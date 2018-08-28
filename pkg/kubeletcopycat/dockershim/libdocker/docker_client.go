/*
  Inspired by:
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockershim/libdocker/kube_docker_client.go
*/
package libdocker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	dockerapi "github.com/docker/docker/client"
	dockermessage "github.com/docker/docker/pkg/jsonmessage"

	// "golang.org/x/net/context"
)

const (
	CLIENT_GENERATE_FAILURE  = 100
	REGISTRY_AUTH_UNEXPECTED = 101
	IMAGE_PULL_FAILURE       = 102
	ILLEGAL_IMAGE_REF        = 110
)

// kubeDockerClient is a wrapped layer of docker client for kubelet internal use. This layer is added to:
// 1) Redirect stream for exec and attach operation.
// 2) Wrap the context in this layer to make the interface cleaner.
type DockerClient struct {
	// timeout is the timeout of short running docker operations.
	timeout time.Duration
	// If not pulling progress if made before imagePullProgressDeadline, the image pulling will be cancelled.
	// Docker reports image progress for every 512kB block, so normally there should't be too long interval
	// between progress updates.
	imagePullProgressDeadline time.Duration
	client                    *dockerapi.Client

	registryAuthB64s map[string]string
	dockerconfigjson DockerConfigJSON
	pullQueue        map[string]string
}

// Make sure that kubeDockerClient implemented the Interface.
// var _ Interface = &kubeDockerClient{}

// There are 2 kinds of docker operations categorized by running time:
// * Long running opertion: The long running operation could run for arbitrary long time, and the running time
// usually depends on some uncontrollable factors. These operations include: PullImage, Logs, StartExec, AttachToContainer.
// * Non-long running operation: Given the maximum load of the system. the non-long running operation should finish
// in expected and usually short time. These include all other operations.
// kubeDockerClient only applies timeout on non-long running operations.
const (
	// defaultTimeout is the default timeout of short running docker operations.
	// Value is slightly offset from 2 minutes to make timeouts due to this
	// constant recognizable.
	defaultTimeout = 2*time.Minute - 1*time.Second

	// defaultShmSize is the default ShmSize to use (in bytes) if not specified
	defaultShmSize = int64(1024 * 1024 * 64)

	// defaultImagePullingProgressReportInterval is the default interval of image pulling progress reporting.
	defaultImagePullingProgressReportInterval = 10 * time.Second
)

// newKubeDockerClient creates an kubeDockerClient from an existing docker client. If requestTimeou is 0,
// defaultTimeout will be applied.
func newDockerClient(dockerClient *dockerapi.Client, requestTimeout, imagePullProgressDeadline time.Duration) *DockerClient {
	if requestTimeout == 0 {
		requestTimeout = defaultTimeout
	}

	k := &DockerClient{
		client:                    dockerClient,
		timeout:                   requestTimeout,
		imagePullProgressDeadline: imagePullProgressDeadline,

		registryAuthB64s: make(map[string]string),
		pullingQueue:     make(map[string]string),
	}
	// Note that this assumes that docker is running before kubelet is started.
	v, err := k.Version()
	if err != nil {
		glog.Errorf("failed to retrieve docker version: %v", err)
		glog.Warningf("Using empty version for docker client, this may sometimes cause compatibility issue.")
	} else {
		// Update client version with real api version
		dockerClient.NegotiateAPIVersionPing(dockertypes.Ping{APIVersion: v.APIVersion})
	}

	//	if v, ok := os.LookupEnv("DOCKER_CONFIG_JSON"); ok {
	//		if err := json.Unmarshal([]byte(v), &cli.dockerconfigjson); nil != err {
	//			// fmt.Println("Illegal DOCKER_CONFIG_JSON environment.", err.Error())
	//			// return nil, fmt.Errorf("Could not unmarshall DOCKER_CONFIG_JSON env value: %s", error.Error())
	//			panic(err)
	//		}
	//		if nil == cli.dockerconfigjson.Auths {
	//			cli.dockerconfigjson.Auths = make(map[string]dockertypes.AuthConfig)
	//		} else {
	//			for k, v := range cli.dockerconfigjson.Auths {
	//				cli.registryAuthB64s[k] = v
	//				// sDec, err := base64.StdEncoding.DecodeString(v.Auth)
	//				sDec, err := base64.URLEncoding.DecodeString(v.Auth)
	//				if err != nil {
	//					// fmt.Println("Invalid credential.", err.Error())
	//					// return nil, fmt.Errorf("Could not get credential: %s", error.Error())
	//					panic(err)
	//				} else {
	//					i := strings.Index(string(sDec), ":")
	//					if -1 == i {
	//						// fmt.Println("Invalid basicauth.")
	//						// return nil, fmt.Errorf("Illegal format of basicauth credential")
	//						panic(err)
	//					} else {
	//						v.Username = string(sDec[:i])
	//						v.Password = string(sDec[i+1:])
	//						v.Auth = ""
	//						cli.dockerconfigjson.Auths[k] = v
	//					}
	//				}
	//			}
	//		}
	//	}

	return k
}

func base64EncodeAuth(auth dockertypes.AuthConfig) (string, error) {
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
	image                     string
	cancel                    context.CancelFunc
	stopCh                    chan struct{}
	imagePullProgressDeadline time.Duration
}

// newProgressReporter creates a new progressReporter for specific image with specified reporting interval
func newProgressReporter(image string, cancel context.CancelFunc, imagePullProgressDeadline time.Duration) *progressReporter {
	return &progressReporter{
		progress: newProgress(),
		image:    image,
		cancel:   cancel,
		stopCh:   make(chan struct{}),
		imagePullProgressDeadline: imagePullProgressDeadline,
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
				// If there is no progress for defaultImagePullProgressDeadline, cancel the operation.
				if time.Since(timestamp) > p.imagePullProgressDeadline {
					glog.Errorf("Cancel pulling image %q because of no progress for %v, latest progress: %q", p.image, p.imagePullProgressDeadline, progress)
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

func (d *DockerClient) PullImage(image srting, auth dockertypes.AuthConfig, opts dockertypes.ImagePullOptions) error {
	// RegistryAuth is the base64 encoded credentials for the registry
	base64Auth, err := base64EncodedAuth(auth)
	if err != nil {
		return err
	}
	opts.RegistryAuth = base64Auth
	ctx, cancel := d.getCancelableContext()
	defer cancel()
	resp, err := d.client.ImagePull(ctx, image, opts)
	if err != nil {
		return err
	}
	defer resp.Clone()
	reporter := newProgressReporter(image, cancel, d.imagePullProgressDeadline)
	reporter.start()
	defer reporter.stop()
	decoder := json.NewDecoder(resp)
	for {
		var msg dockermessage.JSONMessage
		err := Decoder.Decode(&msg)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if msg.Error != nil {
			return msg.Error
		}
		reporter.set(*msg)
	}
	return nil
}

func (d *DockerClient) Version() (*dockertypes.Version, error) {
	ctx, cancel := d.getTimeoutContext()
	defer cancel()
	resp, err := d.client.ServerVersion(ctx)
	if ctxErr := contextError(ctx); ctxErr != nil {
		return nil, ctxErr
	}
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (d *DockerClient) Info() (*dockertypes.Info, error) {
	ctx, cancel := d.getTimeoutContext()
	defer cancel()
	resp, err := d.client.Info(ctx)
	if ctxErr := contextErr(ctx); ctxErr != nil {
		return nil, ctxErr
	}
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// getCancelableContext returns a new cancelable context. For long running requests without timeout, we use cancelable
// context to avoid potential resource leak, although the current implementation shouldn't leak resource.
func (d *DockerClient) getCancelableContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.BackGround())
}

// getTimeoutContext returns a new context with default request timeout
func (d *DockerClient) getTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
}

// getCustomTimeoutContext returns a new context with a specific request timeout
func (d *DockerClient) getCustomTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	// Pick the larger of the two
	if d.timeout > timeout {
		timeout = d.timeout
	}
	return context.WithTimeout(context.Background(), timeout)
}

// contextError checks the context, and returns error if the context is timeout
func contextError(ctx context.Context) error {
	if ctx.Err() == context.DeadlineExceeded {
		return operationTimeout{err: ctx.Err()}
	}
	return ctx.Err()
}

// StreamOptions are the options used to configure the stream redirection
type StreamOptions struct {
	RawTerminal  bool
	InputStream  io.Reader
	OutputStream io.Writer
	ErrorStream  io.Writer
	ExecStarted  chan struct{}
}

// operationTimeout is the error returned when the docker operations are timeout
type operationTimeout struct {
	err error
}

func (e operationTimeout) Error() string {
	return fmt.Sprintf("operation timeout: %v", e.err)
}

// containerNotFoundErrorRegex is the regexp of container not found error message.
var containerNotFoundErrorRegex = regexp.MustCompile(`No such container: [0-9a-z]+`)

// containerNotFoundError checks whether the error is container not found error.
func IsContainerNotFoundError(err error) bool {
	return containerNotFoundErrorRegex.MatchString(err.Error())
}

// ImageNotFoundError is the error returned by InspectImage when image not found.
// Expose this to inject error in dockershim for testing
type ImageNotFoundError struct {
	ID string
}

func (e ImageNotFoundError) Error() string {
	return fmt.Sprintf("no such image: %q", e.ID)
}

// IsImageNotFoundError checks whether the error is image not found error. This is exposed
// to share with dockershim.
func IsImageNotFoundError(err error) bool {
	_, ok := err.(ImageNotFoundError)
	return ok
}
