package dockerclient

import (
	"bytes"
	"context"
	"fmt"

	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pb"
	mobypb "github.com/tangfeixiong/go-to-docker/pb/moby"
	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim/libdocker"
)

const (
	ILLEGAL_PARAMETER = 1000
	DOCKER_NOT_READY  = 1001

	IMAGE_PULL_STARTED  = 1
	IMAGE_PUSH_STARTED  = 1
	IMAGE_BUILD_STARTED = 1
)

func (cli *DockerClient) ListImages(ctx context.Context, req *pb.DockerImageListReqResp) (*pb.DockerImageListReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable to list images: %v", err)
		return resp, fmt.Errorf("unable to list docker images: %v", err)
	}
	client := cli.KubeDockerClient

	affected, err := client.ListImages(req.ImageListOptions.ExportIntoDockerApiType())
	if err != nil {
		glog.Errorf("Failed to list images: %v", err)
		return resp, fmt.Errorf("failed to list docker images: %v", err)
	}
	for _, v := range affected {
		resp.ImageSummaries = append(resp.ImageSummaries, mobypb.ConvertFromDockerApiTypeImageSummary(v))
	}
	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	glog.Infoln("docker image listed")
	return resp, nil
}

func (cli *DockerClient) InspectImage(ctx context.Context, req *pb.DockerImageInspectReqResp) (*pb.DockerImageInspectReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable to inspect image: %v", err)
		return resp, fmt.Errorf("unable to inspect image: %v", err)
	}
	client := cli.KubeDockerClient

	switch req.KeyType {
	case pb.DockerImageInspectReqResp_REF:
		affected, err := client.InspectImageByRef(req.Ref)
		if err != nil {
			glog.Errorf("failed to inspect image by ref: %v", err)
			return resp, fmt.Errorf("failed to inspect image by ref: %v", err)
		}
		resp.ImageInspect = mobypb.ConvertFromDockerApiTypeImageInspect(affected)
		break
	case pb.DockerImageInspectReqResp_ID:
		affected, err := client.InspectImageByID(req.Id)
		if err != nil {
			glog.Errorf("Failed to inspect image: %v", err)
			return resp, fmt.Errorf("failed to inspect image: %v", err)
		}
		resp.ImageInspect = mobypb.ConvertFromDockerApiTypeImageInspect(affected)
		break
	}
	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	glog.Infof("docker image inspected")
	return resp, nil
}

func (cli *DockerClient) RemoveImage(ctx context.Context, req *pb.DockerImageRemoveReqResp) (*pb.DockerImageRemoveReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable to remove image: %v", err)
		return resp, fmt.Errorf("unable to remove docker image: %v", err)
	}
	client := cli.KubeDockerClient

	switch req.KeyType {
	case pb.DockerImageRemoveReqResp_REF:
		affected, err := client.InspectImageByRef(req.Ref)
		if err != nil {
			glog.Errorf("failed to remove image by ref: %v", err)
			return resp, fmt.Errorf("failed to remove image by ref: %v", err)
		}
		resp.Id = affected.ID
		resp.Id = affected.ID
		fallthrough
	case pb.DockerImageRemoveReqResp_ID:
		affected, err := client.RemoveImage(resp.Id, req.ImageRemoveOptions.ExportIntoDockerApiType())
		if err != nil {
			glog.Errorf("Failed to remove image: %v", err)
			return resp, fmt.Errorf("failed to remove docker image: %v", err)
		}
		resp.ImageDeleteResponseItems = make([]*mobypb.ImageDeleteResponseItem, len(affected))
		for _, v := range affected {
			resp.ImageDeleteResponseItems = append(resp.ImageDeleteResponseItems, mobypb.ConvertFromDockerApiTypeImageDeleteResponseItem(v))
		}
	}
	resp.StateCode = 0
	resp.StateMessage = "removed"
	glog.Infof("have removed image %v", resp.Id)
	return resp, nil
}

func (cli *DockerClient) PruneImages(ctx context.Context, req *pb.DockerImagePruneReqResp) (*pb.DockerImagePruneReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Uable to prune images: %v", err)
		return resp, fmt.Errorf("unable to prune images: %v", err)
	}

	report, err := cli.DockerApiClient.ImagesPrune(ctx, req.Filters.ExportIntoDockerApiType())
	if err != nil {
		glog.Errorf("Failed to prune images: %v", err)
		return resp, fmt.Errorf("failed to prune docker images: %v", err)
	}
	resp.ImagesPruneReport = mobypb.ConvertFromDockerApiTypeImagesPruneReport(report)
	resp.StateCode = 0
	resp.StateMessage = "pruned"
	glog.Infoln("have pruned images")
	return resp, nil
}

func (cli *DockerClient) PullImage(ctx context.Context, req *pb.DockerImagePullReqResp) (*pb.DockerImagePullReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if nil != err {
		glog.Errorf("Unable to pull image: %v", err)
		return resp, fmt.Errorf("Unable to pull docker image: %v", err)
	}
	client := cli.KubeDockerClient

	ch := make(chan map[int]string)
	go func() {
		fmt.Printf("goroutine: pull image: %v\n", req.RefStr)
		// dockershim/libdocker/image.go
		err := libdocker.PullImageAsyncMessaging(client, ctx, req.RefStr, req.ImagePullOptions.ExportIntoDockerApiType(), ch)
		if err != nil {
			glog.V(4).Infof("Pull failure: %v", err)
			return
		}
		glog.Infoln("Stopped")
	}()

	status := <-ch
	for k, v := range status {
		if k >= 100 {
			glog.Errorf("Pulling image failure: %v", v)
			return resp, fmt.Errorf("failed to pull image: %v", v)
		}
	}

	resp.StateCode = IMAGE_PULL_STARTED
	resp.StateMessage = "started"
	glog.Infoln("Docker image being pull")
	return resp, nil
}

func (cli *DockerClient) PushImage(ctx context.Context, req *pb.DockerImagePushReqResp) (*pb.DockerImagePushReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable to push image: %v", err)
		return resp, fmt.Errorf("unable to push docker image: %v", err)
	}
	client := cli.KubeDockerClient

	ch := make(chan map[int]string)
	go func() {
		fmt.Printf("goroutine: push image: %v\n", req.RefStr)
		err := libdocker.PushImageAsyncMessaging(client, ctx, req.RefStr, req.ImagePushOptions.ExportIntoDockerApiType(), ch)
		if err != nil {
			glog.V(4).Infof("Push failure: %v", err)
			return
		}
		glog.Infoln("Stoped")
	}()

	status := <-ch
	for k, v := range status {
		if k >= 100 {
			glog.Errorf("Pushing image failure: %v", v)
			return resp, fmt.Errorf("failed to push image: %v", v)
		}
	}

	resp.StateCode = IMAGE_PUSH_STARTED
	resp.StateMessage = "started"
	glog.Infoln("Docker image being push")
	return resp, nil
}

func (cli *DockerClient) BuildImage(ctx context.Context, req *pb.DockerImageBuildReqResp) (*pb.DockerImageBuildReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Uable to build image: %v", err)
		return resp, fmt.Errorf("unable to build docker image: %v", err)
	}
	client := cli.KubeDockerClient

	buf := bytes.NewBuffer(req.BuildContext)

	ch := make(chan map[int]string)
	go func() {
		fmt.Printf("goroutine: build image: %v\n", req.ImageBuildOptions.Target)
		err := libdocker.BuildImageAsyncMessaging(client, ctx, buf, req.ImageBuildOptions.ExportIntoDockerApiType(), ch)
		if err != nil {
			glog.V(4).Infof("Build image failure: %v", err)
			return
		}
		glog.Infoln("Stopped")
	}()

	status := <-ch
	for k, v := range status {
		if k >= 100 {
			glog.Errorf("Building image failure: %v", v)
			return resp, fmt.Errorf("failed to build docker image: %v", v)
		}
	}

	resp.StateCode = IMAGE_BUILD_STARTED
	resp.StateMessage = "started"
	glog.Infoln("Docker image being build")
	return resp, nil
}
