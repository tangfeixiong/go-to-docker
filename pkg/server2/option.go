package server2

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"runtime"
	"strings"

	"github.com/docker/docker/api/types"
)

type Option struct {
	LogLevel       string
	DockerEndpoint string
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
