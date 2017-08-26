package server

import (
	"fmt"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/network"
	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pb/moby"
	"github.com/tangfeixiong/go-to-docker/pkg/dockerctl"
)

func (m *myService) reapDockerNetworking(req *pb.DockerNetworkData) (*pb.DockerNetworkData, error) {
	resp := new(pb.DockerNetworkData)
	resp.NetworkResources = make([]*moby.NetworkResource, 0)

	ctl := dockerctl.NewEngine1_12Client()

	resultbrs, err := ctl.ListNetwork()
	if nil != err {
		return resp, err
	}

	for _, br := range resultbrs {
		resp.NetworkResources = append(resp.NetworkResources, &moby.NetworkResource{
			Name:       br.Name,
			Id:         br.ID,
			Scope:      br.Scope,
			Driver:     br.Driver,
			EnableIpv6: br.EnableIPv6,
			Ipam: &moby.IPAM{
				Driver:  br.IPAM.Driver,
				Options: br.IPAM.Options,
				Config:  make([]*moby.IPAMConfig, len(br.IPAM.Config)),
			},
			Internal:   br.Internal,
			Containers: make(map[string]*moby.EndpointResource, len(br.Containers)),
			Options:    br.Options,
			Labels:     br.Labels,
		})
		l := len(resp.NetworkResources)
		for i := 0; i < len(br.IPAM.Config); i++ {
			resp.NetworkResources[l-1].Ipam.Config[i] = &moby.IPAMConfig{
				Subnet:     br.IPAM.Config[i].Subnet,
				IpRange:    br.IPAM.Config[i].IPRange,
				Gateway:    br.IPAM.Config[i].Gateway,
				AuxAddress: br.IPAM.Config[i].AuxAddress,
			}
		}
		for k, v := range br.Containers {
			resp.NetworkResources[l-1].Containers[k] = &moby.EndpointResource{
				Name:        v.Name,
				EndpointId:  v.EndpointID,
				MacAddress:  v.MacAddress,
				Ipv4Address: v.IPv4Address,
				Ipv6Address: v.IPv6Address,
			}
		}
	}

	return resp, nil
}

func (m *myService) createDockerNetwork(req *pb.DockerNetworkCreationReqResp) (*pb.DockerNetworkCreationReqResp, error) {
	resp := new(pb.DockerNetworkCreationReqResp)
	if nil == req.NetworkCreateRequest {
		return resp, fmt.Errorf("Request required")
	}
	if nil == req.NetworkCreateRequest.NetworkCreate {
		return resp, fmt.Errorf("Network Creation required")
	}
	if nil == req.NetworkCreateRequest.NetworkCreate.Ipam {
		return resp, fmt.Errorf("IPAM required")
	}
	if 0 == len(req.NetworkCreateRequest.NetworkCreate.Ipam.Config) {
		return resp, fmt.Errorf("IPAM Config required")
	}
	if 0 == len(req.NetworkCreateRequest.Name) {
		return resp, fmt.Errorf("Network Name required")
	}
	resp.NetworkCreateRequest = req.NetworkCreateRequest

	ctl := dockerctl.NewEngine1_12Client()

	creation := types.NetworkCreate{
		CheckDuplicate: req.NetworkCreateRequest.NetworkCreate.CheckDuplicate,
		Driver:         req.NetworkCreateRequest.NetworkCreate.Driver,
		EnableIPv6:     req.NetworkCreateRequest.NetworkCreate.EnableIpv6,
		IPAM: network.IPAM{
			Driver:  req.NetworkCreateRequest.NetworkCreate.Ipam.Driver,
			Options: req.NetworkCreateRequest.NetworkCreate.Ipam.Options,
			Config:  make([]network.IPAMConfig, len(req.NetworkCreateRequest.NetworkCreate.Ipam.Config)),
		},
		Internal: req.NetworkCreateRequest.NetworkCreate.Internal,
		Options:  req.NetworkCreateRequest.NetworkCreate.Options,
		Labels:   req.NetworkCreateRequest.NetworkCreate.Labels,
	}
	i := 0
	for _, v := range req.NetworkCreateRequest.NetworkCreate.Ipam.Config {
		creation.IPAM.Config[i] = network.IPAMConfig{
			Subnet:     v.Subnet,
			IPRange:    v.IpRange,
			Gateway:    v.Gateway,
			AuxAddress: v.AuxAddress,
		}
		i++
	}

	ncresp, err := ctl.CreateNetwork(req.NetworkCreateRequest.Name, creation)
	if err != nil {
		resp.StateCode = 1
		resp.StateMessage = fmt.Sprintf("Could not perform creation, Error: %s", err.Error())
		glog.Info(resp.StateMessage)
		return resp, err
	}
	resp.NetworkCreateResponse = &moby.NetworkCreateResponse{
		Id:      ncresp.ID,
		Warning: ncresp.Warning,
	}

	return resp, nil
}
