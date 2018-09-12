package server2

import (
	"github.com/tangfeixiong/go-to-docker/pkg/credentialprovider/simple"
)

type Config struct {
	Port             int
	Insecure         bool
	dockerconfigjson *simple.DockerConfigJSON
}

func NewConfig() *Config {
	config := &Config{
		Port:             10052,
		dockerconfigjson: simple.CustomeDockerConfigJSON,
	}

	return config
}
