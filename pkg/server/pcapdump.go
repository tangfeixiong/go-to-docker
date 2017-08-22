package server

import (
	"fmt"
	// "time"

	"github.com/tangfeixiong/go-to-docker/pb"
	// "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/gopacketctl"
)

func (m *myService) sniffEtherNetworking(req *pb.EthernetSniffingData) (*pb.EthernetSniffingData, error) {
	resp := new(pb.EthernetSniffingData)

	if nil == req || "" == req.Iface {
		resp.StateCode = 10
		resp.StateMessage = "Request required"
		return resp, fmt.Errorf("Request required")
	}

	//	content, err := gopacketctl.PcapdumpOnce(req.Iface, time.Second*3)
	//	if nil != err {
	//		resp.StateCode = 100
	//		resp.StateMessage = err.Error()
	//		return resp, err
	//	}
	//	resp.StatsAndPackets = content
	//	return resp, err

	resp.StateCode = 999
	resp.StateMessage = "Not running with docker container"
	return resp, fmt.Errorf(resp.StateMessage)
}
