package exporter

/*
import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strconv"

    "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/garyburd/redigo/redis"
	"github.com/nats-io/go-nats"

)

func (m *myServer) createCheck(ctx context.Context, req *pb.CheckActionReqResp) (*pb.CheckActionReqResp, error) {
	if _, ok := m.checkermanager[req.Name]; ok {
		return req, fmt.Errorf("Dispatcher exists, delete or update first. name=%s", req.Name)
	}
	if req.Periodic < 1 {
		return req, fmt.Errorf("Periodic must be greater than one second, name: %s", req.Name)
	}
	if req.Duration < 1 {
		return req, fmt.Errorf("Total duration must be greater than one second, name: %s", req.Name)
	}

	checkctl := new(counselor.CheckerController)
	resp, err := checkctl.Dispatch(m.packgesHome, req)
	if err != nil {
		return resp, err
	}

LOOPC:
	for _, priority := range m.priorityCMDB {
		switch priority {
		case "sentinel":
			break
		case "redis":
			if len(m.redisAddresses) > 0 {
				checkctl.WriteCMDBFn = m.writeRedis
				break LOOPC
			}
			break
		case "etcd":
			if m.etcdAddresses != "" {
				checkctl.WriteCMDBFn = m.writeEtcdV3
				break LOOPC
			}
			break
		case "mysql":
			break
		}
	}

LOOPM:
	for _, priority := range m.priorityMQ {
		switch priority {
		case "sentinel":
			break
		case "redis":
			if len(m.redisAddresses) > 0 {
				checkctl.WriteMQFn = m.publishRedis
				break LOOPM
			}
			break
		case "gnatsd":
			if m.gnatsdAddresses != "" {
				checkctl.WriteMQFn = m.publishGnatsd
				break LOOPM
			}
			break
		case "kafka":
			break
		case "rabbit":
			break
		}
	}

	defer checkctl.Start()

	m.checkermanager[req.Name] = checkctl
	return resp, nil
}

func (m *myServer) CreateLegacy(ctx context.Context, req *pb.CheckActionReqResp) (*pb.CheckActionReqResp, error) {
	resp := new(pb.CheckActionReqResp)
	if _, ok := m.checkermanager[req.Name]; ok {
		return resp, fmt.Errorf("CM of %s exists, delete or update first", req.Name)
	}

	var checkpath string = "web1check.py"
	l := len(req.Command)
	if l > 0 {
		switch {
		case req.Command[0] == "python":
			if 1 != l {
				checkpath = req.Command[l-1]
			} else {
				return resp, fmt.Errorf("Programe required, for example: python my.py ...")
			}
		case l == 1:
			checkpath = req.Command[0]
		default:
			return resp, fmt.Errorf("Not implemented")
		}
	}
	glog.Infoln("path:", checkpath)

	var found error = errors.New("Stop recursive searching")
	err := filepath.Walk(m.packgesHome, func(path string, f os.FileInfo, err error) error {
		switch {
		case path == m.packgesHome:
			return nil
		case strings.HasPrefix(path, filepath.Join(m.packgesHome, req.WorkDir)):
			if strings.HasSuffix(path, checkpath) {
				fmt.Printf("Visited: %s\n", path)
				req.DestinationPath = path
				return found
			}
		default:
			if filepath.Dir(path) != m.packgesHome {
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

	println("config file")
	for k, v := range req.Conf {
		confpath := filepath.Join(m.packgesHome, req.WorkDir, k)
		bdec, err := b64.StdEncoding.DecodeString(v)
		if err != nil {
			glog.Infoln("Invalid conf data:", err)
			return resp, fmt.Errorf("Invalid conf data: %s", err.Error())
		}
		if err := ioutil.WriteFile(confpath, bdec, 0644); err != nil {
			glog.Infoln("Failed write conf into file:", err)
			return resp, fmt.Errorf("Could not write conf: %s", err.Error())
		}
	}

	resp.Name = req.Name
	resp.Command = req.Command
	resp.Args = req.Args
	resp.Env = req.Env
	resp.Conf = req.Conf
	resp.WorkDir = req.WorkDir
	resp.Periodic = req.Periodic
	resp.DestinationPath = req.DestinationPath

	cm := new(counselor.CheckerController)
	cm.ActionReqResp = resp
	cm.RootPath = m.packgesHome

LOOPC:
	for _, priority := range m.priorityCMDB {
		switch priority {
		case "sentinel":
			break
		case "redis":
			if len(m.redisAddresses) > 0 {
				cm.WriteCMDBFn = m.writeRedis
				break LOOPC
			}
			break
		case "etcd":
			if m.etcdAddresses != "" {
				cm.WriteCMDBFn = m.writeEtcdV3
				break LOOPC
			}
			break
		case "mysql":
			break
		}
	}

LOOPM:
	for _, priority := range m.priorityMQ {
		switch priority {
		case "sentinel":
			break
		case "redis":
			if len(m.redisAddresses) > 0 {
				cm.WriteMQFn = m.publishRedis
				break LOOPM
			}
			break
		case "gnatsd":
			if m.gnatsdAddresses != "" {
				cm.WriteMQFn = m.publishGnatsd
				break LOOPM
			}
			break
		case "kafka":
			break
		case "rabbit":
			break
		}
	}

	m.checkermanager[resp.Name] = cm
	cm.CreateTicker()

	return resp, nil
}

func (m *myServer) UpdateWebXCheck(ctx context.Context, req *pb.CheckActionReqResp) (*pb.CheckActionReqResp, error) {
	println("go to update check")
	resp := new(pb.CheckActionReqResp)
	if req == nil || req.Name == "" {
		return resp, fmt.Errorf("Request required")
	}
	if _, ok := m.checkermanager[req.Name]; !ok {
		return resp, fmt.Errorf("CM of %s does not exists, create first", req.Name)
	}

	println("workdir")
	checkpath := m.checkermanager[req.Name].ActionReqResp.DestinationPath
	workdir := filepath.Join(m.packgesHome, req.WorkDir)
	if !strings.HasPrefix(checkpath, workdir) {
		glog.Infoln("Mismatched workdir:", workdir)
		return resp, fmt.Errorf("Command file %s mismatch workdir %s", checkpath, workdir)
	}

	println("config file")
	for k, v := range req.Conf {
		confpath := filepath.Join(m.packgesHome, req.WorkDir, k)
		bdec, err := b64.StdEncoding.DecodeString(v)
		if err != nil {
			return resp, fmt.Errorf("Invalid conf data: %s", err.Error())
		}
		if err := ioutil.WriteFile(confpath, bdec, 0644); err != nil {
			return resp, fmt.Errorf("Could not write conf: %s", err.Error())
		}
	}

	resp.Name = req.Name
	// m.checkermanager[req.Name].ActionReqResp.Command = req.Command
	m.checkermanager[req.Name].ActionReqResp.Args = req.Args
	m.checkermanager[req.Name].ActionReqResp.Env = req.Env
	m.checkermanager[req.Name].ActionReqResp.Conf = req.Conf
	m.checkermanager[req.Name].ActionReqResp.WorkDir = req.WorkDir
	m.checkermanager[req.Name].ActionReqResp.Periodic = req.Periodic
	// m.checkermanager[req.Name].ActionReqResp.DestinationPath = req.DestinationPath

	resp = m.checkermanager[req.Name].ActionReqResp
	m.checkermanager[req.Name].UpdateTicker()
	return resp, nil
}

func (m *myServer) ReapCheck(ctx context.Context, req *pb.CheckActionReqResp) (*pb.CheckActionReqResp, error) {
	println("go to reap check")
	resp := new(pb.CheckActionReqResp)
	if req == nil || req.Name == "" {
		return resp, fmt.Errorf("Request required")
	}
	v, ok := m.checkermanager[req.Name]
	if !ok {
		return resp, fmt.Errorf("CM of %s does not exists", req.Name)
	}
	resp = v.ActionReqResp
	return resp, nil
}

func (m *myServer) DeleteCheck(ctx context.Context, req *pb.CheckActionReqResp) (*pb.CheckActionReqResp, error) {
	println("go to delete check")
	resp := new(pb.CheckActionReqResp)
	if req == nil || req.Name == "" {
		return resp, fmt.Errorf("Request required")
	}
	v, ok := m.checkermanager[req.Name]
	if !ok {
		return resp, fmt.Errorf("CM of %s does not exists", req.Name)
	}
	v.DestroyTicker()
	resp = v.ActionReqResp
	delete(m.checkermanager, req.Name)
	return resp, nil
}

func (m *myServer) writeRedis(content *pb.CheckActionReqResp) error {
	glog.Infoln("write cm into cache")
	k := content.Name
	v, err := json.Marshal(content)
	if err != nil {
		glog.Infoln("Could not marshal into JSON,", err.Error())
		return fmt.Errorf("Failed to mashal into JSON: %s", err.Error())
	}

	c, err := redis.DialURL(fmt.Sprintf("redis://%s", m.redisAddresses[0]), redis.DialDatabase(m.redisDB))
	if err != nil {
		// handle connection error
		glog.Infoln("Could not open cache service,", err.Error())
		return fmt.Errorf("Failed to connect Redis: %s", err.Error())
	}
	defer c.Close()

	cmk := m.cmCache + "." + k
	//	if err := c.Send("SET", cmk, string(v)); err != nil {
	//		glog.Infof("Send(%v, %v) returned error %v", cmk, string(v), err)
	//		return fmt.Errorf("Failed to send: %s", err.Error())
	//	}
	//	c.Flush()

	// if reply, err := c.Do("HMSET", cmk, "name", content.Name, "conf", content.Conf, "periodic", content.Periodic, "state_message", content.StateMessage, "timestamp", content.Timestamp); err != nil {
	if reply, err := c.Do("SET", cmk, string(v)); err != nil {
		glog.Infof("Failed to set cm %s: %s", cmk, err.Error())
		return fmt.Errorf("Failed to set cm %s: %s", cmk, err.Error())
	} else {
		glog.Infof("Set CM %s: %v", cmk, reply)
	}
	return nil
}

func (m *myServer) writeEtcdV3(content *pb.CheckActionReqResp) error {
	k := content.Name
	v, err := content.Marshal()
	if err != nil {
		glog.Infoln("Could not serialize data,", err.Error())
		return fmt.Errorf("Failed to serialize data: %s", err.Error())
	}
	//	if content.StateMessage != "" {
	//		k = strings.Join([]string{content.DestinationPath, contet.Timestamp}, "/")
	//		v = []byte(content.StateMessage)
	//	}
	timeout := time.Second * 10

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(m.etcdAddresses, ","),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		glog.Infoln("Could not open persistent service,", err.Error())
		return fmt.Errorf("Failed to connect EtcdV3: %s", err.Error())
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	resp, err := cli.Put(ctx, k, string(v))
	cancel()
	if err != nil {
		// handle error!
		switch err {
		case context.Canceled:
			glog.Infof("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			glog.Infof("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			glog.Infof("client-side error: %v", err)
		default:
			glog.Infof("bad cluster endpoints, which are not etcd servers: %v", err)
		}
		return fmt.Errorf("Failed to persistent data into EtcdV3: %s", err.Error())
	}
	// use the response
	glog.Infof("%q", resp)
	return nil
}

func (m *myServer) publishRedis(content *pb.CheckActionReqResp) error {
	glog.Infoln("publish check...")
	k := content.Name
	v, err := json.Marshal(content.DestConfigurations)
	if err != nil {
		glog.Infoln("Could not marshal into JSON,", err.Error())
		return fmt.Errorf("Failed to mashal into JSON: %s", err.Error())
	}

	c, err := redis.DialURL(fmt.Sprintf("redis://%s", m.redisAddresses[0]), redis.DialDatabase(m.redisDB))
	if err != nil {
		// handle connection error
		glog.Infoln("Could not open subscription service,", err.Error())
		return fmt.Errorf("Failed to connect Redis: %s", err.Error())
	}
	defer c.Close()

	subj := m.pubSubject + "." + k
	if reply, err := c.Do("PUBLISH", subj, string(v)); err != nil {
		glog.Infof("Failed to publish subject %s: %s", subj, err.Error())
		return fmt.Errorf("Failed to publish subject %s: %s", subj, err.Error())
	} else {
		glog.Infof("Published subject %s: %v", subj, reply)
	}
	return nil
}

func (m *myServer) publishGnatsd(content *pb.CheckActionReqResp) error {
	k := content.Name
	v, err := json.Marshal(content)
	if err != nil {
		glog.Infoln("Could not marshal into JSON,", err.Error())
		return fmt.Errorf("Failed to mashal into JSON: %s", err.Error())
	}
	//	if content.StateMessage != "" {
	//		k = strings.Join([]string{content.DestinationPath, contet.Timestamp}, "/")
	//		v = []byte(content.StateMessage)
	//	}

	var urls = &m.gnatsdAddresses

	nc, err := nats.Connect(*urls)
	if err != nil {
		glog.Infof("Can't connect: %v\n", err)
		return fmt.Errorf("Could not connect gnatsd: %v", err)
	}

	subj, msg := k, v

	nc.Publish(subj, msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		glog.Infof("Failed to mashal into Gnatsd: %v", err)
	}
	glog.Infof("Published [%s] : '%s'\n", subj, msg)
	return nil
}
*/
