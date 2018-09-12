package kubeletcopycat

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

func InspectDockerImage(cli libdocker.Interface, ctx context.Context, req *pb.DockerImageInspectReqResp) (*pb.DockerImageInspectReqResp, error) {
	checked, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable to inspect image: %v", err)
		return checked, fmt.Errorv("Unable to inspect image: %v", err)
	}

	resp := checked
	switch req.KeyType {
	case pb.DockerImageInspectReqResp_REF:
		affected, err := cli.InspectImageByRef(req.Key)
		if err != nil {
			glog.Errorf("failed to inspect image by ref: %v", err)
			return resp, fmt.Errorf("failed to inspect image: %v", err)
		}
		resp.ImageInspect = mobypb.ConvertFromDockerApiTypeImageInspect(affected)
		break
	case pb.DockerImageInspectReqResp_ID:
		affected, err := cli.InspectImageByID(req.Key)
		if err != nil {
			glog.Errorf("failed to inspect image: %v", err)
			return resp, fmt.Errorf("failed to inspect image: %v", err)
		}
		resp.ImageInspect = mobypb.ConvertFromDockerApiTypeImageInspect(affected)
		break
	}
	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	glog.Infof("inspected")
	return resp, nil
}

func ListDockerImages(cli libdocker.Interface, ctx context.Context, req *pb.DockerImageListReqResp) (*pb.DockerImageListReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("unable to list images: %v", err)
		return resp, fmt.Errorf("unable to list docker images: %v", err)
	}

	affected, err := cli.ListImages(req.ImageListOptions.ExportIntoDockerApiType())
	if err != nil {
		glog.Errorf("Failed to list images: %v", err)
		return resp, fmt.Errorf("failed to list docker images: %v", err)
	}
	for _, v := range affected {
		resp.ImageSummaries = append(resp.ImageSummaries, mobypb.ConvertFromDockerApiTypeImageSummary(v))
	}
	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	glog.Infoln("listed")
	return resp, nil
}

func RemoveDockerImage(cli libdocker.Interface, ctx context.Context, req *pb.DockerImageRemoveReqResp) (*pb.DockerImageRemoveReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.ErrorF("Unable to remove image: %v", err)
		return resp, fmt.Errorf("unable to remove docker image: %v", err)
	}

	affected, err := cli.RemoveImage(req.Key, req.ImageRemoveOptions.ExportIntoDockerApiType())
	if err != nil {
		glog.Errorf("Failed to remove image: %v", err)
		return resp, fmt.Errorf("failed to remove docker image: %v", err)
	}
	for _, v := range affected {
		resp.ImageDeleteResponseItems = append(resp.ImageDeleteResponseItems, mobypb.ConvertFromDockerApiTypeImageDeleteResponseItem(v))
	}
	resp.StateCode = 0
	resp.StateMessage = "removed"
	glog.Infof("have removed image %v", req.Key)
	return resp, nil
}

func PruneDockerImages(cli libdocker.Interface, ctx context.Context, req *pb.DockerImagePruneReqResp) (*pb.DockerImagePruneReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Uable to prune images: %v", err)
		return resp, fmt.Errorf("unable to prune images: %v", err)
	}

	report, err := libdocker.PruneImages(cli, ctx, req.Filters.ExportIntoDockerApiType())
	if err != nil {
		glog.Errorf("Failed to prune images: %v", err)
		return resp, fmt.Errorf("failed to prune docker images: %v", err)
	}
	resp.ImagePruneReport = mobypb.ConvertFromDockerApiTypeImagePruneReport(report)
	resp.StateCode = 0
	resp.StateMessage = "pruned"
	glog.Infoln("have pruned images")
	return resp, nil
}

func PullDockerImage(cli libdocker.Interface, ctx context.Context, req *pb.DockerImagePullReqResp) (*pb.DockerImagePullReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if nil != err {
		glog.Errorf("Unable to pull image: %v", err)
		return resp, fmt.Errorf("Unable to pull docker image: %v", err)
	}

	ch := make(chan map[int]string)
	go func() {
		fmt.Printf("goroutine: pull image: %v\n", req.RefStr)
		// dockershim/libdocker/image.go
		err := libdocker.PullImageAsyncMessaging(cli, ctx, req.RefStr, req.ImagePullOptions.ExportIntoDockerApiType(), ch)
		if err != nil {
			glog.V(4).Infof("Pull failure: %v", err)
			return
		}
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

func PushDockerImage(cli libdocker.Interface, ctx context.Context, req *pb.DockerImagePushReqResp) (*pb.DockerImagePushReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable to push image: %v", err)
		return resp, fmt.Errorf("unable to push docker image: %v", err)
	}

	ch := make(chan map[int]string)
	go func() {
		fmt.Printf("goroutine: push image: %v\n", req.RefStr)
		err := libdocker.PushImageAsyncMessaging(cli, ctx, req.RefStr, req.ImagePushOptions.ExportIntoDockerApiType(), ch)
		if err != nil {
			glog.V(4).Infof("Push failure: %v", err)
			return
		}
		glog.Infof("Stoped pulling image: %s", msg)
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

func BuildDockerImage(cli libdocker.Interface, ctx context.Context, req *pb.DockerImageBuildReqResp) (*pb.DockerImageBuildReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Uable to build image: %v", err)
		return resp, fmt.Errorf("unable to build docker image: %v", err)
	}

	buf := bytes.NewBuffer(req.Context)

	ch := make(chan map[int]error)
	go func() {
		fmt.Printf("goroutine: build image: %v\n", req.ImageBuildOptions.Target)
		err := libdocker.BuildImageAsyncMessaging(cli, ctx, buf, req.ImageBuildOptions.ExportIntoDockerApiType(), ch)
		if err != nil {
			glog.V(4).Infof("Build image failure: %v", err)
			return
		}
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
