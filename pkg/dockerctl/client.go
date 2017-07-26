package dockerctl

import (
	"context"
	"fmt"
	"os"
	"time"

	// dockerdigest "github.com/docker/distribution/digest"
	dockerref "github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/golang/glog"

	"k8s.io/kubernetes/pkg/util/parsers"
)

var (
	Default_docker_API_ver = "1.12"
)

type MobyClient struct {
}

func NewMobyClient(apiversion string) *MobyClient {
	var err error = nil
	if v, ok := os.LookupEnv("DOCKER_API_VERSION"); !ok {
		v := Default_docker_API_ver
		if 0 != len(apiversion) {
			v = apiversion
		}
		err = os.Setenv("DOCKER_API_VERSION", v)
	} else {
		if 0 != len(apiversion) {
			v = apiversion
		}
		err = os.Setenv("DOCKER_API_VERSION", v)
	}
	if err != nil {
		fmt.Println("Failed to configure DOCKER_API_VERSION environment.", err.Error())
	}
	return &MobyClient{}
}

func (mc *MobyClient) CreateContainer(config *container.Config, hostconfig *container.HostConfig, networkconfig *network.NetworkingConfig, containername string) (container.ContainerCreateCreatedBody, error) {
	glog.Infoln("Go to create container:", containername, "DOCKER_API_VERSION=", os.Getenv("DOCKER_API_VERSION"))

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instantiate moby:", err.Error())
		return container.ContainerCreateCreatedBody{}, fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	resp, err := cli.ContainerCreate(context.Background(), config, hostconfig, networkconfig, containername)
	if err != nil {
		glog.V(2).Infoln("Could not create container:", err.Error())
		return container.ContainerCreateCreatedBody{}, fmt.Errorf("Failed to create container. %v", err)
	}
	return resp, nil
}

func (mc *MobyClient) StartContainer(containerid string) error {
	glog.Infoln("Go to start container:", containerid)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instance moby:", err.Error())
		return fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	opt := types.ContainerStartOptions{}
	err = cli.ContainerStart(context.Background(), containerid, opt)
	if nil != err {
		glog.V(2).Infoln("Could not start container:", err.Error())
		return fmt.Errorf("Failed to start container. %v", err)
	}
	return nil
}

func (mc *MobyClient) StopContainer(containerid string, timeout time.Duration) error {
	glog.Infoln("Go to stop container:", containerid)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instance moby:", err.Error())
		return fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	err = cli.ContainerStop(context.Background(), containerid, &timeout)
	if nil != err {
		glog.V(2).Infoln("Could not stop container:", err.Error())
		return fmt.Errorf("Failed to stop container. %v", err)
	}
	return nil
}

func (mc *MobyClient) RemoveContainer(containerid string) error {
	glog.Infoln("Go to remove container:", containerid)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instance moby:", err.Error())
		return fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	opt := types.ContainerRemoveOptions{}
	err = cli.ContainerRemove(context.Background(), containerid, opt)
	if nil != err {
		glog.V(2).Infoln("Could not remove container:", err.Error())
		return fmt.Errorf("Failed to remove container. %v", err)
	}
	return nil
}

func (mc *MobyClient) CompareContainer(containerid string) ([]container.ContainerChangeResponseItem, error) {
	glog.Infoln("Go to check container runtime different:", containerid)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instantiate moby:", err.Error())
		return []container.ContainerChangeResponseItem{}, fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	resp, err := cli.ContainerDiff(context.Background(), containerid)
	if err != nil {
		glog.V(2).Infoln("Could not create container:", err.Error())
		return []container.ContainerChangeResponseItem{}, fmt.Errorf("Failed to create container. %v", err)
	}
	return resp, nil
}

func (mc *MobyClient) CommitContainer(container string, options types.ContainerCommitOptions) (types.IDResponse, error) {
	glog.Infoln("Go to commit container:", container)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instantiate moby:", err.Error())
		return types.IDResponse{}, fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	resp, err := cli.ContainerCommit(context.Background(), container, options)
	if err != nil {
		glog.V(2).Infoln("Could not commit container:", err.Error())
		return types.IDResponse{}, fmt.Errorf("Failed to commit container. %v", err)
	}
	return resp, nil
}

func (mc *MobyClient) SearchImages(term string, options types.ImageSearchOptions) ([]registry.SearchResult, error) {
	glog.Infoln("Go to search images into registry:", term)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instantiate moby:", err.Error())
		return []registry.SearchResult{}, fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	resp, err := cli.ImageSearch(context.Background(), term, options)
	if err != nil {
		glog.V(2).Infoln("Could not search images:", err.Error())
		return []registry.SearchResult{}, fmt.Errorf("Failed to search images. %v", err)
	}
	return resp, nil
}

func (mc *MobyClient) InspectImage(imageid string) (types.ImageInspect, []byte, error) {
	glog.Infoln("Go to inspect image:", imageid)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instantiate moby:", err.Error())
		return types.ImageInspect{}, nil, fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	resp, raw, err := cli.ImageInspectWithRaw(context.Background(), imageid)
	if err != nil {
		glog.V(2).Infoln("Could not inspect image:", err.Error())
		return types.ImageInspect{}, nil, fmt.Errorf("Failed to inspect image. %v", err)
	}
	return resp, raw, nil
}

func (mc *MobyClient) ListImages(options types.ImageListOptions) ([]types.ImageSummary, error) {
	glog.Infoln("Go to list images archive")

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instantiate moby:", err.Error())
		return []types.ImageSummary{}, fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	resp, err := cli.ImageList(context.Background(), options)
	if err != nil {
		glog.V(2).Infoln("Could not list images archive:", err.Error())
		return []types.ImageSummary{}, fmt.Errorf("Failed to list images archive. %v", err)
	}
	return resp, nil
}

func (mc *MobyClient) RemoveImage(imageid string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error) {
	glog.Infoln("Go to remove image:", imageid)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infoln("Could not instantiate moby:", err.Error())
		return []types.ImageDeleteResponseItem{}, fmt.Errorf("Failed to instantiate moby. %v", err)
	}

	resp, err := cli.ImageRemove(context.Background(), imageid, options)
	if err != nil {
		glog.V(2).Infoln("Could not remove image:", err.Error())
		return []types.ImageDeleteResponseItem{}, fmt.Errorf("Failed to remove image. %v", err)
	}
	return resp, nil
}

func (mc *MobyClient) PullImage(imgref string, registryauth string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ir, err := cli.ImagePull(context.Background(), imgref, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	if ir == nil {
		panic("image stream not ready")
	}
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
func dockerPull(image string, secrets []api.Secret) error {
	// If the image contains no tag or digest, a default tag should be applied.
	image, err := applyDefaultImageTag(image)
	if err != nil {
		return err
	}

	keyring, err := credentialprovider.MakeDockerKeyring(secrets, p.keyring)
	if err != nil {
		return err
	}

	// The only used image pull option RegistryAuth will be set in kube_docker_client
	opts := dockertypes.ImagePullOptions{}

	creds, haveCredentials := keyring.Lookup(image)
	if !haveCredentials {
		glog.V(1).Infof("Pulling image %s without credentials", image)

		err := p.client.PullImage(image, dockertypes.AuthConfig{}, opts)
		if err == nil {
			// Sometimes PullImage failed with no error returned.
			exist, ierr := p.IsImagePresent(image)
			if ierr != nil {
				glog.Warningf("Failed to inspect image %s: %v", image, ierr)
			}
			if !exist {
				return fmt.Errorf("image pull failed for unknown error")
			}
			return nil
		}

		// Image spec: [<registry>/]<repository>/<image>[:<version] so we count '/'
		explicitRegistry := (strings.Count(image, "/") == 2)
		// Hack, look for a private registry, and decorate the error with the lack of
		// credentials.  This is heuristic, and really probably could be done better
		// by talking to the registry API directly from the kubelet here.
		if explicitRegistry {
			return fmt.Errorf("image pull failed for %s, this may be because there are no credentials on this request.  details: (%v)", image, err)
		}

		return filterHTTPError(err, image)
	}

	var pullErrs []error
	for _, currentCreds := range creds {
		err = p.client.PullImage(image, credentialprovider.LazyProvide(currentCreds), opts)
		// If there was no error, return success
		if err == nil {
			return nil
		}

		pullErrs = append(pullErrs, filterHTTPError(err, image))
	}

	return utilerrors.NewAggregate(pullErrs)
}

func dockerThrottledPull(image string, secrets []api.Secret) error {
	if p.limiter.TryAccept() {
		return p.puller.Pull(image, secrets)
	}
	return fmt.Errorf("pull QPS exceeded.")
}
*/
