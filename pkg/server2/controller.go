package server2

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim/libdocker"
)

const (
	ILLEGAL_PARAMETER = 1000
	DOCKER_NOT_READY  = 1001

	IMAGE_PULL_ACCEPTED = 1
)

var (
	ErrNotImplemented = errors.New("Not Implemented")
)

func (ms *MicroServer) PullImage(ctx context.Context, req *pb.DockerImagePullReqResp) (*pb.DockerImagePullReqResp, error) {
	var resp *pb.DockerImagePullReqResp
	var err error

	resp, err = req.CopyWithRequestValidation()
	if nil != err {
		return resp, fmt.Errorf("Unable to pull image: %v", err)
	}

	ch := make(chan map[int]error)
	go func() {
		fmt.Printf("rpc request of PullImage: %v\n", req.ImageRef)
		// dockershim/libdocker/image.go
		ms.DockerClient.PullImageIntoCachingProgress(ctx, req.RefStr, req.ImagePullOptions, ch)
	}()

	status := <-ch
	for k, v := range status {
		if k != 0 {
			glog.Errorf("Pulling image failure: %v", v)
			resp.StateCode = 100
			resp.StateMessage = v.Error()
			return resp, fmt.Errorf("Pulling image failure: %v", v)
		}
	}

	resp.StateCode = IMAGE_PULL_ACCEPTED
	resp.StateMessage = "Request pulling image accepted"
	return resp, nil
}

func (ms *MicroServer) InspectImageByRef(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) InspectImageById(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) ListImages(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) RemoveImage(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) BuildImage(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) ImageHistory(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) CreateNetwork(ctx context.Context, req *pb.DockerNetworkCreateReqResp) (*pb.DockerNetworkCreateReqResp, error) {
	var resp *pb.DockerNetworkCreateReqResp
	var err error

	resp, err = req.CopyWithRequestValidation()
	if nil != err {
		return resp, fmt.Errorf("Unable to create network: %v", err)
	}

	resp.NetWorkCreateResponse, err = ms.DockerClient.CreateNetwork(ctx, req.Name, req.NetworkCreate)
	if err != nil {
		resp.StateCode = 100
		resp.StateMessage = err.Error()
		return resp, fmt.Errorf("Creating network failure: %v", err)
	}

	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	return resp
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

	return ErrNotImplemented
}

func (ms *MicroServer) CreateContainer(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) StartContainer(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) ContainerLogs(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) StopContainer(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) ListContainers(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) InspectContainer(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) InspectContainerWithSize(ctx context.Context) error {

	return ErrNotImplemented
}

func (ms *MicroServer) UpdateContainerResources(ctx context.Context) error {

	return ErrNotImplemented
}
