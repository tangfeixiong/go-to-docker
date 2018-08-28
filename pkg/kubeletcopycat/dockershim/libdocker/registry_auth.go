/*
  Refer to
  - https://github.com/docker/distribution/blob/master/reference/normalize.go
*/

package libdocker

import (
	dockertypes "github.com/docker/docker/api/types"
)

var (
	legacyDefaultDomain = "index.docker.io"
	defaultDomain       = "docker.io"
	officialRepoName    = "library"
	defaultTag          = "latest"
)

type DockerConfigJSON struct {
	Auths map[string]dockertypes.AuthConfig `json:"auths,omitempty"`
}

func (dc *DockerClient) RegistryAuthConfig(image string) (dockertypes.AuthConfig, bool) {
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
			return dockertypes.AuthConfig{}, false
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
	return dockertypes.AuthConfig{}, false
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
