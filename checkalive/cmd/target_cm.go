package cmd

import (
	// "context"
	"flag"
	// "fmt"
	// "log"
	// "os"
	// "sync"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	// "github.com/tangfeixiong/go-to-docker/pb"
	// "github.com/tangfeixiong/go-to-docker/pb/moby"
	"github.com/tangfeixiong/go-to-docker/checkalive/pkg/server"
)

func CommandFor(name string) *cobra.Command {
	return createRootCommand(name)
}

func createRootCommand(name string) *cobra.Command {
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "Configuration Mgmt of dockerized applications",
		Long: `
        target-cm
        
        This is running on Docker Engine,
        and designed to spawn alive checker for some container,
        It is based on wrapping 3-rd party cli.
        `,
	}
	root.AddCommand(createServeCommand())
	// root.AddCommand(createSnoopCommand())
	// root.AddCommand(createSniffCommand())

	return root
}

func createServeCommand() *cobra.Command {
	var loglevel string

	command := &cobra.Command{
		Use:   "serve",
		Short: "Start gRPC-Gateway service",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", loglevel)
			flag.Parse()
			server.Run()
		},
	}

	command.Flags().StringVar(&loglevel, "loglevel", "2", "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return command
}

func createSnoopCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "snoop",
	}
	return command
}

func createSniffCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "sniff",
	}
	return command
}
