package exporter

import (
	//	"bufio"
	//	"bytes"
	//	b64 "encoding/base64"
	//	"errors"
	"fmt"
	//	"io/ioutil"
	//	"math"
	//	"os"
	//	"path/filepath"
	//	"regexp"
	"strings"
	"sync"
	"time"

	//	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/metering/pb"
)

type MeterDispatcher interface {
	RPC(CollectorClientRPC) MeterDispatcher
	StartMetering() map[string]Job
	ReapMetrics(url string) (*pb.MeteringReqResp, error)
	StopMetering()
}

type CollectorClientRPC interface {
	Transit(*pb.MeteringReqResp) (*pb.MeteringReqResp, error)
}

type ExporterManager struct {
	MeteringNameURLs    map[string][]string
	MetricsCollectorRPC string
	Dispatchers         map[string]MeterDispatcher
	name                string
	command             []string
	args                []string
	env                 []string
	conf                map[string]string
	workdir             string
	periodic            int32
	duration            int32
	destinationPath     string
	RootPath            string
	ticker              *time.Ticker
	timestamp           time.Time
	result              []byte
	mutex               sync.Mutex
}

func (em *ExporterManager) Dispatch(ch <-chan bool) {
	for k, v := range em.MeteringNameURLs {
		switch strings.Trim(strings.ToLower(k), " ") {
		case "cadvisor":
			dispatcher := NewCAdvisorManager(v)
			em.Dispatchers["cadvisor"] = dispatcher
			if em.MetricsCollectorRPC != "" {
				dispatcher.RPC(NewClient_gRPC(em.MetricsCollectorRPC)).StartMetering()
			} else {
				dispatcher.StartMetering()
			}
		default:
			fmt.Println("Not implemented", k)
		}
	}
	if len(em.Dispatchers) != 0 {
		<-ch
		for _, v := range em.Dispatchers {
			v.StopMetering()
		}
	}
}
