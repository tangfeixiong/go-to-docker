package dockerctl

import (
	"context"
	"fmt"
	"os"
	"time"

	// dockerdigest "github.com/docker/distribution/digest"
	// dockerref "github.com/docker/distribution/reference"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/network"
	// "github.com/docker/engine-api/types/registry"
	"github.com/golang/glog"

	// "k8s.io/kubernetes/pkg/util/parsers"
)

type Engine1_12Client struct {
	version      string
	dockerclient *client.Client
}

func NewEngine1_12Client() *Engine1_12Client {
	var ver string = Default_docker_API_ver
	var err error = nil
	if v, ok := os.LookupEnv("DOCKER_API_VERSION"); !ok {
		v := Default_docker_API_ver
		err = os.Setenv("DOCKER_API_VERSION", v)
	} else {
		if Default_docker_API_ver != v {
			v := Default_docker_API_ver
			err = os.Setenv("DOCKER_API_VERSION", v)
		}
	}
	if err != nil {
		fmt.Println("Failed to configure DOCKER_API_VERSION environment.", err.Error())
	}
	return &Engine1_12Client{
		version: ver,
	}
}

func (mc *Engine1_12Client) DockerClient() (*client.Client, error) {
	return client.NewEnvClient()
}

func (mc *Engine1_12Client) CreateContainer(config *container.Config, hostconfig *container.HostConfig, networkconfig *network.NetworkingConfig, containername string) (types.ContainerCreateResponse, error) {
	glog.Infoln("Go to create container:", containername, "DOCKER_API_VERSION=", os.Getenv("DOCKER_API_VERSION"))

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infof("Could not instantiate docker engine api(%s): %s", mc.version, err.Error())
		return types.ContainerCreateResponse{}, fmt.Errorf("Failed to instantiate docker engine api(%s): %v", mc.version, err)
	}

	resp, err := cli.ContainerCreate(context.Background(), config, hostconfig, networkconfig, containername)
	if err != nil {
		glog.V(2).Infoln("Could not create container:", err.Error())
		return types.ContainerCreateResponse{}, fmt.Errorf("Failed to create container. %v", err)
	}
	return resp, nil
}

func (mc *Engine1_12Client) StartContainer(containerid string) error {
	glog.Infoln("Go to start container:", containerid)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infof("Could not instantiate docker engine api(%s): %s", mc.version, err.Error())
		return fmt.Errorf("Failed to instantiate docker engine api(%s): %v", mc.version, err)
	}

	err = cli.ContainerStart(context.Background(), containerid)
	if nil != err {
		glog.V(2).Infoln("Could not start container:", err.Error())
		return fmt.Errorf("Failed to start container. %v", err)
	}
	return nil
}

func (mc *Engine1_12Client) StopContainer(containerid string, timeout time.Duration) error {
	glog.Infoln("Go to stop container:", containerid)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infof("Could not instantiate docker engine api(%s): %s", mc.version, err.Error())
		return fmt.Errorf("Failed to instantiate docker engine api(%s): %v", mc.version, err)
	}

	err = cli.ContainerStop(context.Background(), containerid, int(timeout.Seconds()))
	if nil != err {
		glog.V(2).Infoln("Could not stop container:", err.Error())
		return fmt.Errorf("Failed to stop container. %v", err)
	}
	return nil
}

func (mc *Engine1_12Client) RemoveContainer(containerid string) error {
	glog.Infoln("Go to remove container:", containerid)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infof("Could not instantiate docker engine api(%s): %s", mc.version, err.Error())
		return fmt.Errorf("Failed to instantiate docker engine api(%s): %v", mc.version, err)
	}

	opt := types.ContainerRemoveOptions{}
	err = cli.ContainerRemove(context.Background(), containerid, opt)
	if nil != err {
		glog.V(2).Infoln("Could not remove container:", err.Error())
		return fmt.Errorf("Failed to remove container. %v", err)
	}
	return nil
}

func (mc *Engine1_12Client) ProcessStatusContainers(opt types.ContainerListOptions) ([]types.Container, error) {
	glog.Infoln("Go to list container:", opt)

	cli, err := client.NewEnvClient()
	if err != nil {
		glog.V(2).Infof("Could not instantiate docker engine api(%s): %s", mc.version, err.Error())
		return nil, fmt.Errorf("Failed to instantiate docker engine api(%s): %v", mc.version, err)
	}

	result, err := cli.ContainerList(context.Background(), opt)
	if nil != err {
		glog.V(2).Infoln("Could not remove container:", err.Error())
		return nil, fmt.Errorf("Failed to remove container. %v", err)
	}
	return result, nil
}
