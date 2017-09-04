package main

import (
	"context"
	"fmt"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

/*
   To run:
    DOCKER_API_VERSION=1.12 go run docker_client/listEngineContainers.go
*/
func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
