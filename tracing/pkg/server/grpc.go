package server

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/tangfeixiong/go-to-docker/tracing/pb"
	// "github.com/tangfeixiong/go-to-docker/tracing/pkg/"
)

func (ma *myAgent) Sample(ctx context.Context, req *pb.TracingReqResp) (*pb.TracingReqResp, error) {
	fmt.Printf("go to sample: %q\n", req)
	if req == nil {
		return req, fmt.Errorf("Request is required")
	}
	resp := new(pb.TracingReqResp)
	return resp, nil
}

func (mc *myCollector) Sample(ctx context.Context, req *pb.TracingReqResp) (*pb.TracingReqResp, error) {
	fmt.Printf("go to sample: %q\n", req)
	if req == nil {
		return req, fmt.Errorf("Request is required")
	}
	resp := new(pb.TracingReqResp)
	return resp, nil
}
