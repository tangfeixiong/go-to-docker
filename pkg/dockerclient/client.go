package dockerclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"golang.org/x/net/context"
)

const (
	CLIENT_GENERATE_FAILURE  = 100
	REGISTRY_AUTH_UNEXPECTED = 101
	IMAGE_PULL_FAILURE       = 102
	ILLEGAL_IMAGE_REF        = 110
)

type DockerClient struct {
	// ctx              context.Context
	dockerClient     *client.Client
	registryAuthB64s map[string]string
	dockerconfigjson DockerConfigJSON
	pullQueue        map[string]string
}

func NewOrDie() *DockerClient {
	// ctx := context.Background()
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	cli := &DockerClient{
		dockerClient:     dockerClient,
		registryAuthB64s: make(map[string]string),
		pullingQueue:     make(map[string]string),
	}

	if v, ok := os.LookupEnv("DOCKER_CONFIG_JSON"); ok {
		if err := json.Unmarshal([]byte(v), &cli.dockerconfigjson); nil != err {
			// fmt.Println("Illegal DOCKER_CONFIG_JSON environment.", err.Error())
			// return nil, fmt.Errorf("Could not unmarshall DOCKER_CONFIG_JSON env value: %s", error.Error())
			panic(err)
		}
		if nil == cli.dockerconfigjson.Auths {
			cli.dockerconfigjson.Auths = make(map[string]types.AuthConfig)
		} else {
			for k, v := range cli.dockerconfigjson.Auths {
				cli.registryAuthB64s[k] = v
				// sDec, err := base64.StdEncoding.DecodeString(v.Auth)
				sDec, err := base64.URLEncoding.DecodeString(v.Auth)
				if err != nil {
					// fmt.Println("Invalid credential.", err.Error())
					// return nil, fmt.Errorf("Could not get credential: %s", error.Error())
					panic(err)
				} else {
					i := strings.Index(string(sDec), ":")
					if -1 == i {
						// fmt.Println("Invalid basicauth.")
						// return nil, fmt.Errorf("Illegal format of basicauth credential")
						panic(err)
					} else {
						v.Username = string(sDec[:i])
						v.Password = string(sDec[i+1:])
						v.Auth = ""
						cli.dockerconfigjson.Auths[k] = v
					}
				}
			}
		}
	}

	return cli
}

type DockerConfigJSON struct {
	Auths map[string]types.AuthConfig `json:"auths,omitempty"`
}

func (dc *DockerClient) RegistryAuthConfig(image string) (types.AuthConfig, bool) {
	s := strings.Split(image, "/")
	if len(s) == 1 && s[0] == "docker.io" {
		if v, ok := dc.dockerconfigjson.Auths["docker.io"]; ok {
			return v, true
		} else {
			for k, v := range dc.dockerconfigjson.Auths {
				if strings.Contains(k, "docker.io") /* "https://index.docker.io/v1/" */ {
					return v, true
				}
			}
			return types.AuthConfig{}, false
		}
	}
	if v, ok := dc.dockerconfigjson.Auths[s[0]]; ok {
		return v, true
	} else {
		for k, v := range dc.dockerconfigjson.Auths {
			if strings.Contains(k, "docker.io") /* "https://index.docker.io/v1/" */ {
				return v, true
			}
		}
	}
	return types.AuthConfig{}, false
}

func (dc *DockerClient) RegistryAuthB64Encoded(image string) (string, bool) {
	s := strings.Split(image, "/")
	if len(S) == 1 && s[0] == "docker.io" {
		if v, ok := dc.registryAuthB64s["docker.io"]; ok {
			return v, true
		}
		if v, ok := dc.registryAuthB64s["https://index.docker.io/v1/"]; ok {
			return v, true
		}
	}
	if v, ok := dc.registryAuthB64s[s[0]]; ok {
		return v, true
	}
	return "", false
}
