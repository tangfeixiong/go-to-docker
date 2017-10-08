package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	goruntime "runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-lib/metrics/go-kit"
	"github.com/uber/jaeger-lib/metrics/go-kit/expvar"
	basicB "github.com/uber/jaeger/cmd/builder"
	"github.com/uber/jaeger/cmd/collector/app"
	// "github.com/uber/jaeger/cmd/collector/app/builder"
	"github.com/uber/jaeger/cmd/collector/app/zipkin"
	// "github.com/uber/jaeger/cmd/flags"
	// casFlags "github.com/uber/jaeger/cmd/flags/cassandra"
	// esFlags "github.com/uber/jaeger/cmd/flags/es"
	"github.com/uber/jaeger/pkg/healthcheck"
	"github.com/uber/jaeger/pkg/recoveryhandler"
	jc "github.com/uber/jaeger/thrift-gen/jaeger"
	zc "github.com/uber/jaeger/thrift-gen/zipkincore"
	"github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/thrift"

	"go.uber.org/zap"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tangfeixiong/go-to-docker/tracing/pb"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/collector/app/builder"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/dispatcher"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/flags"
	casFlags "github.com/tangfeixiong/go-to-docker/tracing/pkg/flags/cassandra"
	esFlags "github.com/tangfeixiong/go-to-docker/tracing/pkg/flags/es"
)

type myCollector struct {
	myServer
	collectormanager *dispatcher.CollectorManager
}

func RunTracingCollector(v *viper.Viper, casOptions *casFlags.Options, esOptions *esFlags.Options, storage string) {
	var signalsChannel = make(chan os.Signal, 0)
	signal.Notify(signalsChannel, os.Interrupt, syscall.SIGTERM)

	logger, _ := zap.NewProduction()
	serviceName := "jaeger-collector"
	casOptions.InitFromViper(v)
	esOptions.InitFromViper(v)

	baseMetrics := xkit.Wrap(serviceName, expvar.NewFactory(10))

	builderOpts := new(builder.CollectorOptions).InitFromViper(v)
	sFlags := new(flags.SharedFlags).InitFromViper(v)

	hc, err := healthcheck.Serve(http.StatusServiceUnavailable, builderOpts.CollectorHealthCheckHTTPPort, logger)
	if err != nil {
		logger.Fatal("Could not start the health check server.", zap.Error(err))
	}

	handlerBuilder, err := builder.NewSpanHandlerBuilder(
		builderOpts,
		sFlags,
		basicB.Options.CassandraSessionOption(casOptions.GetPrimary()),
		basicB.Options.ElasticClientOption(esOptions.GetPrimary()),
		basicB.Options.LoggerOption(logger),
		basicB.Options.MetricsFactoryOption(baseMetrics),
	)
	if err != nil {
		logger.Fatal("Unable to set up builder", zap.Error(err))
	}

	ch, err := tchannel.NewChannel(serviceName, &tchannel.ChannelOptions{})
	if err != nil {
		logger.Fatal("Unable to create new TChannel", zap.Error(err))
	}
	server := thrift.NewServer(ch)
	zipkinSpansHandler, jaegerBatchesHandler := handlerBuilder.BuildHandlers()
	server.Register(jc.NewTChanCollectorServer(jaegerBatchesHandler))
	server.Register(zc.NewTChanZipkinCollectorServer(zipkinSpansHandler))

	portStr := ":" + strconv.Itoa(builderOpts.CollectorPort)
	listener, err := net.Listen("tcp", portStr)
	if err != nil {
		logger.Fatal("Unable to start listening on channel", zap.Error(err))
	}
	ch.Serve(listener)

	r := mux.NewRouter()
	apiHandler := app.NewAPIHandler(jaegerBatchesHandler)
	apiHandler.RegisterRoutes(r)
	httpPortStr := ":" + strconv.Itoa(builderOpts.CollectorHTTPPort)
	recoveryHandler := recoveryhandler.NewRecoveryHandler(logger, true)

	go startZipkinHTTPAPI(logger, builderOpts.CollectorZipkinHTTPPort, zipkinSpansHandler, recoveryHandler)

	logger.Info("Starting Jaeger Collector HTTP server", zap.Int("http-port", builderOpts.CollectorHTTPPort))

	go func() {
		if err := http.ListenAndServe(httpPortStr, recoveryHandler(r)); err != nil {
			logger.Fatal("Could not launch service", zap.Error(err))
		}
		hc.Set(http.StatusInternalServerError)
	}()

	hc.Ready()

	go func() {
		g := gRPCCollector(storage)
		g.run()
	}()

	select {
	case <-signalsChannel:
		logger.Info("Jaeger Collector is finishing")
	}
}

func startZipkinHTTPAPI(
	logger *zap.Logger,
	zipkinPort int,
	zipkinSpansHandler app.ZipkinSpansHandler,
	recoveryHandler func(http.Handler) http.Handler,
) {
	if zipkinPort != 0 {
		r := mux.NewRouter()
		zipkin.NewAPIHandler(zipkinSpansHandler).RegisterRoutes(r)
		httpPortStr := ":" + strconv.Itoa(zipkinPort)
		logger.Info("Listening for Zipkin HTTP traffic", zap.Int("zipkin.http-port", zipkinPort))

		if err := http.ListenAndServe(httpPortStr, recoveryHandler(r)); err != nil {
			logger.Fatal("Could not launch service", zap.Error(err))
		}
	}
}

func gRPCCollector(storage string) *myCollector {
	m := new(myCollector)
	m.grpcHost = ":12355"
	m.httpHost = ":12356"
	m.collectormanager = new(dispatcher.CollectorManager)
	// m.collectormanager.MetricsStorageDuration = time.Second * 10

	if v, ok := os.LookupEnv("TRACINGCOLLECTOR_GRPC_PORT"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			m.grpcHost = v
		} else {
			m.grpcHost = ":" + v
		}
	}
	if v, ok := os.LookupEnv("TRACINGCOLLECTOR_HTTP_PORT"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			m.httpHost = v
		} else {
			m.httpHost = ":" + v
		}
	}

	if storage != "" {
		if err := os.Setenv("TRACES_STORAGE_DRIVER", storage); err != nil {
			fmt.Println("Environment not set, error:", err)
		}
	}
	if v, ok := os.LookupEnv("TRACES_STORAGE_DRIVER"); ok && 0 != len(v) {
		if strings.HasPrefix(v, "elasticsearch=") {
			m.collectormanager.StorageDriver = v
		} else {
			fmt.Println("RPC driver not support,", v)
		}
	}

	return m
}

func (m *myCollector) run() {
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
		   default to write to stdout?
		*/
		m.collectormanager.Dispatch(m.dispatchersignal)
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

func (m *myCollector) startGRPC(ch chan<- bool) {
	host := m.grpcHost

	s := grpc.NewServer()
	pb.RegisterTracingAgentServiceServer(s, m)

	l, err := net.Listen("tcp", host)
	if err != nil {
		// panic(err)
		ch <- false
	}

	fmt.Println("Start gRPC on host", l.Addr())
	if err := s.Serve(l); nil != err {
		// panic(err)
		ch <- false
	}
}

func (m *myCollector) startGateway() {
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
		fmt.Fprintln(os.Stderr, "Server died.", err.Error())
	}
}
