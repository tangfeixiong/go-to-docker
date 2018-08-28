package libdocker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/golang/glog"

	mobypb "github.com/tangfeixiong/go-to-docker/pb/moby"
)

func (dc *DockerClient) CreateNetwork(ctx context.Context, name string, opt *mobypb.NetworkCreate) (*mobypb.NetworkCreateResponse, error) {
	resp := new(mobypb.NetworkCreateResponse)
	result, err := dc.MobyClient.NetworkCreate(ctx, name, opt.ConvertIntoDockerApiTypes())
	if nil != err {
		glog.Errorf("Failed to create Docker network: %v", err)
		return resp, fmt.Errorf("Creating network failure: %v", err)
	}
	return resp.ConvertFromDockerApiTypes(result), nil
}
