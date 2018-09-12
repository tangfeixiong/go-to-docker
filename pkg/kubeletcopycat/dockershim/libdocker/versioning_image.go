/*
  Inspired by
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockertools/kube_docker_client.go
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockertools/docker.go
*/

package libdocker

import (
	// "bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	// "sync"
	// "time"

	dockertypes "github.com/docker/docker/api/types"
	dockerfilterstypes "github.com/docker/docker/api/types/filters"

	// dockerapi "github.com/docker/docker/client"
	dockermessage "github.com/docker/docker/pkg/jsonmessage"
	"github.com/golang/glog"

	"golang.org/x/net/context"

	// "github.com/tangfeixiong/go-to-docker/pkg/credentialprovider"
	// utilerrors "github.com/tangfeixiong/go-to-docker/pkg/util/errors"
	"github.com/tangfeixiong/go-to-docker/pkg/util/parsers"
)

func PruneImages(cli Interface, ctx context.Context, filters dockerfilterstypes.Args) (dockertypes.ImagesPruneReport, error) {
	kc, ok := cli.(*kubeDockerClient)
	if !ok {
		glog.Errorln("Unable to get docker client")
		return dockertypes.ImagesPruneReport{}, fmt.Errorf("unable to get docker client")
	}
	return kc.client.ImagesPrune(ctx, filters)
}

func PullImage(cli Interface, image string, opts dockertypes.ImagePullOptions) error {
	repoToPull, _, _, err := parsers.ParseImageName(image)
	if err != nil {
		glog.Errorf("Uable to parse image name %q: %v", image, err)
		return err
	}

	username, password, err := decodeDockerConfigFieldAuth(opts.RegistryAuth)
	if err != nil {
		glog.Errorf("Unable to parse as BasicAuth: %v", err)
		return err
	}
	auth := dockertypes.AuthConfig{
		Username: username,
		Password: password,
	}

	if err := cli.PullImage(repoToPull, auth, opts); err != nil {
		glog.Errorf("Uable to pull image %q: %v", image, err)
		return err
	}
	return nil
}

func PullImageAsyncMessaging(cli Interface, gctx context.Context, image string, opts dockertypes.ImagePullOptions, ch chan<- map[int]string) error {
	kc, ok := cli.(*kubeDockerClient)
	if !ok {
		ch <- map[int]string{100: "unable to get docker client"}
		glog.Errorln("Unable to get docker client")
		return fmt.Errorf("unable to get docker client")
	}
	repoToPull, _, _, err := parsers.ParseImageName(image)
	if err != nil {
		ch <- map[int]string{101: err.Error()}
		glog.Errorf("Uable to parse image name %q: %v", image, err)
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result, err := kc.client.ImagePull(ctx, repoToPull, opts)
	if err != nil {
		ch <- map[int]string{102: err.Error()}
		glog.Errorf("Failed to pull image: %v", err)
		return err
	} else {
		ch <- map[int]string{1: "executed"}
	}
	defer result.Close()
	reporter := newProgressReporter(repoToPull, cancel, kc.imagePullProgressDeadline)
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
			return err
		}
		if msg.Error != nil {
			return msg.Error
		}
		reporter.set(&msg)
	}
	s, t := reporter.get()
	glog.Infof("%q, %s", t, s)
	return nil
}

func PushImageAsyncMessaging(cli Interface, gctx context.Context, image string, opts dockertypes.ImagePushOptions, ch chan<- map[int]string) error {
	kc, ok := cli.(*kubeDockerClient)
	if !ok {
		ch <- map[int]string{100: "unable to get docker client"}
		glog.Errorln("Unable to get docker client")
		return fmt.Errorf("unable to get docker client")
	}
	repoToPush, _, _, err := parsers.ParseImageName(image)
	if err != nil {
		ch <- map[int]string{101: err.Error()}
		glog.Errorf("Uable to parse image name %q: %v", image, err)
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result, err := kc.client.ImagePush(ctx, repoToPush, opts)
	if err != nil {
		ch <- map[int]string{102: err.Error()}
		glog.Errorf("Failed to push image: %v", err)
		return err
	} else {
		ch <- map[int]string{1: "executed"}
	}
	defer result.Close()
	reporter := newProgressReporter(repoToPush, cancel, kc.imagePullProgressDeadline)
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
			return err
		}
		if msg.Error != nil {
			return msg.Error
		}
		reporter.set(&msg)
	}
	s, t := reporter.get()
	glog.Infof("%q, %s", t, s)
	return nil
}

// decodeDockerConfigFieldAuth deserializes the "auth" field from dockercfg into a
// username and a password. The format of the auth field is base64(<username>:<password>).
func decodeDockerConfigFieldAuth(field string) (username, password string, err error) {
	decoded, err := base64.StdEncoding.DecodeString(field)
	if err != nil {
		return
	}

	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		err = fmt.Errorf("unable to parse auth field")
		return
	}

	username = parts[0]
	password = parts[1]

	return
}

func BuildImageAsyncMessaging(cli Interface, gctx context.Context, buildContext io.Reader, opts dockertypes.ImageBuildOptions, ch chan<- map[int]string) error {
	kc, ok := cli.(*kubeDockerClient)
	if !ok {
		ch <- map[int]string{100: "unable to get docker client"}
		glog.Errorln("Unable to get docker client")
		return fmt.Errorf("unable to get docker client")
	}
	repoToBuild := opts.Target

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// return: dockertypes.ImageBuildResponse, error
	var result io.ReadCloser
	imageBuildResponse, err := kc.client.ImageBuild(ctx, buildContext, opts)
	if err != nil {
		ch <- map[int]string{102: err.Error()}
		glog.Errorf("Failed to build image: %v", err)
		return err
	} else {
		ch <- map[int]string{1: "executed"}
		result = imageBuildResponse.Body
	}
	defer result.Close()
	reporter := newProgressReporter(repoToBuild, cancel, kc.imagePullProgressDeadline)
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
			return err
		}
		if msg.Error != nil {
			return msg.Error
		}
		reporter.set(&msg)
	}
	s, t := reporter.get()
	glog.Infof("%q, %s", t, s)
	return nil
}
