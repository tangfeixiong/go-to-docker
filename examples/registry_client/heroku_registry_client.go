/*
  Inspired by:
    https://github.com/heroku/docker-registry-client/blob/master/registry/registry.go
    https://github.com/jcbsmpsn/golang-https-example/blob/master/https_client.go
*/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	//"github.com/docker/distribution/digest"
	//"github.com/docker/distribution/manifest"
	//"github.com/docker/libtrust"
	"github.com/heroku/docker-registry-client/registry"
)

func main() {
	url := "https://172.17.4.50:5000/" // "https://registry-1.docker.io/"
	username := "admin"                // anonymous
	password := "password"             // anonymous

	caCert, err := ioutil.ReadFile("server.cert")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	//	client := &http.Client{
	//		Transport: &http.Transport{
	//			TLSClientConfig: &tls.Config{
	//				RootCAs: caCertPool,
	//			},
	//		},
	//	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
		},
	}

	//	hub, err := registry.New(url, username, password)
	hub, err := newFromTransport(url, username, password, transport, registry.Log)
	if err != nil {
		fmt.Println(err)
		return
	}

	repositories, err := hub.Repositories()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, r := range repositories {
		fmt.Println(r)
		tags, err := hub.Tags(r)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, t := range tags {
				fmt.Println(r + ":" + t)
			}
		}
		fmt.Println()
	}

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
