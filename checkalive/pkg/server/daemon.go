package server

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nats-io/go-nats"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	healthcheckerpb "github.com/tangfeixiong/go-to-docker/checkalive/pb"
	"github.com/tangfeixiong/go-to-docker/checkalive/pkg/counselor"
)

type myCounselor struct {
	grpcHost               string
	httpHost               string
	packgesHome            string
	checkermanager         map[string]*counselor.CheckerController
	redisSentinelAddresses string
	redisAddresses         []string
	redisDB                int
	etcdAddresses          string
	mysqlAddress           string
	gnatsdAddresses        string
	kafkaAddresses         string
	zookeeperAddresses     string
	rabbitAddress          string
	priorityCMDB           []string
	priorityMQ             []string
	subSubject             string
	subQueue               string
	pubSubject             string
	cmCache                string
	unsubCh                chan string
	pubCh                  chan []byte
}

func Run() {
	mc := new(myCounselor)
	mc.grpcHost = ":10061"
	mc.httpHost = ":10062"
	mc.checkermanager = make(map[string]*counselor.CheckerController)
	mc.redisAddresses = make([]string, 0)
	mc.pubSubject = "checkalive"
	mc.cmCache = "checkalive"
	mc.priorityCMDB = []string{"sentinel", "redis", "etcd", "mysql"}
	mc.priorityMQ = []string{"sentinel", "redis", "gnatsd", "kafka", "rabbit"}

	if v, ok := os.LookupEnv("CHECKALIVE_GRPC_PORT"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			mc.grpcHost = v
		} else {
			mc.grpcHost = "localhost:" + v
		}
	}

	if v, ok := os.LookupEnv("CHECKALIVE_HTTP_PORT"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			mc.httpHost = v
		} else {
			mc.httpHost = ":" + v
		}
	}

	mc.packgesHome = "examples/python/checkalive"
	if v, ok := os.LookupEnv("CHECKALIVE_PACKAGE_HOME"); ok && 0 != len(v) {
		mc.packgesHome = v
	}

	if v, ok := os.LookupEnv("DATABUS_REDIS_SENTINEL_HOSTS"); ok {
		mc.redisSentinelAddresses = v
	}
	if v, ok := os.LookupEnv("DATABUS_REDIS_HOST"); ok {
		mc.redisAddresses = append(mc.redisAddresses, v)
	}
	if v, ok := os.LookupEnv("DATABUS_REDIS_DB"); ok {
		if n, err := strconv.Atoi(v); err != nil {
			glog.Infoln("Failed to parse Redis Db number, using default: ", err)
		} else {
			mc.redisDB = n
		}
	}

	if v, ok := os.LookupEnv("DATABUS_ETCD_HOSTS"); ok {
		mc.etcdAddresses = v
	}
	if v, ok := os.LookupEnv("DATABUS_MYSQL_HOST"); ok {
		mc.mysqlAddress = v
	}

	if v, ok := os.LookupEnv("DATABUS_GNATSD_HOSTS"); ok {
		mc.gnatsdAddresses = v
	}
	if v, ok := os.LookupEnv("DATABUS_KAFKA_HOSTS"); ok {
		mc.kafkaAddresses = v
		if v, ok := os.LookupEnv("DATABUS_ZOOKEEPER_HOSTS"); ok {
			mc.zookeeperAddresses = v
		}
	}
	if v, ok := os.LookupEnv("DATABUS_RABBIT_HOST"); ok {
		mc.rabbitAddress = v
	}

	if v, ok := os.LookupEnv("CHECK_MESSAGE_PUBSUB"); ok && v != "" {
		mc.pubSubject = v
	}

	if v, ok := os.LookupEnv("CHECK_CM_CACHE"); ok && v != "" {
		mc.cmCache = v
	}

	if v, ok := os.LookupEnv("DATABUS_PRIORITY_CMDB"); ok && v != "" {
		mc.priorityCMDB = strings.Split(v, " ")
	}
	if v, ok := os.LookupEnv("DATABUS_PRIORITY_MQ"); ok && v != "" {
		mc.priorityMQ = strings.Split(v, " ")
	}

	// mc.unsubMutex = new(sync.Mutex)
	mc.unsubCh = make(chan string)
	mc.pubCh = make(chan []byte)
	mc.run()
}

func (m *myCounselor) run() {
	wg := sync.WaitGroup{}
	ch := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.startGRPC(ch)
	}()
	time.Sleep(time.Millisecond * 500)

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.startGateway(ch)
	}()

	// wg.Add(1)
	go func() {
		// defer wg.Done()
		/*
		   alternate to read from cmdb: redis/etcd/mysql
		*/
		wg.Wait() //placeholder

	}()

	// wg.Wait()

	/*
	   https://github.com/kubernetes/kubernetes/blob/release-1.1/build/pause/pause.go
	*/
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	<-c
}

func (m *myCounselor) startGRPC(ch chan<- bool) {
	ch <- true
	host := m.grpcHost

	s := grpc.NewServer()
	healthcheckerpb.RegisterCounselorServiceServer(s, m)

	l, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start gRPC on host", l.Addr())
	if err := s.Serve(l); nil != err {
		panic(err)
	}
}

func (m *myCounselor) startGateway(ch <-chan bool) {
	<-ch
	gRPCHost := m.grpcHost
	host := m.httpHost

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	// mux.HandleFunc("/swagger/", serveSwagger2)
	//	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
	//		io.Copy(w, strings.NewReader(healthcheckerpb.Swagger))
	//	})

	dopts := []grpc.DialOption{grpc.WithInsecure()}

	fmt.Println("Start gRPC Gateway into host", gRPCHost)
	gwmux := runtime.NewServeMux()
	if err := healthcheckerpb.RegisterCounselorServiceHandlerFromEndpoint(ctx, gwmux, gRPCHost, dopts); err != nil {
		fmt.Println("Failed to run HTTP server. ", err)
		return
	}

	mux.Handle("/", gwmux)
	// serveSwagger(mux)
	//	fmt.Printf("Start HTTP")
	//	if err := http.ListenAndServe(host, allowCORS(mux)); nil != err {
	//		fmt.Fprintf(os.Stderr, "Server died: %s\n", err)
	//	}

	lstn, err := net.Listen("tcp", host)
	if nil != err {
		panic(err)
	}

	fmt.Printf("http on host: %s\n", lstn.Addr())
	srv := &http.Server{
		Handler: func /*allowCORS*/ (h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if origin := r.Header.Get("Origin"); origin != "" {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
						func /*preflightHandler*/ (w http.ResponseWriter, r *http.Request) {
							headers := []string{"Content-Type", "Accept"}
							w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
							methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
							w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
							// glog.Infof("preflight request for %s", r.URL.Path)
							return
						}(w, r)
						return
					}
				}
				h.ServeHTTP(w, r)
			})
		}(mux),
	}

	if err := srv.Serve(lstn); nil != err {
		fmt.Fprintln(os.Stderr, "Server died.", err.Error())
	}
}

func (m *myCounselor) CreateCheck(ctx context.Context, req *healthcheckerpb.CheckActionReqResp) (*healthcheckerpb.CheckActionReqResp, error) {
	glog.Infof("go to create check: %q", req)
	if req == nil || req.Name == "" {
		return req, fmt.Errorf("Request name is required")
	}
	//	if len(req.DestConfigurations) == 0 {
	//		return m.CreateLegacy(ctx, req)
	//	}
	return m.createCheck(ctx, req)
}

func (m *myCounselor) createCheck(ctx context.Context, req *healthcheckerpb.CheckActionReqResp) (*healthcheckerpb.CheckActionReqResp, error) {
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

func (m *myCounselor) CreateLegacy(ctx context.Context, req *healthcheckerpb.CheckActionReqResp) (*healthcheckerpb.CheckActionReqResp, error) {
	resp := new(healthcheckerpb.CheckActionReqResp)
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

func (m *myCounselor) UpdateCheck(ctx context.Context, req *healthcheckerpb.CheckActionReqResp) (*healthcheckerpb.CheckActionReqResp, error) {
	return m.UpdateWebXCheck(ctx, req)
}

func (m *myCounselor) UpdateWebXCheck(ctx context.Context, req *healthcheckerpb.CheckActionReqResp) (*healthcheckerpb.CheckActionReqResp, error) {
	println("go to update check")
	resp := new(healthcheckerpb.CheckActionReqResp)
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

func (m *myCounselor) ReapCheck(ctx context.Context, req *healthcheckerpb.CheckActionReqResp) (*healthcheckerpb.CheckActionReqResp, error) {
	println("go to reap check")
	resp := new(healthcheckerpb.CheckActionReqResp)
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

func (m *myCounselor) DeleteCheck(ctx context.Context, req *healthcheckerpb.CheckActionReqResp) (*healthcheckerpb.CheckActionReqResp, error) {
	println("go to delete check")
	resp := new(healthcheckerpb.CheckActionReqResp)
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

func (m *myCounselor) writeRedis(content *healthcheckerpb.CheckActionReqResp) error {
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

func (m *myCounselor) writeEtcdV3(content *healthcheckerpb.CheckActionReqResp) error {
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

func (m *myCounselor) publishRedis(content *healthcheckerpb.CheckActionReqResp) error {
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

func (m *myCounselor) publishGnatsd(content *healthcheckerpb.CheckActionReqResp) error {
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
