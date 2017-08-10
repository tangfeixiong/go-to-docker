package server

import (
	// "fmt"

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
