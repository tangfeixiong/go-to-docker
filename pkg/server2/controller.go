package server2

import (
	"context"
	"errors"

	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pb"
)

var (
	errNotImplemented = errors.New("Not Implemented")
)

// docker container op

func (ms *MicroServer) RunDockerContainer(ctx context.Context, req *pb.DockerContainerRunReqResp) (*pb.DockerContainerRunReqResp, error) {
	glog.V(4).Infof("go to run docker container %v", req)
	return ms.DockerClient.RunContainer(ctx, req)
}

func (ms *MicroServer) CreateContainer(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) StartContainer(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) ListDockerContainers(ctx context.Context, req *pb.DockerContainerListReqResp) (*pb.DockerContainerListReqResp, error) {
	glog.V(4).Infof("go to list docker containers %v", req)
	return ms.DockerClient.ListContainers(ctx, req)
}

func (ms *MicroServer) InspectDockerContainer(ctx context.Context, req *pb.DockerContainerInspectReqResp) (*pb.DockerContainerInspectReqResp, error) {
	glog.V(4).Infof("go to inspect docker container %v", req)
	return ms.DockerClient.InspectContainer(ctx, req)
}

func (ms *MicroServer) InspectContainerWithSize(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) ContainerLogs(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) UpdateContainerResources(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) StopContainer(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) RemoveDockerContainer(ctx context.Context, req *pb.DockerContainerRemoveReqResp) (*pb.DockerContainerRemoveReqResp, error) {
	glog.V(4).Infof("go to remove docker container %v", req)
	return ms.DockerClient.RemoveContainer(ctx, req)
}

func (ms *MicroServer) PruneDockerContainers(ctx context.Context, req *pb.DockerContainerPruneReqResp) (*pb.DockerContainerPruneReqResp, error) {
	glog.V(4).Infof("go to prune docker containers")
	return ms.DockerClient.PruneContainers(ctx, req)
}

// docker image op

func (ms *MicroServer) ListDockerImages(ctx context.Context, req *pb.DockerImageListReqResp) (*pb.DockerImageListReqResp, error) {
	glog.V(4).Infof("go to list docker images: %v", req)
	return ms.DockerClient.ListImages(ctx, req)
}

func (ms *MicroServer) InspectDockerImage(ctx context.Context, req *pb.DockerImageInspectReqResp) (*pb.DockerImageInspectReqResp, error) {
	glog.V(4).Infof("go to inspect docker image: %v", req)
	return ms.DockerClient.InspectImage(ctx, req)
}

func (ms *MicroServer) InspectImageByRef(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) InspectImageById(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) RemoveDockerImage(ctx context.Context, req *pb.DockerImageRemoveReqResp) (*pb.DockerImageRemoveReqResp, error) {
	glog.V(4).Infof("go to remove docker image: %v", req)
	return ms.DockerClient.RemoveImage(ctx, req)
}

func (ms *MicroServer) PruneDockerImages(ctx context.Context, req *pb.DockerImagePruneReqResp) (*pb.DockerImagePruneReqResp, error) {
	glog.V(4).Infof("go to prune docker image: %v", req)
	return ms.DockerClient.PruneImages(ctx, req)
}

func (ms *MicroServer) PullDockerImage(ctx context.Context, req *pb.DockerImagePullReqResp) (*pb.DockerImagePullReqResp, error) {
	glog.V(4).Infof("go to pull docker image: %v", req)
	return ms.DockerClient.PullImage(ctx, req)
}

func (ms *MicroServer) PushDockerImage(ctx context.Context, req *pb.DockerImagePushReqResp) (*pb.DockerImagePushReqResp, error) {
	glog.V(4).Infof("go to push docker image: %v", req)
	return ms.DockerClient.PushImage(ctx, req)
}

func (ms *MicroServer) BuildDockerImage(ctx context.Context, req *pb.DockerImageBuildReqResp) (*pb.DockerImageBuildReqResp, error) {
	glog.V(4).Infof("go to build docker image: %v", req)
	return ms.DockerClient.BuildImage(ctx, req)
}

func (ms *MicroServer) ImageHistory(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) PullDockerImageStreaming(req *pb.DockerImagePullReqResp, streaming pb.SimpleService_PullDockerImageStreamingServer) error {
	return errNotImplemented
}

func (ms *MicroServer) BuildDockerImageStreaming(streaming pb.SimpleService_BuildDockerImageStreamingServer) error {
	return errNotImplemented
}

func (ms *MicroServer) PushDockerImageStreaming(streaming pb.SimpleService_PushDockerImageStreamingServer) error {
	return errNotImplemented
}

// docker network op

func (ms *MicroServer) CreateDockerNetwork(ctx context.Context, req *pb.DockerNetworkCreateReqResp) (*pb.DockerNetworkCreateReqResp, error) {
	glog.V(4).Infof("go to create docker network: %v", req)
	return ms.DockerClient.CreateNetwork(ctx, req)
}

func (ms *MicroServer) ListDockerNetworks(ctx context.Context, req *pb.DockerNetworkListReqResp) (*pb.DockerNetworkListReqResp, error) {
	glog.V(4).Infof("go to list docker networks: %v", req)
	return ms.DockerClient.ListNetworks(ctx, req)
}

func (ms *MicroServer) InspectDockerNetwork(ctx context.Context, req *pb.DockerNetworkInspectReqResp) (*pb.DockerNetworkInspectReqResp, error) {
	glog.V(4).Infof("go to inspect docker network: %v", req)
	return ms.DockerClient.InspectNetwork(ctx, req)
}

func (ms *MicroServer) ConnectDockerNetwork(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) DisconnectDockerNetwork(ctx context.Context) error {
	return errNotImplemented
}

func (ms *MicroServer) RemoveDockerNetwork(ctx context.Context, req *pb.DockerNetworkRemoveReqResp) (*pb.DockerNetworkRemoveReqResp, error) {
	glog.V(4).Infof("go to remove docker network: %v", req)
	return ms.DockerClient.RemoveNetwork(ctx, req)
}

func (ms *MicroServer) PruneDockerNetworks(ctx context.Context, req *pb.DockerNetworkPruneReqResp) (*pb.DockerNetworkPruneReqResp, error) {
	glog.V(4).Infof("go to prune docker networks: %v", req)
	return ms.DockerClient.PruneNetworks(ctx, req)
}
