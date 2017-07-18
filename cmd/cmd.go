package cmd

import (
	"flag"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/tangfeixiong/go-to-docker/pkg/server"
)

var (
	/*tmplLong = templates.LongDesc*/ tmplLong string = (`
		%[2]s

		The %[3]s helps you build, deploy, and manage your applications on top of
		Docker containers. To start an all-in-one server with the default configuration, run:

		    $ %[1]s serve &`)
)

type Option struct {
	endpoint string
	example  string
}

func NewRootCommand(name string) *cobra.Command {
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "Deploy dockerized applications",
		Long:  fmt.Sprintf(tmplLong, name, name, name),
	}
	root.AddCommand(newServeCommand())

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()

	return root
}

func newServeCommand() *cobra.Command {
	var opt Option

	command := &cobra.Command{
		Use:   "serve",
		Short: "Start gRPC-Gateway service",
		Run: func(cmd *cobra.Command, args []string) {
			wg := sync.WaitGroup{}

			ch := make(chan string)
			wg.Add(1)
			go func() {
				defer wg.Done()
				server.StartGRPC(ch)
			}()
			opt.endpoint = <-ch

			wg.Add(1)
			go func() {
				defer wg.Done()
				server.StartGateway(opt.endpoint)
			}()

			wg.Wait()
		},
	}

	command.Flags().StringVar(&opt.example, "example", "", "for example")

	return command
}
