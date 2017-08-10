package server

import (
	"testing"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pb/moby"
)

func TestDocker_Container_inspect(t *testing.T) {
	req := &pb.DockerContainerInspection{
		ContainerInfo: &moby.ContainerJSON{
			ContainerJsonBase: &moby.ContainerJSONBase{
				Id: "57258341ab92",
			},
		},
	}

	s := newServer()

	resp, err := s.inspectContainer(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%q", resp)
}
