package dockerclient

import (
	"context"
	"fmt"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	dockerfilters "github.com/docker/docker/api/types/filters"
	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pb"
	mobypb "github.com/tangfeixiong/go-to-docker/pb/moby"
	containerpb "github.com/tangfeixiong/go-to-docker/pb/moby/container"
	"github.com/tangfeixiong/go-to-docker/pkg/credentialprovider"
	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim/libdocker"
	utilerrors "github.com/tangfeixiong/go-to-docker/pkg/util/errors"
	"github.com/tangfeixiong/go-to-docker/pkg/util/parsers"
)

func (cli *DockerClient) RunContainer(ctx context.Context, req *pb.DockerContainerRunReqResp) (*pb.DockerContainerRunReqResp, error) {
	resp, err := req.DeepCopyChecked()
	if err != nil {
		glog.Errorf("Unable running container: %v", err)
		return resp, fmt.Errorf("unable running container: %v", err)
	}
	client := cli.KubeDockerClient

	err = ensureDockerImageExists(client, req.Config.Image)
	if err != nil {
		glog.Errorf("Unable to ensure image %q of running container: %v", req.Config.Image, err)
		return resp, fmt.Errorf("failed to ensure image (%q) of running container: %v", req.Config.Image, err)
	}

	created, err := client.CreateContainer(req.ExportAsDockerApiTypeContainerCreateConfig())
	if err != nil {
		glog.Errorf("Failed to create container: %v", err)
		return resp, err
	}
	resp.ContainerCreateCreatedBody = containerpb.ConvertFromDockerApiTypeContainerCreateCreatedBody(created)
	resp.StateCode = 1
	resp.StateMessage = "created"
	glog.Infof("docker container %v created", resp.ContainerCreateCreatedBody.Id)

	err = client.StartContainer(created.ID)
	if err != nil {
		glog.Errorf("Failed to start container: %v", err)
		return resp, err
	}
	resp.StateCode = 0
	resp.StateMessage = "started"
	glog.Infof("docker container %v started", resp.ContainerCreateCreatedBody.Id)
	return resp, nil

}

// inspired by: https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockershim/helpers.go#L315
// ensureSandboxImageExists pulls the sandbox image when it's not present
//
func ensureDockerImageExists(client libdocker.Interface, image string) error {
	_, err := client.InspectImageByRef(image)
	if err == nil {
		return nil
	}
	if !libdocker.IsImageNotFoundError(err) {
		return fmt.Errorf("failed to inspect image %q: %v", image, err)
	}

	repoToPull, _, _, err := parsers.ParseImageName(image)
	if err != nil {
		return err
	}

	keyring := credentialprovider.NewDockerKeyring()
	creds, withCredentials := keyring.Lookup(repoToPull)
	if !withCredentials {
		glog.V(3).Infof("Pulling image %q without credentials", image)

		err := client.PullImage(image, dockertypes.AuthConfig{}, dockertypes.ImagePullOptions{})
		if err != nil {
			return fmt.Errorf("failed pulling image %q: %v", image, err)
		}

		return nil
	}

	var pullErrs []error
	for _, currentCreds := range creds {
		authConfig := credentialprovider.LazyProvide(currentCreds)
		err := client.PullImage(image, authConfig, dockertypes.ImagePullOptions{})
		// If there was no error, return success
		if err == nil {
			return nil
		}

		pullErrs = append(pullErrs, err)
	}

	return utilerrors.NewAggregate(pullErrs)
}

func (cli *DockerClient) ListContainers(ctx context.Context, req *pb.DockerContainerListReqResp) (*pb.DockerContainerListReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable listing containers: %v", err)
		return resp, fmt.Errorf("unable to list containers: %v", err)
	}
	client := cli.KubeDockerClient

	affected, err := client.ListContainers(req.ContainerListOptions.ExportIntoDockerApiType())
	if nil != err {
		glog.Errorf("Failed to list containers: %v", err)
		return resp, fmt.Errorf("failed to list docker containers: %v", err)
	}
	resp.Containers = make([]*mobypb.Container, len(affected))
	for _, v := range affected {
		resp.Containers = append(resp.Containers, mobypb.ConvertFromDockerApiTypeContainer(v))
	}
	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	glog.Infoln("docker container listed")
	return resp, nil
}

func (cli *DockerClient) InspectContainer(ctx context.Context, req *pb.DockerContainerInspectReqResp) (*pb.DockerContainerInspectReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable inspecting container: %v", err)
		return resp, fmt.Errorf("uable to inspect docker container: %v", err)
	}
	client := cli.KubeDockerClient

	switch req.KeyType {
	case pb.DockerContainerInspectReqResp_NAME:
		cond := dockerfilters.NewArgs()
		cond.Add("name", req.Name)
		affected, err := client.ListContainers(dockertypes.ContainerListOptions{
			Filters: cond,
		})
		if nil != err {
			glog.Errorf("Unable inspecting container by name: %v", err)
			return resp, fmt.Errorf("unalbe to inspect docker container by name: %v", err)
		}
		var ok bool
	LOOP:
		for _, v := range affected {
			for _, v1 := range v.Names {
				if v1 == req.Name {
					req.Id = v.ID
					resp.Id = v.ID
					ok = true
					break LOOP
				}
			}
		}
		if !ok {
			resp.StateCode = 100
			glog.Errorln("Could not find correspond cotainer name in list")
			return resp, fmt.Errorf("Could not find correspond container name %s", req.Name)
		}
		fallthrough
	case pb.DockerContainerInspectReqResp_ID:
		detailed, err := client.InspectContainer(resp.Id)
		if nil != err {
			glog.Errorf("Failed to inspect docker container: %v", err)
			return resp, fmt.Errorf("Failed to inspect docker container: %v", err)
		}
		resp.ContainerJson = mobypb.ConvertFromDockerApiTypeContainerJSON(detailed)
	}
	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	glog.Infof("docker container %s inspected", resp.Id)
	return resp, nil
}

func (cli *DockerClient) RemoveContainer(ctx context.Context, req *pb.DockerContainerRemoveReqResp) (*pb.DockerContainerRemoveReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable removing container: %v", err)
		return resp, fmt.Errorf("unable to remove container: %v", err)
	}
	client := cli.KubeDockerClient

	switch req.KeyType {
	case pb.DockerContainerRemoveReqResp_NAME:
		cond := dockerfilters.NewArgs()
		cond.Add("name", req.Name)
		affected, err := client.ListContainers(dockertypes.ContainerListOptions{
			Filters: cond,
		})
		if nil != err {
			glog.Errorf("Unable to remove container by name: %v", err)
			return resp, fmt.Errorf("unalbe to remove docker container by name: %v", err)
		}
		var ok bool
	LOOP:
		for _, v := range affected {
			for _, v1 := range v.Names {
				if v1 == req.Name {
					req.Id = v.ID
					resp.Id = v.ID
					ok = true
					break LOOP
				}
			}
		}
		if !ok {
			resp.StateCode = 100
			glog.Errorln("Could not find correspond cotainer name in list")
			return resp, fmt.Errorf("Could not find correspond container name %s", req.Name)
		}
		fallthrough
	case pb.DockerContainerRemoveReqResp_ID:
		if err := client.StopContainer(resp.Id, 60*time.Second); err != nil {
			// Stopping an already stopped container will not cause an error in dockerapi
			glog.Warningf("Could not stop container before removing: %v", err)
		}
		err := client.RemoveContainer(resp.Id, req.ContainerRemoveOptions.ExportIntoDockerApiType())
		if nil != err {
			glog.Errorf("Failed to remove docker container: %v", err)
			return resp, fmt.Errorf("Failed to remove docker container: %v", err)
		}
	}
	resp.StateCode = 0
	resp.StateMessage = "removed"
	glog.Infof("docker container %s removed", resp.Id)
	return resp, nil
}

func (cli *DockerClient) PruneContainers(ctx context.Context, req *pb.DockerContainerPruneReqResp) (*pb.DockerContainerPruneReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable to prune docker containers: %v", err)
		return resp, fmt.Errorf("Unable to prune docker containers: %v", err)
	}
	report, err := cli.DockerApiClient.ContainersPrune(ctx, req.Filters.ExportIntoDockerApiType())
	if err != nil {
		glog.Errorf("Failed to prune docker containers: %v", err)
		return resp, fmt.Errorf("Failed to prune docker containers: %v", err)
	}

	resp.ContainersPruneReport = mobypb.ConvertFromDockerApiTypeContainersPruneReport(report)
	resp.StateCode = 0
	resp.StateMessage = "pruned"
	glog.Infof("docker containers pruned")
	return resp, nil
}
