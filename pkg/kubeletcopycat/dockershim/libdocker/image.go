/*
  Inspired by
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockershim/libdocker/kube_docker_client.go
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockertools/kube_docker_client.go
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockertools/docker.go
*/

package libdocker

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	dockerapi "github.com/docker/docker/client"
	dockermessage "github.com/docker/docker/pkg/jsonmessage"
	"github.com/golang/glog"

	"golang.org/x/net/context"

	mobypb "github.com/tangfeixiong/go-to-docker/pb/moby"
)

func (dc *DockerClient) PullImageIntoCachingProgress(ctx context.Context, ref string, opts *mobypb.ImagePullOptions, ch chan<- map[int]error) (int, string, error) {
	if img, err := applyDefaultImageTag(ref); nil != err {
		glog.Warningf("Could not parse image reference: %v", err)
		ch <- map[int]error{ILLEGAL_IMAGE_REF: err}
		return ILLEGAL_IMAGE_REF, fmt.Sprintf("Couldn't parse image ref: %v", err), err
	}

	//	cli, err := client.NewEnvClient()
	//	if nil != err {
	//		glog.Warningf("Failed to generate Docker client: %v", err)
	//		ch <- map[int]error{100: err}
	//		return 100, fmt.Sprintf("Failed to gen Docker client: %v", err), err
	//	}
	cli := dc.MobyClient

	auth := types.AuthConfig{
		// Username: "",
		// Password: "",
		Auth:          "",
		Email:         "",
		ServerAddress: "127.0.0.1:5000",
		// IdentityToken: "",
		// RegistryToken: "",
	}
	var ok bool
	auth, ok = dc.RegistryAuthConfig(img)

	// RegistryAuth is the base64 encoded credentials for the registry
	base64Auth, err := base64EncodeAuth(auth)
	if err != nil {
		glog.Warningf("Failed to encode registry auth from base64: %v", err)
		ch <- map[int]error{101: err}
		return 101, fmt.Sprintf("Failed to encode registry auth from base64: %v", err), err
	}
	opts := types.ImagePullOptions{
		RegistryAuth: base64Auth,
	}
	if !ok {
		opts = types.ImagePullOptions{}
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	glog.Infoln("Go to pull image", ref)
	ch <- map[int]error{0: nil}

	result, err := cli.ImagePull(ctx, img, opts)
	if err != nil {
		glog.Warningf("Failed to pull image: %v", err)
		// ch <- map[int]error{102, err}
		return 102, "Could not pull image", err
	}
	defer result.Close()
	reporter := newProgressReporter(img, cancel)
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
			return 103, "Docker response message decode error", err
		}
		if msg.Error != nil {
			statemsg := fmt.Sprintf("code: %d, message: %s, %s", msg.Error.Code, msg.Error.Message, msg.ErrorMessage)
			return 104, statemsg, fmt.Errorf("Failed to pull image %s; %s", ref, statemsg)
		}
		reporter.set(&msg)
	}
	statemsg, _ := reporter.get()
	return 0, statemsg, nil
}
