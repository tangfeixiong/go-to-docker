package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	dockerref "github.com/docker/distribution/reference"
	dockermessage "github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/engine-api/types"
	"github.com/golang/glog"

	"golang.org/x/net/context"
	"k8s.io/kubernetes/pkg/util/parsers"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pkg/dockerctl"
)

// There are 2 kinds of docker operations categorized by running time:
// * Long running operation: The long running operation could run for arbitrary long time, and the running time
// usually depends on some uncontrollable factors. These operations include: PullImage, Logs, StartExec, AttachToContainer.
// * Non-long running operation: Given the maximum load of the system, the non-long running operation should finish
// in expected and usually short time. These include all other operations.
// kubeDockerClient only applies timeout on non-long running operations.
const (
	// defaultTimeout is the default timeout of short running docker operations.
	defaultTimeout = 2 * time.Minute

	// defaultShmSize is the default ShmSize to use (in bytes) if not specified.
	defaultShmSize = int64(1024 * 1024 * 64)

	// defaultImagePullingProgressReportInterval is the default interval of image pulling progress reporting.
	defaultImagePullingProgressReportInterval = 10 * time.Second

	// defaultImagePullingStuckTimeout is the default timeout for image pulling stuck. If no progress
	// is made for defaultImagePullingStuckTimeout, the image pulling will be cancelled.
	// Docker reports image progress for every 512kB block, so normally there shouldn't be too long interval
	// between progress updates.
	// TODO(random-liu): Make this configurable
	defaultImagePullingStuckTimeout = 1 * time.Minute
)

func base64EncodeAuth(auth types.AuthConfig) (string, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(auth); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buf.Bytes()), nil
}

// progress is a wrapper of dockermessage.JSONMessage with a lock protecting it.
type progress struct {
	sync.RWMutex
	// message stores the latest docker json message.
	message *dockermessage.JSONMessage
	// timestamp of the latest update.
	timestamp time.Time
}

func newProgress() *progress {
	return &progress{timestamp: time.Now()}
}

func (p *progress) set(msg *dockermessage.JSONMessage) {
	p.Lock()
	defer p.Unlock()
	p.message = msg
	p.timestamp = time.Now()
}

func (p *progress) get() (string, time.Time) {
	p.RLock()
	defer p.RUnlock()
	if p.message == nil {
		return "No progress", p.timestamp
	}
	// The following code is based on JSONMessage.Display
	var prefix string
	if p.message.ID != "" {
		prefix = fmt.Sprintf("%s: ", p.message.ID)
	}
	if p.message.Progress == nil {
		return fmt.Sprintf("%s%s", prefix, p.message.Status), p.timestamp
	}
	return fmt.Sprintf("%s%s %s", prefix, p.message.Status, p.message.Progress.String()), p.timestamp
}

// progressReporter keeps the newest image pulling progress and periodically report the newest progress.
type progressReporter struct {
	*progress
	image  string
	cancel context.CancelFunc
	stopCh chan struct{}
}

// newProgressReporter creates a new progressReporter for specific image with specified reporting interval
func newProgressReporter(image string, cancel context.CancelFunc) *progressReporter {
	return &progressReporter{
		progress: newProgress(),
		image:    image,
		cancel:   cancel,
		stopCh:   make(chan struct{}),
	}
}

// start starts the progressReporter
func (p *progressReporter) start() {
	go func() {
		ticker := time.NewTicker(defaultImagePullingProgressReportInterval)
		defer ticker.Stop()
		for {
			// TODO(random-liu): Report as events.
			select {
			case <-ticker.C:
				progress, timestamp := p.progress.get()
				// If there is no progress for defaultImagePullingStuckTimeout, cancel the operation.
				if time.Now().Sub(timestamp) > defaultImagePullingStuckTimeout {
					glog.Errorf("Cancel pulling image %q because of no progress for %v, latest progress: %q", p.image, defaultImagePullingStuckTimeout, progress)
					p.cancel()
					return
				}
				glog.V(2).Infof("Pulling image %q: %q", p.image, progress)
			case <-p.stopCh:
				progress, _ := p.progress.get()
				glog.V(2).Infof("Stop pulling image %q: %q", p.image, progress)
				return
			}
		}
	}()
}

// stop stops the progressReporter
func (p *progressReporter) stop() {
	close(p.stopCh)
}

// applyDefaultImageTag parses a docker image string, if it doesn't contain any tag or digest,
// a default tag will be applied.
// https://github.com/kubernetes/kubernetes/pkg/kubelet/dockertools/docker.go
func applyDefaultImageTag(image string) (string, error) {
	named, err := dockerref.ParseNamed(image)
	if err != nil {
		return "", fmt.Errorf("couldn't parse image reference %q: %v", image, err)
	}
	_, isTagged := named.(dockerref.Tagged)
	_, isDigested := named.(dockerref.Digested)
	if !isTagged && !isDigested {
		named, err := dockerref.WithTag(named, parsers.DefaultImageTag)
		if err != nil {
			return "", fmt.Errorf("failed to apply default image tag %q: %v", image, err)
		}
		image = named.String()
	}
	return image, nil
}

/*
  Inspired from https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockertools/kube_docker_client.go
*/
func (m *myService) pullImage(req *pb.DockerPullData) (*pb.DockerPullData, error) {
	glog.Infoln("Go to pull image", req.Image)
	resp := new(pb.DockerPullData)
	if nil == req || 0 == len(req.Image) {
		return resp, fmt.Errorf("Request required")
	}
	if img, err := applyDefaultImageTag(req.Image); nil != err {
		return resp, err
	} else {
		resp.Image = img
	}

	ctl := dockerctl.NewEngine1_12Client()
	cli, err := ctl.DockerClient()
	if nil != err {
		resp.StateCode = 100
		resp.StateMessage = err.Error()
		return resp, err
	}

	auth := types.AuthConfig{
		// Username: "",
		// Password: "",
		Auth:          "",
		Email:         "",
		ServerAddress: "127.0.0.1:5000",
		// IdentityToken: "",
		// RegistryToken: "",
	}
	auth = ctl.RegistryAuth(resp.Image)

	// RegistryAuth is the base64 encoded credentials for the registry
	base64Auth, err := base64EncodeAuth(auth)
	if err != nil {
		resp.StateCode = 101
		resp.StateMessage = err.Error()
		return resp, err
	}
	opts := types.ImagePullOptions{
		RegistryAuth: base64Auth,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	result, err := cli.ImagePull(ctx, resp.Image, opts)
	if err != nil {
		resp.StateCode = 102
		resp.StateMessage = err.Error()
		return resp, err
	}
	defer result.Close()
	reporter := newProgressReporter(resp.Image, cancel)
	reporter.start()
	defer reporter.stop()
	decoder := json.NewDecoder(result)
	for {
		var msg dockermessage.JSONMessage
		err := decoder.Decode(&msg)
		if err == io.EOF {
			break
		}
		if err != nil {
			resp.StateCode = 103
			resp.StateMessage = err.Error()
			return resp, err
		}
		if msg.Error != nil {
			resp.StateCode = 104
			resp.StateMessage = fmt.Sprintf("code: %d, message: %s, %s", msg.Error.Code, msg.Error.Message, msg.ErrorMessage)
			return resp, fmt.Errorf("Failed to pull image %s; %s", resp.Image, resp.StateMessage)
		}
		reporter.set(&msg)
	}
	resp.ProgressReport, _ = reporter.get()
	return resp, nil
}
