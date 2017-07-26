package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"google.golang.org/grpc"

	"k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pb/moby"
	"github.com/tangfeixiong/go-to-docker/pkg/server"
	"github.com/tangfeixiong/go-to-docker/pkg/util"
)

var (
	/*tmplLong = templates.LongDesc*/ tmplLong string = (`
		%[2]s

		The %[3]s helps you build, deploy, and manage your applications on top of
		Docker containers. To start an all-in-one server with the default configuration, run:

		    $ %[1]s serve &`)
)

type Option struct {
	endpoint    string
	loglevel    string
	runTestType string
	image       string
	name        string
}

func NewRootCommand(name string) *cobra.Command {
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "Deploy dockerized applications",
		Long:  fmt.Sprintf(tmplLong, name, name, name),
	}
	root.AddCommand(newServeCommand())
	root.AddCommand(newRunTestCommand())

	return root
}

func newServeCommand() *cobra.Command {
	var opt Option

	command := &cobra.Command{
		Use:   "serve",
		Short: "Start gRPC-Gateway service",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", opt.loglevel)
			flag.Parse()

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

	command.Flags().StringVar(&opt.loglevel, "loglevel", "2", "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return command
}

func newRunTestCommand() *cobra.Command {
	var opt Option

	command := &cobra.Command{
		Use:   "runTest",
		Short: "Start runtime test",
	}
	command.AddCommand(newRunContainerCommand(&opt))

	command.PersistentFlags().StringVar(&opt.runTestType, "type", "container", "for image or container")

	return command
}

func newRunContainerCommand(opt *Option) *cobra.Command {
	command := &cobra.Command{
		Use:   "run",
		Short: "Run a new container",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial("localhost:10051", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := pb.NewEchoServiceClient(conn)

			// Contact the server and print out its response.
			if len(os.Args) > 1 {
				log.Printf("args: %+v", os.Args[1:])
			}
			req := &pb.DockerRunData{
				Config: &moby.Config{
					Image: "springcloud/eureka",
					ExposedPorts: &moby.PortSet{
						Value: map[string]string{
							"8761": "webui",
						},
					},
				},
				HostConfig: &moby.HostConfig{
					PortBindings: &moby.PortMap{
						Value: map[string]*moby.PortBinding{
							"8761": &moby.PortBinding{
								HostIp:   "",
								HostPort: "8761",
							},
						},
					},
				},
				NetworkConfig: &moby.NetworkingConfig{},
				ContainerName: "runtest-" + rand.String(12),
			}
			r, err := c.RunContainer(context.Background(), req)
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			util.Logger.Printf("Greeting: %v", r)
		},
	}

	command.Flags().StringVar(&opt.image, "image", "nginx", "image for create container")
	command.Flags().StringVar(&opt.name, "name", "nginx", "container name")

	return command
}
