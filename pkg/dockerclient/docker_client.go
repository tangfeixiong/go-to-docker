package dockerclient

import (
	// "net/http"
	// "os"
	"time"

	dockerapi "github.com/docker/docker/client"
	// "github.com/docker/go-connections/tlsconfig"
	"github.com/golang/glog"

	// openshiftstiapi "github.com/openshift/source-to-image/pkg/api"
	// openshiftstidocker "github.com/openshift/source-to-image/pkg/docker"

	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim"
	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim/libdocker"
)

type DockerClient struct {
	DockerApiClient *dockerapi.Client
	// criHandler dockershim.DockerService
	KubeDockerClient libdocker.Interface
	//OpenshiftStiDockerClient: docker.Docker
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

// Refer to
// - https://github.com/openshift/source-to-image/blob/master/pkg/docker/docker.go#L276-L277
// NewEngineAPIClient creates a new Docker engine API client
/*
func NewEngineAPIClient(config *openshiftstiapi.DockerConfig) (*dockerapi.Client, error) {
	var httpClient *http.Client

	if config.UseTLS || config.TLSVerify {
		tlscOptions := tlsconfig.Options{
			InsecureSkipVerify: !config.TLSVerify,
		}

		if _, err := os.Stat(config.CAFile); !os.IsNotExist(err) {
			tlscOptions.CAFile = config.CAFile
		}
		if _, err := os.Stat(config.CertFile); !os.IsNotExist(err) {
			tlscOptions.CertFile = config.CertFile
		}
		if _, err := os.Stat(config.KeyFile); !os.IsNotExist(err) {
			tlscOptions.KeyFile = config.KeyFile
		}

		tlsc, err := tlsconfig.Client(tlscOptions)
		if err != nil {
			return nil, err
		}

		httpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsc,
			},
		}
	}
	return dockerapi.NewClient(config.Endpoint, os.Getenv("DOCKER_API_VERSION"), httpClient, nil)
}
*/

// Refer to
// - https://github.com/openshift/source-to-image/blob/master/pkg/docker/docker.go#L309
// New creates a new implementation of the STI Docker interface
/*
func New(client Client, auth api.AuthConfig) Docker {
	return &stiDocker{
		client: client,
		pullAuth: dockertypes.AuthConfig{
			Username:      auth.Username,
			Password:      auth.Password,
			Email:         auth.Email,
			ServerAddress: auth.ServerAddress,
		},
	}
}
*/
