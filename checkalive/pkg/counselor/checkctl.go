package counselor

import (
	"bytes"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"

	checkalivepb "github.com/tangfeixiong/go-to-docker/checkalive/pb"
	"github.com/tangfeixiong/go-to-docker/checkalive/pkg/manipulate"
)

type WriteCMDBFunc func(*checkalivepb.CheckActionReqResp) error
type WriteMQFunc func(*checkalivepb.CheckActionReqResp) error

type status struct {
	ticker   *time.Ticker
	modified time.Time
	latest   []byte
}

type WebXCheckerConfigMgmt struct {
	ActionReqResp   *checkalivepb.CheckActionReqResp
	name            string
	command         []string
	args            []string
	env             []string
	workdir         string
	periodic        int32
	destinationPath string
	RootPath        string
	ticker          *time.Ticker
	timestamp       time.Time
	result          []byte
	mgmt            map[string]*status
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

	bn := filepath.Base(cm.destinationPath)
	switch bn {
	case "web1check.py", "web2check.py":
		var result []byte
		var err error
		if bn == "web1check.py" {
			result, err = manipulate.Client.Web1Check(filepath.Join(cm.RootPath, cm.workdir))
		} else {
			result, err = manipulate.Client.Web2Check(filepath.Join(cm.RootPath, cm.workdir))
		}
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
	case "web1-2-tag-1-0x20_checker.py":
		opt := append([]string{}, cm.ActionReqResp.Command[1:]...)
		result, err := manipulate.Client.Web1_2_check(cm.workdir, opt...)
		if err != nil {
			fmt.Printf("Failed to execute: %s", err.Error())
		}
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
	default:
		fmt.Println("not implemented")
	}
}

func (cm *WebXCheckerConfigMgmt) ConfigCreation(pkgHome string, actionReq *checkalivepb.CheckActionReqResp) (*checkalivepb.CheckActionReqResp, error) {
	resp := &checkalivepb.CheckActionReqResp{
		Name:               actionReq.Name,
		Periodic:           actionReq.Periodic,
		Selector:           actionReq.Selector,
		DestConfigurations: make(map[string]*checkalivepb.DestinationConfig),
	}
	cm.ActionReqResp = resp
	cm.RootPath = pkgHome
	cm.mgmt = make(map[string]*status)

	for ak, req := range actionReq.DestConfigurations {
		fmt.Println("action key:", ak)
		var checkpath string = ""
		l := len(req.Command)
		if l == 0 {
			return resp, fmt.Errorf("check command required")
		}
		switch {
		case req.Command[0] == "python":
			if 1 != l {
				checkpath = req.Command[l-1]
			} else {
				return resp, fmt.Errorf("Program required, for example: python my.py ...")
			}
		case l != 1:
			return resp, fmt.Errorf("Not implemented of command format: %v", req.Command)
		default:
			checkpath = req.Command[0]
		}
		fmt.Println("action path:", checkpath)

		var found error = errors.New("Stop recursive searching")
		err := filepath.Walk(pkgHome, func(path string, f os.FileInfo, err error) error {
			switch {
			case path == pkgHome:
				return nil
			case strings.HasPrefix(path, filepath.Join(pkgHome, req.WorkDir)):
				if strings.HasSuffix(path, checkpath) {
					fmt.Printf("Visited: %s\n", path)
					req.DestinationPath = path
					return found
				}
			default:
				if filepath.Dir(path) != pkgHome {
					return filepath.SkipDir
				}
			}
			return nil
		})
		fmt.Printf("filepath.Walk() returned %v\n", err)
		if err == nil || err != found {
			glog.Infoln("Command file not found:", checkpath, err)
			return resp, fmt.Errorf("Command file %s not found", checkpath)
		}
		if _, err := os.Stat(req.DestinationPath); os.IsNotExist(err) {
			glog.Infoln("Program file not found:", checkpath, err)
			return resp, fmt.Errorf("Program not found: %s", err.Error())
		}

		for k, v := range req.Conf {
			confpath := filepath.Join(pkgHome, req.WorkDir, k)
			bdec, err := b64.StdEncoding.DecodeString(v)
			if err != nil {
				glog.Infoln("Invalid conf data:", err)
				return resp, fmt.Errorf("Invalid conf data: %s", err.Error())
			}
			if err := ioutil.WriteFile(confpath, bdec, 0644); err != nil {
				glog.Infoln("Failed write conf into file:", err)
				return resp, fmt.Errorf("Could not write conf: %s", err.Error())
			}
			fmt.Println("action config:", k)
		}

		resp.DestConfigurations[ak] = &checkalivepb.DestinationConfig{
			Name:            req.Name,
			Command:         req.Command,
			Args:            req.Args,
			Env:             req.Env,
			Conf:            req.Conf,
			WorkDir:         req.WorkDir,
			Periodic:        req.Periodic,
			DestinationPath: req.DestinationPath,
		}
	}
	return resp, nil
}

func (cm *WebXCheckerConfigMgmt) StartCheck() {
	for dck, req := range cm.ActionReqResp.DestConfigurations {
		periodic := cm.ActionReqResp.Periodic
		if req.Periodic != 0 {
			periodic = req.Periodic
		}
		if periodic <= 0 || periodic >= 3600 {
			fmt.Println("Tick away")
			cm.doCheck(dck)
			return
		}

		ticker := time.NewTicker(time.Second * time.Duration(cm.periodic))
		cm.mgmt[dck] = &status{
			ticker: ticker,
		}
		go func() {
			for t := range ticker.C {
				fmt.Println("Tick at", t)
				cm.doCheck(dck)
			}
		}()
		time.Sleep(time.Millisecond * 500)
	}
}

func (cm *WebXCheckerConfigMgmt) RestartCheck() {
	cm.StopCheck()
	cm.StartCheck()
}

func (cm *WebXCheckerConfigMgmt) StopCheck() {
	for dck, _ := range cm.ActionReqResp.DestConfigurations {
		ticker := cm.mgmt[dck].ticker
		if ticker != nil {
			ticker.Stop()
			fmt.Println("Ticker stopped")
		}
	}
}

func (cm *WebXCheckerConfigMgmt) doCheck(dck string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	workdir := cm.ActionReqResp.DestConfigurations[dck].WorkDir
	destinationPath := cm.ActionReqResp.DestConfigurations[dck].DestinationPath
	timestamp := time.Now()
	req := cm.ActionReqResp.DestConfigurations[dck]

	switch filepath.Base(destinationPath) {
	case "awd1-4-tag-1-0x20_checker.py", "awd1-8-tag-1-0x20_checker.py":
		fallthrough
	case "web1-2-tag-1-0x20_checker.py":
		opts := append(append([]string{}, req.Command[1:]...), req.Args...)
		result, err := manipulate.Client.Python_00_check(workdir, opts...)
		if err != nil {
			fmt.Printf("Failed to execute: %s", err.Error())
		}
		fmt.Println(string(result))

		buf := bytes.Buffer{}
		if len(cm.mgmt[dck].latest) != 0 {
			buf.Write(cm.mgmt[dck].latest)
			buf.WriteByte('\n')
		}
		buf.Write(result)
		cm.ActionReqResp.DestConfigurations[dck].Timestamp = timestamp.Format(time.RFC3339)
		cm.ActionReqResp.DestConfigurations[dck].StateMessage = buf.String()
		cm.mgmt[dck].latest = result
		cm.mgmt[dck].modified = timestamp

		if cm.WriteCMDBFn != nil {
			cm.WriteCMDBFn(cm.ActionReqResp)
		}
		if cm.WriteMQFn != nil {
			cm.WriteMQFn(cm.ActionReqResp)
		}
	default:
		fmt.Println("not implemented")
	}
}
