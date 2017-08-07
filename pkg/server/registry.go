package server

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/docker/engine-api/types"
	"github.com/heroku/docker-registry-client/registry"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pkg/dockerctl"
)

func (m *myService) reapRegistryForRepositories(req *pb.RegistryRepositoryData) (*pb.RegistryRepositoryData, error) {
	resp := new(pb.RegistryRepositoryData)
	resp.Registries = make([]*pb.RegistryRepositoryData_Registry, 0)
	if nil == req || 0 == len(req.Registries) {
		return resp, fmt.Errorf("Request required")
	}

	// certs, err := gainCerts()
	if nil != m.err {
		resp.StateCode = 100
		resp.StateMessage = m.err.Error()
		return resp, m.err
	}

	ctl := dockerctl.NewEngine1_12Client()

	auth := types.AuthConfig{
		// Username: "",
		// Password: "",
		Auth:          "",
		Email:         "",
		ServerAddress: "127.0.0.1:5000",
		// IdentityToken: "",
		// RegistryToken: "",
	}

	url := "https://registry-1.docker.io/"
	username := "" // anonymous
	password := "" // anonymous
	var transport *http.Transport

	for _, r := range req.Registries {
		auth = ctl.RegistryAuth(r.Name)
		username = auth.Username
		password = auth.Password
		if r.TlsDisabled {
			url = "http://" + r.Name + "/"
		} else {
			url = "https://" + r.Name + "/"
			cert, ok := m.certs[r.Name]
			if ok {
				if len(cert.CertificateAuthority) > 0 {
					caCertPool := x509.NewCertPool()
					caCertPool.AppendCertsFromPEM(cert.CertificateAuthority)
					if len(cert.ClientCertificate) > 0 && len(cert.ClientKey) > 0 {
						ck, err := tls.X509KeyPair(cert.ClientCertificate, cert.ClientKey)
						if err != nil {
							fmt.Println(err)
							return resp, fmt.Errorf("Failed to make X509 key pair: %s", err.Error())
						}
						transport = &http.Transport{
							TLSClientConfig: &tls.Config{
								RootCAs:      caCertPool,
								Certificates: []tls.Certificate{ck},
							},
						}
					} else {
						transport = &http.Transport{
							TLSClientConfig: &tls.Config{
								RootCAs: caCertPool,
							},
						}
					}
				} else {
					transport = &http.Transport{
						TLSClientConfig: &tls.Config{
							InsecureSkipVerify: true,
						},
					}
				}
			}
		}

		hub, err := newFromTransport(url, username, password, transport, registry.Log)
		if err != nil {
			fmt.Println(err)
			return resp, fmt.Errorf("Failed to creat transport: %s", err.Error())
		}

		repositories, err := hub.Repositories()
		if err != nil {
			fmt.Println(err)
			return resp, fmt.Errorf("Failed to reap catalogs: %s", err.Error())
		}
		resp.Registries = append(resp.Registries, &pb.RegistryRepositoryData_Registry{
			Name:     r.Name,
			Catalogs: make([]*pb.RegistryRepositoryData_Catalog, 0),
		})
		for _, r := range repositories {
			resp.Registries[len(resp.Registries)-1].Catalogs = append(resp.Registries[len(resp.Registries)-1].Catalogs, &pb.RegistryRepositoryData_Catalog{
				Name: r,
				Tags: make([]*pb.RegistryRepositoryData_Tag, 0),
			})
			tags, err := hub.Tags(r)
			if err != nil {
				fmt.Println(err)
				return resp, fmt.Errorf("Failed to reap tags: %s", err.Error())
			}
			for _, t := range tags {
				fmt.Println(r + ":" + t)
				resp.Registries[len(resp.Registries)-1].Catalogs[len(resp.Registries[len(resp.Registries)-1].Catalogs)].Tags = append(resp.Registries[len(resp.Registries)-1].Catalogs[len(resp.Registries[len(resp.Registries)-1].Catalogs)].Tags, &pb.RegistryRepositoryData_Tag{
					Name: t,
				})
			}
		}
	}

	return resp, nil
}

func newFromTransport(registryUrl, username, password string, transport http.RoundTripper, logf registry.LogfCallback) (*registry.Registry, error) {
	url := strings.TrimSuffix(registryUrl, "/")
	transport = registry.WrapTransport(transport, url, username, password)
	registry := &registry.Registry{
		URL: url,
		Client: &http.Client{
			Transport: transport,
		},
		Logf: logf,
	}

	if err := registry.Ping(); err != nil {
		return nil, err
	}

	return registry, nil
}

func gainCerts() (map[string]Certificates, error) {
	certs := make(map[string]Certificates)
	if s, ok := os.LookupEnv("REGISTRY_CERTS_JSON"); ok {
		if err := json.Unmarshal([]byte(s), &certs); err != nil {
			return certs, fmt.Errorf("Failed to deserization: %s", err.Error())
		}
		for k, v := range certs {
			if len(v.CA) != 0 {
				b, err := base64.StdEncoding.DecodeString(v.CA)
				if err != nil {
					return certs, fmt.Errorf("Failed to decode CA: %s", err.Error())
				}
				v.CertificateAuthority = b
				certs[k] = v
			}
			if len(v.Crt) != 0 {
				b, err := base64.StdEncoding.DecodeString(v.Crt)
				if err != nil {
					return certs, fmt.Errorf("Failed to decode crt: %s", err.Error())
				}
				v.ClientCertificate = b
				certs[k] = v
			}
			if len(v.Key) != 0 {
				b, err := base64.StdEncoding.DecodeString(v.Key)
				if err != nil {
					return certs, fmt.Errorf("Failed to decode key: %s", err.Error())
				}
				v.ClientKey = b
				certs[k] = v
			}
		}
	}
	return certs, nil
}
