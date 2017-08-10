package server

import (
	"testing"

	"github.com/tangfeixiong/go-to-docker/pb"
)

func TestLbr_NetworkLandscape(t *testing.T) {
	req := &pb.BridgedNetworkingData{
		LinuxBridges: []*pb.BridgedNetworkingData_LinuxBridgeInfo{
			{
				Name: "docker0",
			},
		},
	}

	s := newServer()

	resp, err := s.reapBridgedNetworkLandscape(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%q", resp)
}
