package counselor

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	healthcheckerpb "github.com/tangfeixiong/go-to-docker/checkalive/pb"
	"github.com/tangfeixiong/go-to-docker/checkalive/pkg/manipulate"
)

type WriteCMDBFunc func(*healthcheckerpb.CheckActionReqResp) error
type WriteMQFunc func(*healthcheckerpb.CheckActionReqResp) error

type WebXCheckerConfigMgmt struct {
	ActionReqResp   *healthcheckerpb.CheckActionReqResp
	name            string
	command         []string
	args            []string
	env             map[string]string
	periodic        int32
	RootPath        string
	workdir         string
	destinationPath string
	hosts           []string
	result          []byte
	timestamp       time.Time
	ticker          *time.Ticker
	mutex           sync.Mutex
	WriteCMDBFn     WriteCMDBFunc
	WriteMQFn       WriteMQFunc
}

func (cm *WebXCheckerConfigMgmt) CreateTicker() {
	if cm.ActionReqResp != nil {
		cm.periodic = cm.ActionReqResp.Periodic
	}
	if cm.periodic <= 0 || cm.periodic >= 3600 {
		fmt.Println("Tick out")
		cm.checkout()
		return
	}

	cm.ticker = time.NewTicker(time.Second * time.Duration(cm.periodic))
	ticker := cm.ticker
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			cm.checkout()
		}
	}()
	time.Sleep(time.Millisecond * 500)
}

func (cm *WebXCheckerConfigMgmt) UpdateTicker() {
	cm.DestroyTicker()
	cm.CreateTicker()
}

func (cm *WebXCheckerConfigMgmt) DestroyTicker() {
	ticker := cm.ticker
	if ticker != nil {
		ticker.Stop()
		fmt.Println("Ticker stopped")
	}
}

func (cm *WebXCheckerConfigMgmt) checkout() {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.timestamp = time.Now()
	if cm.ActionReqResp != nil {
		cm.workdir = cm.ActionReqResp.WorkDir
		cm.destinationPath = cm.ActionReqResp.DestinationPath
		cm.ActionReqResp.Timestamp = cm.timestamp.Format(time.RFC3339)
	}

	switch {
	case strings.HasSuffix(cm.destinationPath, "web1check.py"):
		result, err := manipulate.Client.Web1Check(filepath.Join(cm.RootPath, cm.workdir))
		if err != nil {
			fmt.Printf("Failed to execute: %s", err.Error())
		}
		fmt.Println(string(result))
		if cm.ActionReqResp != nil {
			buf := bytes.Buffer{}
			if len(cm.result) != 0 {
				buf.Write(cm.result)
				buf.WriteByte('\n')
			}
			buf.Write(result)
			cm.ActionReqResp.StateMessage = buf.String()
			if cm.WriteCMDBFn != nil {
				cm.WriteCMDBFn(cm.ActionReqResp)
			}
			if cm.WriteMQFn != nil {
				cm.WriteMQFn(cm.ActionReqResp)
			}
		}
		cm.result = result
	case strings.HasPrefix(cm.destinationPath, "web2check.py"):
		result, err := manipulate.Client.Web2Check(filepath.Join(cm.RootPath, cm.workdir))
		if err != nil {
			fmt.Printf("Failed to execute: %s", err.Error())
		}
		fmt.Println(string(result))
		cm.result = result
	default:
		fmt.Println("not implemented")
	}
}
