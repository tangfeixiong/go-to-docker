package dockerctl

import (
	"context"
	"fmt"
	"os"

	// dockerdigest "github.com/docker/distribution/digest"
	dockerref "github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
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
	glog.Infof("Go to create container: %s", containername)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infof("Could not instance moby client: %v", err)
		return container.ContainerCreateCreatedBody{}, err
	}

	containerCreateResponse, err := cli.ContainerCreate(context.Background(), config, hostconfig, networkconfig, containername)
	if err != nil {
		glog.V(2).Infof("Could not create container: %v", err)
		return container.ContainerCreateCreatedBody{}, err
	}
	return containerCreateResponse, nil
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
