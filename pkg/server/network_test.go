package server

import (
	"testing"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pb/moby"
)

func TestDocker_network_ls(t *testing.T) {
	req := &pb.DockerNetworkData{
		NetworkResources: []*moby.NetworkResource{
			{
				Name: "bridge",
			},
		},
	}

	s := newServer()

	resp, err := s.reapDockerNetworking(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%q", resp)
}
