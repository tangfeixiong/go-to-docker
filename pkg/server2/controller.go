package server2

import (
	"context"
	"fmt"

	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pkg/dockerclient"
)

const (
	ILLEGAL_PARAMETER = 1000
	DOCKER_NOT_READY  = 1001

	IMAGE_TO_PULL = 1
)

func (ms *MicroServer) PullImage(ctx context.Context, req *pb.DockerImagePullReqResp) (*pb.DockerImagePullReqResp, error) {
	resp := new(pb.DockerImagePullReqResp)
	if req == nil || 0 == len(req.ImageRef) {
		glog.Warningln("Image ref required")
		return resp, fmt.Errorf("Image ref reqired")
	}

	resp.ImageRef = req.ImageRef

	ch := make(chan map[int]error)
	go func() {
		fmt.Printf("rpc request of PullImage: %v\n", req.ImageRef)
		// dockerclient/image.go
		ms.dockerClient.PullImage(req.ImageRef, ch)
	}()

	status := <-ch
	for k, v := range status {
		if k != 0 {
			glog.Warningf("Failed to pull image: %v", v)
			return resp, fmt.Errorf("Failed to pull image: %v", v)
		}
	}

	resp.StateCode = IMAGE_TO_PULL
	resp.StateMessage = "Request of pulling image accepted"
	return resp, nil
}

func (ms *MicroServer) InspectImageByRef(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) InspectImageById(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) ListImages(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) RemoveImage(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) BuildImage(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) ImageHistory(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) CreateNetwork(ctx context.Context, req *pb.DockerNetworkCreateReqResp) (*pb.DockerNetworkCreateReqResp, error) {
	resp := new(pb.DockerNetworkCreateReqResp)
	if req == nil {
		glog.Warningln("Network request required")
		return nil, fmt.Errorf("Network request required")
	}

	return resp, nil
}

func (ms *MicroServer) RemoveNetwork(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) InspectNetwork(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) ListNetworks(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) RunContainer(ctx context.Context, req *pb.DockerContainerRunReqResp) (*pb.DockerContainerRunReqResp, error) {
	resp := new(pb.DockerContainerRunReqResp)
	if req == null {
		glog.Warningln("Container request required")
		return nil, fmt.Errorf("Container request required")
	}

	return resp, nil
}

func (ms *MicroServer) RemoveContainer(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) CreateContainer(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) StartContainer(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) ContainerLogs(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) StopContainer(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) ListContainers(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) InspectContainer(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) InspectContainerWithSize(ctx context.Context) error {

	return nil
}

func (ms *MicroServer) UpdateContainerResources(ctx context.Context) error {

	return nil
}
