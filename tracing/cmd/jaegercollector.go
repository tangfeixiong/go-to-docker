package cmd

import (
	"flag"
	// "fmt"
	// "log"
	// "os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	// "github.com/uber/jaeger/cmd/collector/app/builder"
	// "github.com/uber/jaeger/pkg/config"
	// "github.com/uber/jaeger/cmd/flags"
	// casFlags "github.com/uber/jaeger/cmd/flags/cassandra"
	// esFlags "github.com/uber/jaeger/cmd/flags/es"

	// "go.uber.org/zap"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-docker/tracing/pkg/collector/app/builder"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/config"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/flags"
	casFlags "github.com/tangfeixiong/go-to-docker/tracing/pkg/flags/cassandra"
	esFlags "github.com/tangfeixiong/go-to-docker/tracing/pkg/flags/es"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/server"
)

func createJaegerCollectorCommand() *cobra.Command {
	var storage string
	var loglevel string

	//	var signalsChannel = make(chan os.Signal, 0)
	//	signal.Notify(signalsChannel, os.Interrupt, syscall.SIGTERM)

	// logger, _ := zap.NewProduction()
	// serviceName := "jaeger-collector"
	casOptions := casFlags.NewOptions("cassandra")
	esOptions := esFlags.NewOptions("es")

	v := viper.New()
	command := &cobra.Command{
		Use:   "jaeger-collector",
		Short: "THe collector is receives and processes traces from agent with streaming and a gRPC-Gateway for interactive stuff",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", loglevel)
			flag.Parse()

			//			casOptions.InitFromViper(v)
			//			esOptions.InitFromViper(v)

			//			baseMetrics := xkit.Wrap(serviceName, expvar.NewFactory(10))

			//			builderOpts := new(builder.CollectorOptions).InitFromViper(v)
			//			sFlags := new(flags.SharedFlags).InitFromViper(v)

			//			hc, err := healthcheck.Serve(http.StatusServiceUnavailable, builderOpts.CollectorHealthCheckHTTPPort, logger)
			//			if err != nil {
			//				logger.Fatal("Could not start the health check server.", zap.Error(err))
			//			}

			//			handlerBuilder, err := builder.NewSpanHandlerBuilder(
			//				builderOpts,
			//				sFlags,
			//				basicB.Options.CassandraSessionOption(casOptions.GetPrimary()),
			//				basicB.Options.ElasticClientOption(esOptions.GetPrimary()),
			//				basicB.Options.LoggerOption(logger),
			//				basicB.Options.MetricsFactoryOption(baseMetrics),
			//			)
			//			if err != nil {
			//				logger.Fatal("Unable to set up builder", zap.Error(err))
			//			}

			//			ch, err := tchannel.NewChannel(serviceName, &tchannel.ChannelOptions{})
			//			if err != nil {
			//				logger.Fatal("Unable to create new TChannel", zap.Error(err))
			//			}
			//			server := thrift.NewServer(ch)
			//			zipkinSpansHandler, jaegerBatchesHandler := handlerBuilder.BuildHandlers()
			//			server.Register(jc.NewTChanCollectorServer(jaegerBatchesHandler))
			//			server.Register(zc.NewTChanZipkinCollectorServer(zipkinSpansHandler))

			//			portStr := ":" + strconv.Itoa(builderOpts.CollectorPort)
			//			listener, err := net.Listen("tcp", portStr)
			//			if err != nil {
			//				logger.Fatal("Unable to start listening on channel", zap.Error(err))
			//			}
			//			ch.Serve(listener)

			//			r := mux.NewRouter()
			//			apiHandler := app.NewAPIHandler(jaegerBatchesHandler)
			//			apiHandler.RegisterRoutes(r)
			//			httpPortStr := ":" + strconv.Itoa(builderOpts.CollectorHTTPPort)
			//			recoveryHandler := recoveryhandler.NewRecoveryHandler(logger, true)

			//			go startZipkinHTTPAPI(logger, builderOpts.CollectorZipkinHTTPPort, zipkinSpansHandler, recoveryHandler)

			//			logger.Info("Starting Jaeger Collector HTTP server", zap.Int("http-port", builderOpts.CollectorHTTPPort))

			//			go func() {
			//				if err := http.ListenAndServe(httpPortStr, recoveryHandler(r)); err != nil {
			//					logger.Fatal("Could not launch service", zap.Error(err))
			//				}
			//				hc.Set(http.StatusInternalServerError)
			//			}()

			//			hc.Ready()
			//			select {
			//			case <-signalsChannel:
			//				logger.Info("Jaeger Collector is finishing")
			//			}

			server.RunTracingCollector(v, casOptions, esOptions, storage)
		},
	}

	config.AddFlags(
		v,
		command,
		flags.AddFlags,
		builder.AddFlags,
		casOptions.AddFlags,
		esOptions.AddFlags,
	)

	command.Flags().StringVar(&storage, "storage", "", "for storage address, e.g. elasticsearch=http://localhost:9200")
	command.Flags().StringVar(&loglevel, "loglevel", "2", "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return command
}

//func startZipkinHTTPAPI(
//	logger *zap.Logger,
//	zipkinPort int,
//	zipkinSpansHandler app.ZipkinSpansHandler,
//	recoveryHandler func(http.Handler) http.Handler,
//) {
//	if zipkinPort != 0 {
//		r := mux.NewRouter()
//		zipkin.NewAPIHandler(zipkinSpansHandler).RegisterRoutes(r)
//		httpPortStr := ":" + strconv.Itoa(zipkinPort)
//		logger.Info("Listening for Zipkin HTTP traffic", zap.Int("zipkin.http-port", zipkinPort))

//		if err := http.ListenAndServe(httpPortStr, recoveryHandler(r)); err != nil {
//			logger.Fatal("Could not launch service", zap.Error(err))
//		}
//	}
//}
