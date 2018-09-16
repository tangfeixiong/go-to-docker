package dockerclient

import (
	// "net/http"
	// "os"
	"time"

	dockerapi "github.com/docker/docker/client"
	// "github.com/docker/go-connections/tlsconfig"
	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim"
	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim/libdocker"
)

type DockerClient struct {
	DockerApiClient *dockerapi.Client
	// criHandler dockershim.DockerService
	KubeDockerClient libdocker.Interface
	//OpenshiftS2iDockerClient: docker.Docker
}

// refer to
//   https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L614-L620
//
//     ds, err := dockershim.NewDockerService(kubeDeps.DockerClientConfig, crOptions.PodSandboxImage, streamingConfig,
//		   &pluginSettings, runtimeCgroups, kubeCfg.CgroupDriver, crOptions.DockershimRootDirectory, !crOptions.RedirectContainerStreaming)
//	   if err != nil {
//		   return nil, err
//	   }
//	   if crOptions.RedirectContainerStreaming {
//		   klet.criHandler = ds
//	   }
// ...
func NewOrDie() *DockerClient {

	apiClient, err := getDockerClient("")
	if err != nil {
		glog.Fatalf("Couldn't connect to docker: %v", err)
	}

	config := &dockershim.ClientConfig{
		RuntimeRequestTimeout:     2*time.Minute - 1*time.Second,
		ImagePullProgressDeadline: 1 * time.Minute,
	}

	/*
	   https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockershim/docker_service.go#L89-L190
	   - func NewDockerService(config *ClientConfig, ...
	*/
	client := dockershim.NewDockerClientFromConfig(config)

	if client == nil {
		panic("couldn't initialize Docker client from config")
	}

	return &DockerClient{
		DockerApiClient:  apiClient,
		KubeDockerClient: client,
	}
}

// Refer to
// - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockershim/libdocker/client.go#L71-L72
// Get a *Dockerapi.Client, either using the endpoint passed in, or using
// DOCKER_HOST, DOCKER_TLS_VERIFY, and DOCKER_CERT path per their spec
func getDockerClient(dockerEndpoint string) (*dockerapi.Client, error) {
	if len(dockerEndpoint) > 0 {
		glog.Infof("Connecting to docker on %s", dockerEndpoint)
		return dockerapi.NewClient(dockerEndpoint, "", nil, nil)
	}
	return dockerapi.NewEnvClient()
}
