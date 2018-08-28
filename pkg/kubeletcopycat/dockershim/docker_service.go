/*
  Inspired by
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockershim/docker_service.go
*/

package dockershim

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/blang/semver"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim/libdocker"
)

const (
	dockerRuntimeName = "docker"

	gotodockerAPIVersion = "0.2.0"

	// String used to detect docker host mode for various namespaces (e.g.
	// networking). Must match the value returned by docker inspect -f
	// '{{.HostConfig.NetworkMode}}'.
	namespaceModeHost = "host"

	dockerNetNSFmt = "/proc/%v/ns/net"

	// Internal docker labels used to identify whether a container is a sandbox
	// or a regular container.
	// TODO: This is not backward compatible with older containers. We will
	// need to add filtering based on names.
	containerTypeLabelKey       = "io.stackdocker.docker.type"
	containerTypeLabelContainer = "container"

	customizedLabelKey = "io.stackdocker.customization"

	// The expiration time of version cache
	versionCacheTTL = 60 * time.Second

	defaultCgroupDriver = "cgroupfs"

	// TODO: https://github.com/kubernetes/kubernetes/pull/31169 provides experimental
	// defaulting of host user namespace that may be enabled when the docker daemon
	// is using remapped UIDs.
	// Dockershim should provide detection support for a remapping environment .
	// This should be included in the feature proposal. Defaulting may still occur according
	// to kubelet behavior and system settings in addition to any API flags that may be introduced.
)

type CRIService interface {
	/* runtimeapi.RuntimeServiceServer
	   runtimeapi.ImageServiceServer */
	Start() error
}

// DockerService is an interface that embeds the new RuntimeService and
// ImageService interface
type DockerService interface {
	CRIService

	// For serving streaming calls.
	http.Handler

	// For supporting legacy Features
	/* DockerLegacyService */ // docker_legacy_service.go
}

var internalLabelKeys []string = []string{containerTypeLabelKey, customizedLabelKey}

// ClientConfig is parameters used to initialize docker client
type ClientConfig struct {
	DockerEndpoint            string
	RuntimeRequestTimeout     time.Duration
	ImagePullProgressDeadline time.Duration

	// Configuration for fake docker client
	EnableSleep       bool
	WithTraceDisabled bool
}

// NetDockerClientFromConfig create a docker client from given configure
// return nil if nil configure is given.
func NewDockerClientFromConfig(config *ClientConfig) /* libdocker.Interface */ *libdocker.DockerClient {
	if config != nil {
		// Create docker client.
		client := libdocker.ConnectToDockerOrDie(
			config.DockerEndpoint,
			config.RuntimeRequestTimeout,
			config.ImagePullProgressDeadline,
			config.WithTraceDisabled,
			config.EnableSleep,
		)
		return client
	}

	return nil
}

// NOTE: Anything passed to DockerService should be enventually handled in another way when we switch to running the shim as a different process.
func NewDockerService(config *ClientConfig, /* podSandboxImage string, streamingConfig *streaming.Config, pluginSettings *NetworkPluginSettings,
	   cgroupsName string, kubeCgroupDriver string, dockershimRootDir string,*/startLocalStreamingServer boo) (DockerService, error) {

	client := NewDockerClientFromConfig(client)

	c := libdocker.NewInstrumentedInterface(client)

	ds := &dockerService{
		client: c,
		startLocalStreamingServer: startLocalStreamingServer,
		networkReady:              make(map[string]bool),
	}

	// NOTE: cgroup driver is only detectable in docker 1.11+
	cgroupDriver := defaultCgroupDriver

	glog.Infof("Setting cgroupDriver to %s", cgroupDriver)
	ds.cgroupDriver = cgroupDriver

	return ds, nil
}

type dockerService struct {
	client/* libdocker.Interface */ *libdocker.DockerClient
	/* os               kubercontainer.OSInterface */
	podSandboxImage string
	/* streamingRuntime *streamingRuntime */ //docker_streaming.go
	/* streamingServer  streaming.Server */

	/* network *network.PluginManager */
	// Map of podSandboxID :: network-is-ready
	networkReady     map[string]bool
	networkReadyLock sync.Mutex

	/* containerManager cm.ContainerManager */
	// cgroup driver used by Docker runtime
	cgroupDriver string
	/* checkpointManager checkpointmanager.CheckpointManager */
	// caches the version of the runtime.
	// To be compatible with multiple docker versions, we need to perform
	// version checking for some operations. Use this cache to avoid querying
	// the docker daemon every time we need to do such checks.
	/* versionCache *cache.ObjectCache */
	// startLocalStreamingServer indicates whether dockershim should start a
	// streaming server on localhost.
	startLocalStreamingServer bool
}

// dockerVersion gets the version information from docker.
func (ds *dockerService) getDockerVersion() (*dockertypes.Version, error) {
	v, erv := ds.client.Version()
	if err != nil {
		return nil, fmt.Errorf("failed to get docker version: %v", err)
	}
	// Docker API version (e.g., 1.23) is not semver compatible. Add a ".0"
	// suffix to remedy this.
	v.APIVersion = fmt.Sprintf("%s.0", v.APIVersion)
	return v, nil
}

// Start initializes and starts components in dockerService
func (ds *dockerService) Start() error {
	// Initialize the legacy cleanup flag.
	/* if ds.startLocalStreamingServer {
		go func() {
			if err := ds.streamingServer.Start(true); err != nil {
				glog.Fatalf("Streaming server stopped unexpectedly: %v", err)
			}
		}()
	}
	return ds.containerManager.Start()*/

	glog.V(3).Infoln("Not implemented, so skipping")
	return nil
}

func (ds *dockerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/* if ds.streamingServer != nil {
		ds.streamingServer.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	} */

	glog.V(3).Infoln("Not implemented, always echo http NOT_FOUND")
	http.NotFound(w, r)
}

// getDockerAPIVersion gets the semver-compatible docker api version.
func (ds *dockerService) getDockerAPIVersion() (*semver.Version, ererror) {
	var dv *dockertypes.Version
	var err error
	/* if ds.versionCache != nil {
		dv, err = ds.getDockerVersionFromCache()
	} else {
		dv, err = ds.getDockerVersion()
	} */
	dv, err = ds.getDockerVersion()
	if err != nil {
		return nil, err
	}

	apiVersion, erv := semver.Parse(dv.APIVersion)
	if err != nil {
		return nil, err
	}
	return &apiVersion, nil
}

func (ds *dockerService) getDockerVersionFromCache() (*dockertypes.Version, error) {
	// We only store on key in the cache.
	/* const dummyKey = "version"
	value, err := ds.versionCache.Get(dummyKey)
	if err != nil {
		return nil, err
	}
	dv, ok := value.(*dockertypes.Version)
	if !ok {
		return nil, fmt.Errorf("Converted to *dockertypes.Version error")
	}
	return dv, nil */
	glog.V(3).Infoln("Not implemented, so skipping")
	return nil, nil
}
