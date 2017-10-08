package cmd

import (
	"flag"
	// "fmt"
	// "log"
	// "os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	// "github.com/uber/jaeger/cmd/agent/app"
	// "github.com/uber/jaeger/pkg/config"
	// "github.com/uber/jaeger/pkg/metrics"

	// "go.uber.org/zap"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-docker/tracing/pkg/agent/app"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/config"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/metrics"
	"github.com/tangfeixiong/go-to-docker/tracing/pkg/server"
)

func createJaegerAgentCommand() *cobra.Command {
	var meterdriver, collectorrpc string
	var loglevel string

	// logger, _ := zap.NewProduction()
	v := viper.New()
	command := &cobra.Command{
		Use:   "jaeger-agent",
		Short: "The agent is collecting tracing data with streaming and a gRPC-Gateway for interactive stuff",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", loglevel)
			flag.Parse()

			//			builder := &app.Builder{}
			//			builder.InitFromViper(v)
			//			runtime.GOMAXPROCS(runtime.NumCPU())

			//			// TODO illustrate discovery service wiring
			//			// TODO illustrate additional reporter

			//			agent, err := builder.CreateAgent(logger)
			//			if err != nil {
			//				return errors.Wrap(err, "Unable to initialize Jaeger Agent")
			//			}

			//			logger.Info("Starting agent")
			//			if err := agent.Run(); err != nil {
			//				return errors.Wrap(err, "Failed to run the agent")
			//			}
			//			select {}

			server.RunTracingAgent(v, collectorrpc, meterdriver)
		},
	}

	config.AddFlags(
		v,
		command,
		app.AddFlags,
		metrics.AddFlags,
	)

	command.Flags().StringVar(&collectorrpc, "collector", "", "for collector rpc address, e.g. grpc=localhost:12305")
	command.Flags().StringVar(&meterdriver, "meter", "", "for meter driver, e.g. cadvisor=http://localhost:8080;http://...,...")
	command.Flags().StringVar(&loglevel, "loglevel", "2", "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return command
}
