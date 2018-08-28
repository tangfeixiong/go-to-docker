/*
  Inspired by
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/credentialprovider/secrets/secrets.go
*/

package secrets

import (
	"encoding/json"

	"k8s.io/api/core/v1"
	// "k8s.io/kubernetes/pkg/credentialprovider"

	"github.com/tangfeixiong/go-to-docker/pk/credentialprovider"
)

// MakeDockerKeyring inspects the passedSecrets to see if they contain any DockerConfig secrets.  If they do,
// then a DockerKeyring is built based on every hit and unioned with the defaultKeyring.
// If they do not, then the default keyring is returned
func MakeDockerKeyring(passedSecrets []v1.Secret, defaultKeyring credentialprovider.DockerKeyring) (credentialprovider.DockerKeyring, error) {
	passedCredentials := []credentialprovider.DockerConfig{}
	for _, passedSecret := range passedSecrets {
		if dockerConfigJSONBytes, dockerConfigJSONExists := passedSecret.Data[v1.DockerConfigJsonKey]; (passedSecret.Type == v1.SecretTypeDockerConfigJson) && dockerConfigJSONExists && (len(dockerConfigJSONBytes) > 0) {
			dockerConfigJSON := credentialprovider.DockerConfigJson{}
			if err := json.Unmarshal(dockerConfigJSONBytes, &dockerConfigJSON); err != nil {
				return nil, err
			}

			passedCredentials = append(passedCredentials, dockerConfigJSON.Auths)
		} else if dockercfgBytes, dockercfgExists := passedSecret.Data[v1.DockerConfigKey]; (passedSecret.Type == v1.SecretTypeDockercfg) && dockercfgExists && (len(dockercfgBytes) > 0) {
			dockercfg := credentialprovider.DockerConfig{}
			if err := json.Unmarshal(dockercfgBytes, &dockercfg); err != nil {
				return nil, err
			}

			passedCredentials = append(passedCredentials, dockercfg)
		}
	}

	if len(passedCredentials) > 0 {
		basicKeyring := &credentialprovider.BasicDockerKeyring{}
		for _, currCredentials := range passedCredentials {
			basicKeyring.Add(currCredentials)
		}
		return credentialprovider.UnionDockerKeyring{basicKeyring, defaultKeyring}, nil
	}

	return defaultKeyring, nil
}
