package server

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/tangfeixiong/go-to-docker/metering/pb"
)

func (m *myExporter) Measure(ctx context.Context, req *pb.MetricReqResp) (*pb.MetricReqResp, error) {
	fmt.Printf("go to make metrics: %q\n", req)
	if req == nil {
		return req, fmt.Errorf("Request name is required")
	}
	if req.Meter != pb.KnownMeter_CADVISOR {
		return req, fmt.Errorf("Not supported currently")
	}
	return new(pb.MetricReqResp), nil
}

func (m *myCollector) Transit(ctx context.Context, req *pb.MetricReqResp) (*pb.MetricReqResp, error) {
	fmt.Printf("go to deal with transitions: %q\n", req)
	if req == nil {
		return req, fmt.Errorf("Request name is required")
	}
	if req.Meter != pb.KnownMeter_CADVISOR {
		return req, fmt.Errorf("Not supported currently")
	}
	return new(pb.MetricReqResp), nil
}

func (m *myCollector) BatchTransit(streamer pb.CollectorService_BatchTransitServer) error {
	return fmt.Errorf("Not implemented")
}
