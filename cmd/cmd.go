package cmd

import (
	"flag"
	"fmt"

	//"os"
	//"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/tangfeixiong/go-to-docker/pkg/server2"
)

var (
	/*tmplLong = templates.LongDesc*/ tmplLong string = (`
		%[2]s

		The %[3]s helps you build, deploy, and manage your applications on top of
		Docker containers. To start an all-in-one server with the default configuration, run:

		    $ %[1]s serve &`)
)

func NewRootCommand(name string) *cobra.Command {
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "A Docker micro-service",
		Long:  fmt.Sprintf(tmplLong, name, name, name),
	}
	root.AddCommand(newServe2Command())

	return root
}

func newServe2Command() *cobra.Command {
	opt := server2.NewOption()

	command := &cobra.Command{
		Use:   "serve2",
		Short: "Start gRPC with HTTP gateway service",
		Run: func(cmd *cobra.Command, args []string) {
			flag.Set("v", opt.LogLevel)
			flag.Parse()

			server2.Start(opt)
		},
	}

	command.Flags().StringVar(&opt.LogLevel, "loglevel", "2", "for glog")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}
