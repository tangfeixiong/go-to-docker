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
		Port:             10053,
		dockerconfigjson: simple.CustomDockerConfigJSON,
	}

	return config
}
