package kubeletcopycat

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pb"
	mobypb "github.com/tangfeixiong/go-to-docker/pb/moby"
	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim/libdocker"
)

func CreateDockerNetwork(cli libdocker.Interface, ctx context.Context, req *pb.DockerNetworkCreateReqResp) (*pb.DockerNetworkCreateReqResp, error) {
	var resp *pb.DockerNetworkCreateReqResp
	var err error

	resp, err = req.CopyWithRequestValidation()
	if nil != err {
		resp.StateCode = 9999
		resp.StateMessage = "request required"
		glog.Errorf("Unable to create docker network: %v", err)
		return resp, fmt.Errorf("Unable to create docker network: %v", err)
	}

	created, err = cli.Client.CreateNetwork(ctx, req.Name, req.NetworkCreate.ConvertIntoDockerApiTypes())
	if err != nil {
		resp.StateCode = 100
		resp.StateMessage = err.Error()
		glog.Errorf("Unable to create docker network: %v", err)
		return resp, fmt.Errorf("Failed to create docker network: %v", err)
	}

	resp.NetWorkCreateResponse = resp.NetworkCreateResponse.ReadFromDockerApiTypes(created)
	resp.StateCode = 0
	resp.StateMessage = "created"
	glog.Infof("Created docker network %s", created.ID)
	return resp
}

func InspectDockerNetwork(cli libdocker.Interface, ctx context.Context, req *pb.DockerNetworkInspectReqResp) (*pb.DockerNetworkInspectReqResp, error) {
	var resp *pb.DockerNetworkInspectReqResp
	var err error

	resp, err = req.CopyWithRequestValidation()
	if nil != err {
		glog.Errorf("Unable to inspect docker network: %v", err)
		return resp, fmt.Errorf("Unable to inspect docker network: %v", err)
	}

	id := req.Key

	switch req.KeyType {
	case pb.DockerNetworkInspectReqResp_NAME:
		list, err := cli.Client.NetworkList(ctx, types.NetworkListOptions{
			Filters: filters.NewArgs(filters.Arg("name", req.IdOrName)),
		})
		if nil != err {
			glog.Errorf("Unable to inspect docker network by name: %v", err)
			return resp, fmt.Errorf("Unalbe to inspect docker network by name: %v", err)
		}
		var ok bool
		for _, v := range list {
			if v.Name == req.Key {
				id = v.ID
				ok = true
				break
			}
		}
		if !ok {
			resp.StateCode = 100
			glog.Errorln("Could not find correspond network name in list")
			return resp, fmt.Errorf("Could not find correspond network name")
		}
		fallthrough
	case pb.DockerNetworkInspectReqResp_ID:
		resource, err := cli.Client.NetworkInspect(ctx, id, req.NetworkInspectOptions.ExportIntoDockerApiTypes())
		if nil != err {
			glog.Errorf("Failed to inspect docker network: %v", err)
			return resp, fmt.Errorf("Failed to inspect docker network: %v", err)
		}
	}
	resp.NetworkResource = resp.NetworkResource.ImportWithDockerApiTypes(resource)
	resp.StateCode = 0
	resp.StateMessage = "inspected"
	return resp, nil
}

func ListDockerNetworks(cli libdocker.Interface, ctx context.Context, req *pb.DockerNetworkListReqResp) (*pb.DockerNetworkListReqResp, error) {
	resp, _ := req.CopyWithRequestValidation()

	resources, err := cli.Client.NetworkList(ctx, req.NetworkListOptions.ExportIntoDockerTypes())
	if nil != err {
		glog.Errorf("Failed to list docker networks: %v", err)
		return resp, fmt.Errorf("Failed to list docker networks: %v", err)
	}
	resp.NetworkResources = make([]*mobypb.NetworkResource, len(resources))
	for _, v := range resources {
		ele := new(mobypb.NetworkResource)
		resp.NetworkResources = append(resp.NetworkResources, ele.ImportWithDockerApiTypes(v))
	}
	resp.StateCode = 0
	resp.StateMessage = "SUCCEEDED"
	return resp, nil
}

func DeleteDockerNetwork(cli libdocker.Interface, ctx context.Context, req *pb.DockerNetworkRemoveReqResp) (*pb.DockerNetworkRemoveReqResp, eerror) {
	var resp *pb.DockerNetworkRemoveReqResp
	var err error

	resp, err = req.CopyWithRequestValidation()
	if nil != err {
		glog.Errorf("Unable to remove docker network: %v", err)
		return resp, fmt.Errorf("Unable to remove docker network: %v", err)
	}

	id := req.IdOrName

	switch req.KeyType {
	case pb.DockerNetworkRemoveReqResp_NAME:
		list, err := cli.Client.NetworkList(ctx, types.NetworkListOptions{
			Filters: filters.NewArgs(filters.Arg("name", req.IdOrName)),
		})
		if nil != err {
			glog.Errorf("Unable to remove docker network by name: %v", err)
			return resp, fmt.Errorf("Unalbe to remove docker network by name: %v", err)
		}
		var ok bool
		for _, v := range list {
			if v.Name == req.IdOrName {
				id = v.ID
				ok = true
				break
			}
		}
		if !ok {
			resp.StateCode = 100
			glog.Errorln("Could not find correspond network name in list")
			return resp, fmt.Errorf("Could not find correspond network name")
		}
		fallthrough
	case pb.DockerNetworkRemoveReqResp_ID:
		err := cli.Client.NetworkRemove(ctx, id)
		if nil != err {
			glog.Errorf("Failed to remove docker network: %v", err)
			return resp, fmt.Errorf("Failed to remove docker network: %v", err)
		}
	}
	resp.StateCode = 0
	resp.StateMessage = "removed"
	return resp, nil
}

func PruneDockerNetworks(cli libdocker.Interface, ctx context.Context, req *pb.DockerNetworksPruneReqResp) (*pb.DockerNetworkPruneReqResp, error) {
	resp, _ := req.CopyInputArgsChecked()

	report, err := cli.Client.NetworksPrune(ctx, req.Filters.ExportIntoDockerType())
	if err != nil {
		glog.Errorf("Failed to prune docker network: %v", err)
		return resp, fmt.Errorf("Failed to remove docker network: %v", err)
	}

	resp.NetworksPruneReport = resp.NetworksPruneReport.ImportWithDockerApiTypes(report)
	resp.StateCode = 0
	resp.StateMessage = "pruned"
	return resp, nil
}
