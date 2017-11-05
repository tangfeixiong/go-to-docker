package counselor

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	// "reflect"
	"regexp"
	// "runtime"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"

	checkalivepb "github.com/tangfeixiong/go-to-docker/checkalive/pb"
	"github.com/tangfeixiong/go-to-docker/checkalive/pkg/manipulate"
)

type WriteCMDBFunc func(*checkalivepb.CheckActionReqResp) error
type WriteMQFunc func(*checkalivepb.CheckActionReqResp) error

type Message struct {
	modified time.Time
	latest   []byte
}

type CheckerController struct {
	ActionReqResp   *checkalivepb.CheckActionReqResp
	name            string
	command         []string
	args            []string
	env             []string
	conf            map[string]string
	workdir         string
	periodic        int32
	duration        int32
	tplconf         map[string]*checkalivepb.DestinationConfig
	destinationPath string
	RootPath        string
	ticker          *time.Ticker
	timestamp       time.Time
	result          []byte
	messages        map[string][]Message
	mutex           sync.Mutex
	WriteCMDBFn     WriteCMDBFunc
	WriteMQFn       WriteMQFunc
}

func (cm *CheckerController) CreateTicker() {
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

func (cm *CheckerController) UpdateTicker() {
	cm.DestroyTicker()
	cm.CreateTicker()
}

func (cm *CheckerController) DestroyTicker() {
	ticker := cm.ticker
	if ticker != nil {
		ticker.Stop()
		fmt.Println("Ticker stopped")
	}
}

func (cm *CheckerController) checkout() {
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

func (cc *CheckerController) Dispatch(pkgHome string, req *checkalivepb.CheckActionReqResp) (*checkalivepb.CheckActionReqResp, error) {
	resp := &checkalivepb.CheckActionReqResp{
		Name:               req.Name,
		Periodic:           req.Periodic,
		Duration:           req.Duration,
		DestConfigurations: make(map[string]*checkalivepb.DestinationConfig),
	}

	l := len(req.Command)
	if l == 0 {
		return resp, fmt.Errorf("Command must be required")
	}
	var checkpath string = ""
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

	var found error = errors.New("Stop recursive searching")
	err := filepath.Walk(pkgHome, func(path string, f os.FileInfo, err error) error {
		switch {
		case path == pkgHome:
			break
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
		resp.StateCode = 11
		return resp, fmt.Errorf("Command file %s not found", checkpath)
	}
	if _, err := os.Stat(req.DestinationPath); os.IsNotExist(err) {
		glog.Infoln("Program file not found:", checkpath, err)
		resp.StateCode = 11
		return resp, fmt.Errorf("Program not found: %s", err.Error())
	}
	fmt.Println("command:", checkpath)
	resp.DestinationPath = req.DestinationPath

	for ck, _ := range req.DestConfigurations {
		req.DestConfigurations[ck].Args = make([]string, len(req.Args))
		copy(req.DestConfigurations[ck].Args, req.Args)
		req.DestConfigurations[ck].Env = make([]string, len(req.Env))
		copy(req.DestConfigurations[ck].Env, req.Env)
		req.DestConfigurations[ck].Conf = make(map[string]string, len(req.Conf))
	}
	for k, v := range req.Conf {
		fmt.Println("config file:", k)
		bdec, err := b64.StdEncoding.DecodeString(v)
		if err != nil {
			glog.Infoln("Invalid conf data:", err)
			resp.StateCode = 20
			return resp, fmt.Errorf("Invalid conf data: %s", err.Error())
		}
		for ck, _ := range req.DestConfigurations {
			req.DestConfigurations[ck].Conf[k] = string(bdec)
		}
	}
	resp.Command = req.Command
	resp.Args = req.Args
	resp.Env = req.Env
	resp.Conf = req.Conf
	resp.DestConfigurations = req.DestConfigurations

	for ck, cv := range req.DestConfigurations {
		for tk, tv := range cv.Tpl {
			r := regexp.MustCompile(`\$\(` + tk + `\)`)
			for i, v := range req.DestConfigurations[ck].Args {
				req.DestConfigurations[ck].Args[i] = r.ReplaceAllString(v, tv)
			}
			for i, v := range req.DestConfigurations[ck].Env {
				req.DestConfigurations[ck].Env[i] = r.ReplaceAllString(v, tv)
			}
			for k, v := range req.DestConfigurations[ck].Conf {
				req.DestConfigurations[ck].Conf[k] = r.ReplaceAllString(v, tv)
			}
		}
		fmt.Printf("Dest config of %s: %v", ck, req.DestConfigurations[ck])
	}

	for _, cv := range req.DestConfigurations {
		for k, v := range cv.Conf {
			confpath := filepath.Join(pkgHome, req.WorkDir, k)
			if err := ioutil.WriteFile(confpath, []byte(v), 0644); err != nil {
				glog.Infoln("Failed to write conf into file; error:", err)
				resp.StateCode = 10
				return resp, fmt.Errorf("Could not write conf: %s", err.Error())
			}
		}
	}

	cc.messages = make(map[string][]Message)
	for ck, _ := range req.DestConfigurations {
		cc.messages[ck] = make([]Message, 0)
	}
	cc.ActionReqResp = req
	cc.RootPath = pkgHome
	return resp, nil
}

func (cc *CheckerController) Start() {
	periodic := cc.ActionReqResp.Periodic
	if periodic < 1 {
		return
	}
	duration := float64(cc.ActionReqResp.Duration)
	if duration <= 0 {
		duration = math.MaxFloat64
	}

	ticker := time.NewTicker(time.Second * time.Duration(periodic))
	cc.ticker = ticker

	go func() {
		var count int64 = 0
		ts := time.Now()
		for t := range ticker.C {
			elapsed := time.Since(ts).Seconds()
			amount := 0
			fmt.Println(cc.ActionReqResp)
			for k, _ := range cc.ActionReqResp.DestConfigurations {
				switch {
				case cc.ActionReqResp.Periodic <= 0:
					if count == 0 {
						fmt.Println("Run once:", k)
						go func() {
							cc.doCheck(k)
						}()
					}
					amount += 1
				default:
					if elapsed < duration+1 {
						fmt.Println("Tick at", t, "-> key:", k)
						go func() {
							cc.doCheck(k)
						}()
					} else {
						amount += 1
					}
				}
			}
			count++
			if amount == len(cc.ActionReqResp.DestConfigurations) {
				break
			}
		}
		// runtime.Goexit()
		ticker.Stop()
		fmt.Println("Stopped")
	}()

	time.Sleep(time.Millisecond * 100)
}

func (cc *CheckerController) Refresh(req *checkalivepb.CheckActionReqResp) (*checkalivepb.CheckActionReqResp, error) {
	//	resp := &checkalivepb.CheckActionReqResp{
	//		Name:     req.Name,
	//		Periodic: req.Periodic,
	//		Duration: req.Duration,
	//	}
	//	resp.Command = req.Command
	//	resp.Args = req.Args
	//	resp.Env = req.Env
	//	resp.Conf = req.Conf
	//	resp.DestConfigurations = req.DestConfigurations
	if cc.ActionReqResp.Periodic == req.Periodic && cc.ActionReqResp.Duration == req.Duration {
		println("Identical papameters")
		return cc.ActionReqResp, nil
	}
	cc.Stop()
	cc.ActionReqResp.Periodic = req.Periodic
	cc.ActionReqResp.Duration = req.Duration
	defer cc.Start()
	return cc.ActionReqResp, nil
}

func (cc *CheckerController) Restart() {
	cc.Stop()
	cc.Start()
}

func (cc *CheckerController) Stop() {
	ticker := cc.ticker
	if ticker != nil {
		ticker.Stop()
		fmt.Println("Ticker stopped")
	}
}

func (cc *CheckerController) doCheck(key string) {
	cc.mutex.Lock()
	defer cc.mutex.Unlock()

	destinationPath := cc.ActionReqResp.DestinationPath
	workdir := filepath.Join(cc.RootPath, cc.ActionReqResp.WorkDir)
	timestamp := time.Now()
	value := cc.ActionReqResp.DestConfigurations[key]

	var result []byte
	var err error
	prev := cc.ActionReqResp.DestConfigurations[key].StateCode
	switch filepath.Base(destinationPath) {
	case "awd1-4-checker.py", "awd1-8-checker.py":
		fallthrough
	case "awd1_lemon_cms_check.py", "awd2_dynpage_check.py", "awd4_tomcat_check.py", "awd5_gracer_check.py":
		fallthrough
	case "awd6_cms_check.py", "awd7_upload_check.py", "awd8_blog_check.py", "awd9_money_check.py":
		fallthrough
	case "awd10_nothing_check.py", "awd11_maccms_check.py", "awd12_phpsqllitecms_check.py":
		fallthrough
	case "awd2_daydayweb_check.py", "awd4_chinaz_check.py", "awd5_babyblog_check.py":
		opts := append(append([]string{}, cc.ActionReqResp.Command[1:]...), value.Args...)
		result, err = manipulate.Client.Python_00_check(workdir, opts...)
		if err != nil {
			glog.Infof("Failed to execute: %s", err.Error())
			return
		}
		scanner := bufio.NewScanner(bytes.NewReader(result))
		for scanner.Scan() {
			if ok, _ := regexp.MatchString(`.*\[Errno.+\] Connection refused.*`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString(`\('.*', 'error:', '.*'\)`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString(`\(False, '.*'\)`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString(`\(True, '.*'\)`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode = 0
				break
			}
			fmt.Println("Unkown content:", scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Faild to reading standard input:", err)
		}
	case "awd3_electronics_check.py", "awd1_xmanweb2_check.py":
		opts := append(append([]string{}, cc.ActionReqResp.Command[1:]...), value.Args...)
		result, err = manipulate.Client.Python_00_check(workdir, opts...)
		if err != nil {
			glog.Infof("Failed to execute: %s", err.Error())
			return
		}
		scanner := bufio.NewScanner(bytes.NewReader(result))
		for scanner.Scan() {
			if ok, _ := regexp.MatchString(`.*\[Errno.+\] Connection refused.*`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString(`\('.*', 'error:', '.*'\)`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString(`{'status': *'down', *'msg':.*}`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString(`{'status': *'up', *'msg':.*}`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode = 0
				break
			}
			fmt.Println("Unkown content:", scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Faild to reading standard input:", err)
		}
	case "awd3_shadow_check.py":
		opts := append(append([]string{}, cc.ActionReqResp.Command[1:]...), value.Args...)
		result, err = manipulate.Client.Python_00_check(workdir, opts...)
		if err != nil {
			glog.Infof("Failed to execute: %s", err.Error())
			return
		}
		scanner := bufio.NewScanner(bytes.NewReader(result))
		for scanner.Scan() {
			if ok, _ := regexp.MatchString(`.*\[Errno.+\] Connection refused.*`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString(`\('.*', 'error:', '.*'\)`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString(`\[no\].*`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString(`\[ok\].*`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode = 0
				break
			}
			fmt.Println("Unkown content:", scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Faild to reading standard input:", err)
		}
	case "web1-checker.py", "web2-checker.py":
		opts := append(append([]string{}, cc.ActionReqResp.Command[1:]...), value.Args...)
		result, err = manipulate.Client.Python_00_check(workdir, opts...)
		if err != nil {
			glog.Infof("Failed to execute: %s", err.Error())
			return
		}
		scanner := bufio.NewScanner(bytes.NewReader(result))
		for scanner.Scan() {
			if ok, _ := regexp.MatchString("-+", scanner.Text()); ok {
				continue
			}
			if ok, _ := regexp.MatchString("checking host:.+", scanner.Text()); ok {
				continue
			}
			if ok, _ := regexp.MatchString(`.*\[Errno.+\] Connection refused.*`, scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString("Host:.+seems down", scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode += 1
				break
			}
			if ok, _ := regexp.MatchString("Host:.+seems ok", scanner.Text()); ok {
				cc.ActionReqResp.DestConfigurations[key].StateCode = 0
				break
			}
			fmt.Println("Unkown content:", scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Faild to reading standard input:", err)
		}
		// fallthrough
	default:
		switch {
		case cc.ActionReqResp.Command[0] == "python":
			opts := append(append([]string{}, cc.ActionReqResp.Command[1:]...), value.Args...)
			result, err = manipulate.Client.Python_00_check(workdir, opts...)
			if err != nil {
				glog.Infof("Failed to execute: %s", err.Error())
				return
			}
			scanner := bufio.NewScanner(bytes.NewReader(result))
			for scanner.Scan() {
				if ok, _ := regexp.MatchString(`\(False, '.*'\)`, scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode += 1
					break
				}
				if ok, _ := regexp.MatchString(`\(True, '.*'\)`, scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode = 0
					break
				}
				if ok, _ := regexp.MatchString(`{'status': *'down', *'msg':.*}`, scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode += 1
					break
				}
				if ok, _ := regexp.MatchString(`{'status': *'up', *'msg':.*}`, scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode = 0
					break
				}
				if ok, _ := regexp.MatchString(`\[no\].*`, scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode += 1
					break
				}
				if ok, _ := regexp.MatchString(`\[ok\].*`, scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode = 0
					break
				}
				if ok, _ := regexp.MatchString("Host:.+seems down", scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode += 1
					break
				}
				if ok, _ := regexp.MatchString("Host:.+seems ok", scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode = 0
					break
				}
				// runtime error
				if ok, _ := regexp.MatchString(`.*\[Errno.+\] Connection refused.*`, scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode += 1
					break
				}
				if ok, _ := regexp.MatchString(`\('.*', 'error:', '.*'\)`, scanner.Text()); ok {
					cc.ActionReqResp.DestConfigurations[key].StateCode += 1
					break
				}
				// treat as information
				if ok, _ := regexp.MatchString("-+", scanner.Text()); ok {
					continue
				}
				if ok, _ := regexp.MatchString("checking host:.+", scanner.Text()); ok {
					continue
				}
				fmt.Println("Unkown content:", scanner.Text())
			}
			if err = scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Faild to reading standard input:", err)
			}
		default:
			fmt.Println("Not implemented bin:", destinationPath)
			return
		}
	}
	fmt.Println(string(result))

	cc.messages[key] = append(cc.messages[key], Message{timestamp, result})
	if len(cc.messages[key]) > 3 {
		cc.messages[key] = cc.messages[key][1:]
	}

	var buf bytes.Buffer
	for _, msg := range cc.messages[key] {
		// buf.WriteString(msg.modified.Format(time.RFC3339))
		// buf.WriteByte('\n')
		buf.Write(msg.latest)
	}
	buf.WriteByte('\n')
	cc.ActionReqResp.DestConfigurations[key].StateMessage = buf.String()
	cc.ActionReqResp.DestConfigurations[key].Timestamp = timestamp.Format(time.RFC3339)

	if cc.WriteCMDBFn != nil {
		cc.WriteCMDBFn(cc.ActionReqResp)
	}
	if cc.WriteMQFn != nil {
		if prev != cc.ActionReqResp.DestConfigurations[key].StateCode {
			cc.WriteMQFn(cc.ActionReqResp)
		}
	}
}
