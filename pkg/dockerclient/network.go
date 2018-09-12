package dockerclient

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/golang/glog"

	utilrand "k8s.io/apimachinery/pkg/util/rand"

	"github.com/tangfeixiong/go-to-docker/pb"
	mobypb "github.com/tangfeixiong/go-to-docker/pb/moby"
)

func (cli *DockerClient) CreateNetwork(ctx context.Context, req *pb.DockerNetworkCreateReqResp) (*pb.DockerNetworkCreateReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable creating network: %v", err)
		return resp, fmt.Errorf("unable creating network: %v", err)
	}
	if len(req.Name) == 0 {
		req.Name = fmt.Sprintf("stackdocker-br%s", utilrand.String(9))
	}

	created, err := cli.DockerApiClient.NetworkCreate(ctx, req.Name, req.NetworkCreate.ExportIntoDockerApiType())
	if err != nil {
		resp.StateCode = 100
		resp.StateMessage = err.Error()
		glog.Errorf("Failed to create network: %v", err)
		return resp, fmt.Errorf("failed to create docker network: %v", err)
	}
	resp.Name = req.Name
	resp.NetworkCreateResponse = mobypb.ConvertFromDockerApiTypeNetworkCreateResponse(created)
	resp.StateCode = 0
	resp.StateMessage = "created"
	glog.Infof("docker network %s created", resp.NetworkCreateResponse.Id)
	return resp, nil
}

func (cli *DockerClient) ListNetworks(ctx context.Context, req *pb.DockerNetworkListReqResp) (*pb.DockerNetworkListReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable to inspect network: %v", err)
		return resp, fmt.Errorf("unable to inspect docker network: %v", err)
	}

	resources, err := cli.DockerApiClient.NetworkList(ctx, req.NetworkListOptions.ExportIntoDockerApiType())
	if nil != err {
		glog.Errorf("Failed to list docker networks: %v", err)
		return resp, fmt.Errorf("failed to list docker networks: %v", err)
	}
	resp.NetworkResources = make([]*mobypb.NetworkResource, len(resources))
	for _, v := range resources {
		resp.NetworkResources = append(resp.NetworkResources, mobypb.ConvertFromDockerApiTypeNetworkResource(v))
	}
	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	glog.Infoln("docker network listed")
	return resp, nil
}

func (cli *DockerClient) InspectNetwork(ctx context.Context, req *pb.DockerNetworkInspectReqResp) (*pb.DockerNetworkInspectReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if nil != err {
		glog.Errorf("Unable to inspect docker network: %v", err)
		return resp, fmt.Errorf("unable to inspect docker network: %v", err)
	}

	switch req.KeyType {
	case pb.DockerNetworkInspectReqResp_NAME:
		cond := filters.NewArgs()
		cond.Add("name", req.Name)
		affected, err := cli.DockerApiClient.NetworkList(ctx, types.NetworkListOptions{
			Filters: cond,
		})
		if nil != err {
			glog.Errorf("Unable to inspect docker network by name: %v", err)
			return resp, fmt.Errorf("unable to inspect docker network by name: %v", err)
		}
		var ok bool
		for _, v := range affected {
			if v.Name == req.Name {
				req.Id = v.ID
				resp.Id = v.ID
				ok = true
				break
			}
		}
		if !ok {
			resp.StateCode = 100
			glog.Errorln("Could not find correspond network name in list")
			return resp, fmt.Errorf("could not find correspond network name %s", req.Name)
		}
		fallthrough
	case pb.DockerNetworkInspectReqResp_ID:
		resource, err := cli.DockerApiClient.NetworkInspect(ctx, resp.Id, req.NetworkInspectOptions.ExportIntoDockerApiType())
		if nil != err {
			glog.Errorf("Failed to inspect docker network: %v", err)
			return resp, fmt.Errorf("Failed to inspect docker network: %v", err)
		}
		resp.NetworkResource = mobypb.ConvertFromDockerApiTypeNetworkResource(resource)
	}
	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	glog.Infof("docker network %s inspected", resp.Id)
	return resp, nil
}

func (cli *DockerClient) RemoveNetwork(ctx context.Context, req *pb.DockerNetworkRemoveReqResp) (*pb.DockerNetworkRemoveReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if nil != err {
		glog.Errorf("Unable to remove docker network: %v", err)
		return resp, fmt.Errorf("unable to remove docker network: %v", err)
	}

	switch req.KeyType {
	case pb.DockerNetworkRemoveReqResp_NAME:
		cond := filters.NewArgs()
		cond.Add("name", req.Name)
		affected, err := cli.DockerApiClient.NetworkList(ctx, types.NetworkListOptions{
			Filters: cond,
		})
		if nil != err {
			glog.Errorf("Unable to remove docker network by name: %v", err)
			return resp, fmt.Errorf("unalbe to remove docker network by name: %v", err)
		}
		var ok bool
		for _, v := range affected {
			if v.Name == req.Name {
				req.Id = v.ID
				resp.Id = v.ID
				ok = true
				break
			}
		}
		if !ok {
			resp.StateCode = 100
			glog.Errorln("Could not find correspond network name in list")
			return resp, fmt.Errorf("could not find correspond network name %s", req.Name)
		}
		fallthrough
	case pb.DockerNetworkRemoveReqResp_ID:
		err := cli.DockerApiClient.NetworkRemove(ctx, resp.Id)
		if nil != err {
			glog.Errorf("Failed to remove docker network: %v", err)
			return resp, fmt.Errorf("Failed to remove docker network: %v", err)
		}
	}
	resp.StateCode = 0
	resp.StateMessage = "removed"
	return resp, nil
}

func (cli *DockerClient) PruneNetworks(ctx context.Context, req *pb.DockerNetworkPruneReqResp) (*pb.DockerNetworkPruneReqResp, error) {
	resp, err := req.DeepCopyCheckedArgs()
	if err != nil {
		glog.Errorf("Unable to prune networks: %v", err)
		return resp, fmt.Errorf("unable to prune networks: %v", err)
	}

	report, err := cli.DockerApiClient.NetworksPrune(ctx, req.Filters.ExportIntoDockerApiType())
	if err != nil {
		glog.Errorf("Failed to prune docker network: %v", err)
		return resp, fmt.Errorf("failed to remove docker network: %v", err)
	}
	resp.NetworksPruneReport = mobypb.ConvertFromDockerApiTypeNetworksPruneReport(report)
	resp.StateCode = 0
	resp.StateMessage = "pruned"
	return resp, nil
}
