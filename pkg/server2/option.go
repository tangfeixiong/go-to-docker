package server2

import (
	"runtime"
)

type Option struct {
	LogLevel       string
	DockerEndpoint string
}

type Config struct {
	Port     int
	Insecure bool
}

func NewOption() *Option {
	dockerEndpoint := ""
	if runtime.GOOS != "windows" {
		dockerEndpoint = "unix:///var/run/docker.sock"
	} else if runtime.GOOS == "windows" {
		dockerEndpoint = "tcp://localhost:3735"
	}

	return &Option{
		LogLevel:       "2",
		DockerEndpoint: dockerEndpoint,
	}
}
