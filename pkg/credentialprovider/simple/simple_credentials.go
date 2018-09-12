package simple

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pkg/credentialprovider"
)

type DockerConfigJSON struct {
	Auths map[string]dockertypes.AuthConfig `json:"auths,omitempty"`
}

var CustomDockerConfigJSON *DockerConfigJSON

func decode(jsonstr string) (*DockerConfigJSON, error) {
	var dockerconfigjson DockerConfigJSON
	dockerconfigjson.Auths = make(map[string]dockertypes.AuthConfig)

	if err := json.Unmarshal([]byte(jsonstr), &dockerconfigjson); nil != err {
		glog.Errorf("failed to decode JSON str: %v", err)
		return &dockerconfigjson, fmt.Errorf("failed to decode JSON str: %s", err.Error())
	} else {
		for k, v := range dockerconfigjson.Auths {
			sDec, err := base64.StdEncoding.DecodeString(v.Auth)
			// sDec, err := base64.URLEncoding.DecodeString(v.Auth)
			if err != nil {
				glog.Errorf("failed to parse credential from DockerConfigJSON: %v", err)
				return &dockerconfigjson, fmt.Errorf("failed to parse credential from DockerConfigJSON: %s", err.Error())
			} else {
				i := strings.Index(string(sDec), ":")
				if -1 == i {
					glog.Warningln("Invalid basicauth from a docker AuthConfig.")
				} else {
					v.Username = string(sDec[:i])
					v.Password = string(sDec[i+1:])
					// v.Auth = ""
					dockerconfigjson.Auths[k] = v
				}
			}
		}
	}
	return &dockerconfigjson, nil
}

func init() {
	credentialprovider.RegisterCredentialProvider("stackdocker-configjson",
		&credentialprovider.CachingDockerConfigProvider{
			Provider: NewSimpleProvider(),
			Lifetime: 30 * time.Second,
		})
}

type simpleProvider struct {
}

func NewSimpleProvider() credentialprovider.DockerConfigProvider {
	CustomDockerConfigJSON = &DockerConfigJSON{
		Auths: make(map[string]dockertypes.AuthConfig),
	}
	if v, ok := os.LookupEnv("DOCKER_CONFIG_JSON"); ok {
		dockerconfigjson, err := decode(v)
		if err == nil {
			for k, v := range dockerconfigjson.Auths {
				CustomDockerConfigJSON.Auths[k] = v
			}
		}
	}
	if v, ok := os.LookupEnv("DOCKER_CONFIG_JSON_BASE64"); ok {
		if sDec, err := base64.URLEncoding.DecodeString(v); err != nil {
			glog.Warningf("failed to decode DockerConfigJSON with Base64 from ENV: %v", err.Error())
		} else {
			dockerconfigjson, err := decode(string(sDec))
			if err == nil {
				for k, v := range dockerconfigjson.Auths {
					CustomDockerConfigJSON.Auths[k] = v
				}
			}
		}
	}
	return &simpleProvider{}
}

func (p *simpleProvider) Enabled() bool {
	return true
}

// Provide implements dockerConfigProvider
func (p *simpleProvider) Provide() credentialprovider.DockerConfig {
	cfg := credentialprovider.DockerConfig{}
	for k, v := range CustomDockerConfigJSON.Auths {
		cfg[k] = credentialprovider.DockerConfigEntry{
			Username: v.Username,
			Password: v.Password,
			Email:    v.Email,
		}
	}
	return cfg
}

// LazyProvide implements dockerConfigProvider. Should never be called.
func (p *simpleProvider) LazyProvide() *credentialprovider.DockerConfigEntry {
	return nil
}
