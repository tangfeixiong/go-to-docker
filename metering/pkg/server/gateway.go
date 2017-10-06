package server

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/tangfeixiong/go-to-docker/metering/pb"
	"github.com/tangfeixiong/go-to-docker/metering/pkg/exporter"
)

func (m *myExporter) Measure(ctx context.Context, req *pb.MeteringReqResp) (*pb.MeteringReqResp, error) {
	fmt.Printf("go to measure: %q\n", req)
	if req == nil {
		return req, fmt.Errorf("Request name is required")
	}
	if req.MeterDriver != pb.MeterDriver_CADVISOR {
		return req, fmt.Errorf("Not supported currently")
	}
	if req.MeterUrl == "" {
		return req, fmt.Errorf("Meter not available")
	}
	c := exporter.NewCAdvisorManager([]string{req.MeterUrl})
	return c.ReapMetrics(req.MeterUrl)
}

func (m *myCollector) Transit(ctx context.Context, req *pb.MeteringReqResp) (*pb.MeteringReqResp, error) {
	fmt.Printf("go to deal with transitions: %q\n", req)
	if req == nil {
		return req, fmt.Errorf("Request name is required")
	}
	if req.MeterDriver != pb.MeterDriver_CADVISOR {
		return req, fmt.Errorf("Not supported currently")
	}
	if req.MeterUrl == "" {
		return req, fmt.Errorf("Meter not available")
	}
	return m.collectormanager.Store(req)
}

func (m *myCollector) BatchTransit(streamer pb.CollectorService_BatchTransitServer) error {
	return fmt.Errorf("Not implemented")
}
