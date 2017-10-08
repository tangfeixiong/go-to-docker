package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	goruntime "runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	// "github.com/uber/jaeger/cmd/agent/app"

	"go.uber.org/zap"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tangfeixiong/go-to-docker/tracing/pb"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/agent/app"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/dispatcher"
)

type myAgent struct {
	myServer
	*app.Builder
	agentmanager *dispatcher.AgentManager
}

func RunTracingAgent(v *viper.Viper, collectorrpc, meterdriver string) error {
	logger, _ := zap.NewProduction()

	builder := &app.Builder{}
	builder.InitFromViper(v)
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO illustrate discovery service wiring
	// TODO illustrate additional reporter

	agent, err := builder.CreateAgent(logger)
	if err != nil {
		return errors.Wrap(err, "Unable to initialize Jaeger Agent")
	}

	logger.Info("Starting agent")
	if err := agent.Run(); err != nil {
		return errors.Wrap(err, "Failed to run the agent")
	}

	go func() {
		g := gRPCAgent(collectorrpc, meterdriver)
		g.Builder = builder
		g.run()
	}()

	select {}
}

func gRPCAgent(collectorrpc, meterdriver string) *myAgent {
	m := new(myAgent)
	m.grpcHost = ":12365"
	m.httpHost = ":12366"
	m.agentmanager = new(dispatcher.AgentManager)
	// m.exportermanager.Dispatchers = make(map[string]exporter.MeterDispatcher)
	// m.exportermanager.MeteringNameURLs = make(map[string][]string)
	// m.exportermanager.MetricsCollectorRPC = "localhost:12305"
	// m.exportermanager.MeteringNameURLs["docker"] = []string{"unix:///var/run/docker.sock"}

	if v, ok := os.LookupEnv("TRACINGAGENT_GRPC_PORT"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			m.grpcHost = v
		} else {
			m.grpcHost = ":" + v
		}
	}
	if v, ok := os.LookupEnv("TRACINGAGENT_HTTP_PORT"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			m.httpHost = v
		} else {
			m.httpHost = ":" + v
		}
	}

	if collectorrpc != "" {
		if err := os.Setenv("TRACINGCOLLECTOR_RPC_DRIVER", collectorrpc); err != nil {
			fmt.Println("Environment not set, error:", err)
		}
	}
	if v, ok := os.LookupEnv("TRACINGCOLLECTOR_RPC_DRIVER"); ok && 0 != len(v) {
		if strings.HasPrefix(v, "grpc=") {
			m.agentmanager.CollectorDriver = v[5:]
		} else {
			fmt.Println("RPC driver not support,", v)
		}
	}

	return m
}

func (m *myAgent) run() {
	wg := sync.WaitGroup{}
	ch := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.startGRPC(ch)
	}()

	select {
	case <-ch:
		return
	case <-time.After(time.Millisecond * 500):
		fmt.Println("So gRPC is running")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.startGateway()
	}()

	m.dispatchersignal = make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		/*
		   default to read from 'docker stats'?
		*/
		m.agentmanager.Dispatch(m.dispatchersignal)
	}()

	/*
	   https://github.com/kubernetes/kubernetes/blob/release-1.1/build/pause/pause.go
	*/
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Block until a signal is received.
		<-c

		// to stop stuff...
		m.dispatchersignal <- false
		goruntime.Goexit()
	}()

	wg.Wait()
}

func (m *myAgent) startGRPC(ch chan<- bool) {
	host := m.grpcHost

	s := grpc.NewServer()
	pb.RegisterTracingAgentServiceServer(s, m)

	l, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("Failed to listen net, error:", err.Error())
		// panic(err)
		ch <- false
	}

	fmt.Println("Start gRPC on host", l.Addr())
	if err := s.Serve(l); nil != err {
		fmt.Println("Failed to start gRPC, error:", err.Error())
		// panic(err)
		ch <- false
	}
}

func (m *myAgent) startGateway() {
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
	if err := pb.RegisterTracingAgentServiceHandlerFromEndpoint(ctx, gwmux, gRPCHost, dopts); err != nil {
		fmt.Println("Failed to run gRPC-Gateway server. ", err)
	}

	mux.Handle("/", gwmux)
	// serveSwagger(mux)
	//	fmt.Printf("Start HTTP")
	//	if err := http.ListenAndServe(host, allowCORS(mux)); nil != err {
	//		fmt.Fprintf(os.Stderr, "Server died: %s\n", err)
	//	}

	lstn, err := net.Listen("tcp", host)
	if nil != err {
		fmt.Println("Failed to listen to HTTP server. ", err)
		// panic(err)
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
							glog.Infof("preflight request for %s", r.URL.Path)
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
		fmt.Fprintln(os.Stderr, "HTTP server died.", err.Error())
	}
}
