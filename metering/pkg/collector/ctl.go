package collector

import (
	"fmt"
	//	"bufio"
	//	"bytes"
	"encoding/json"
	//	"errors"
	//	"fmt"
	//	"io/ioutil"
	//	"math"
	//	"os"
	//	"path/filepath"
	//	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	info "github.com/google/cadvisor/info/v1"

	"github.com/tangfeixiong/go-to-docker/metering/pb"
	"github.com/tangfeixiong/go-to-docker/metering/pkg/cache/memory"
	"github.com/tangfeixiong/go-to-docker/metering/pkg/storage/elasticsearch"
)

type CollectorManager struct {
	MetricsStorageDriver       string
	MetricsStorageDuration     time.Duration
	drivername                 string
	storagehost                string
	elasticsearchIndexName     string
	elsaticsearchTypeName      string
	elasticsearchEnableSniffer bool
	chmetrics                  chan *pb.MeteringReqResp
	storageCache               *memory.InMemoryCache
	command                    []string
	args                       []string
	ticker                     *time.Ticker
	timestamp                  time.Time
	result                     []byte
	mutex                      sync.Mutex
}

func (cm *CollectorManager) Start(ch <-chan bool) {
	cm.chmetrics = make(chan *pb.MeteringReqResp, 100)

	s := strings.Split(cm.MetricsStorageDriver, "=")
	cm.drivername = s[0]
	if len(s) > 1 {
		cm.storagehost = s[1]
	}
	if cm.drivername == "elasticsearch" {
		fmt.Printf("%q", s)
		cm.elasticsearchIndexName = "cadvisor"
		cm.elsaticsearchTypeName = "stats"
		cm.elasticsearchEnableSniffer = false
		elasticsearch.Register(&cm.storagehost, &cm.elasticsearchIndexName, &cm.elasticsearchIndexName, &cm.elasticsearchEnableSniffer)
	}
	memoryStorage, err := NewMemoryStorage(&cm.drivername, &cm.MetricsStorageDuration)
	if err != nil {
		glog.Infof("Failed to initialize storage driver: %s", err)
	}
	cm.storageCache = memoryStorage

	ticker := time.NewTicker(cm.MetricsStorageDuration)
	cm.ticker = ticker
	go func() {
		for t := range ticker.C {
			fmt.Println("Tiker at", t)
			in := <-cm.chmetrics
			for k, v := range in.MeterResponse {
				if k == int32(pb.MetricType_CADVISOR_V1_CONTAINERINFO) {
					var containers []*info.ContainerInfo
					if err := json.Unmarshal(v, &containers); err != nil {
						glog.Infoln("Failed to read metrics, error:", err.Error())
					} else {
						for _, container := range containers {
							ref := container.ContainerReference
							for _, stats := range container.Stats {
								if err := cm.storageCache.AddStats(ref, stats); err != nil {
									glog.Infoln("Failed to store metrics, error:", err.Error())
								}
							}
						}
					}
				}
			}

		}
	}()
}

func (cm *CollectorManager) Store(in *pb.MeteringReqResp) (*pb.MeteringReqResp, error) {
	resp := new(pb.MeteringReqResp)
	if in == nil {
		resp.StateCode = 100
		resp.StateMessage = "Data not found"
		return resp, fmt.Errorf("Data not found")
	}
	if in.MeterDriver != pb.MeterDriver_CADVISOR {
		resp.StateCode = 200
		resp.StateMessage = "Not implemeted"
		return resp, fmt.Errorf("Not implemented")
	}
	cm.chmetrics <- in
	return in, nil
}
