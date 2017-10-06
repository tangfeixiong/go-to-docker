package exporter

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/google/cadvisor/client"
	"github.com/google/cadvisor/info/v1"

	"github.com/tangfeixiong/go-to-docker/metering/pb"
)

type ActionType int

const (
	ACTION_START ActionType = iota
	ACTION_PAUSE
	ACTION_CONTINUE
	ACTION_STOP
)

type Job struct {
	Id          *ActionType
	action      chan ActionType
	result      *pb.MeteringReqResp
	Err         error
	LastUpdated time.Time
}

type cAdvisorManager struct {
	jobs      map[string]*Job
	periodic  time.Duration
	ticker    *time.Ticker
	clientrpc CollectorClientRPC
}

func NewCAdvisorManager(urls []string) *cAdvisorManager {
	cm := &cAdvisorManager{
		jobs:     make(map[string]*Job),
		periodic: time.Second * 5,
	}
	for _, url := range urls {
		cm.jobs[url] = &Job{
			action: make(chan ActionType),
		}
	}
	return cm
}

func (cm *cAdvisorManager) RPC(clientrpc CollectorClientRPC) MeterDispatcher {
	cm.clientrpc = clientrpc
	return cm
}

func (cm *cAdvisorManager) StartMetering() map[string]Job {
	resp := make(map[string]Job)
	wg := sync.WaitGroup{}
	cm.ticker = time.NewTicker(cm.periodic)
	ticker := cm.ticker
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			for k, v := range cm.jobs {
				wg.Add(1)
				go func() {
					defer wg.Done()
					v.result, v.Err = cm.ReapMetrics(k)
					if cm.clientrpc != nil {
						cm.clientrpc.Transit(v.result)
					}
					v.LastUpdated = time.Now()
					resp[k] = *v
				}()
			}
		}
	}()
	wg.Wait()
	return resp
}

func (cm *cAdvisorManager) StopMetering() {
	cm.ticker.Stop()
}

func (cm *cAdvisorManager) ControlStarting(url string) error {
	v, ok := cm.jobs[url]
	if !ok {
		fmt.Println("Invalid metering host,", url)
		return fmt.Errorf("Invalid metering host,", url)
	}
	v.action <- ACTION_START
	return nil
}

func (cm *cAdvisorManager) ReapMetrics(url string) (*pb.MeteringReqResp, error) {
	resp := &pb.MeteringReqResp{
		MeterDriver:      pb.MeterDriver_CADVISOR,
		MeterUrl:         url,
		TimestampNanosec: time.Now().UnixNano(),
		MeterResponse:    make(map[int32][]byte),
	}
	c, _ := client.NewClient(url)
	mi, err := c.MachineInfo()
	if err != nil {
		glog.Info("Could not get machine info, error:", err.Error())
		resp.StateCode = 100
		resp.StateMessage = err.Error()
		return resp, err
	}
	b, err := json.Marshal(mi)
	if err != nil {
		resp.StateCode = 200
		resp.StateMessage = err.Error()
		return resp, err
	}
	resp.MeterResponse[int32(pb.MetricType_CADVISOR_V1_MACHINEINFO)] = b

	query := v1.DefaultContainerInfoRequest()
	adc, err := c.AllDockerContainers(&query)
	if err != nil {
		glog.Info("Could not get all docker containers, error:", err.Error())
		resp.StateCode = 101
		resp.StateMessage = err.Error()
		return resp, err
	}
	b, err = json.Marshal(adc)
	if err != nil {
		resp.StateCode = 201
		resp.StateMessage = err.Error()
		return resp, err
	}
	resp.MeterResponse[int32(pb.MetricType_CADVISOR_V1_CONTAINERINFO)] = b

	// dc, err := c.DockerContainer(name, &query)
	einfo, err := c.EventStaticInfo("?oom_events=true")
	if err != nil {
		glog.Errorf("got error retrieving event info: %v", err)
		resp.StateCode = 102
		resp.StateMessage = err.Error()
		return resp, err
	}
	for idx, ev := range einfo {
		glog.Infof("static einfo %v: %v", idx, ev)
	}

	b, err = json.Marshal(einfo)
	if err != nil {
		resp.StateCode = 202
		resp.StateMessage = err.Error()
		return resp, err
	}
	resp.MeterResponse[int32(pb.MetricType_CADVISOR_V1_EVENT)] = b

	fmt.Printf("%q", resp)
	return resp, nil
}
