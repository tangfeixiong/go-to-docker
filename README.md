# go-to-docker

Working around Docker (http://docs.docker.com/) Engine API (version 1.12, 1.13) & Registry API (v2)

HTTP/TCP Wrapper with gRPC and gRPC-Gateway (insecure mode)

## Development

Make binary
```
[vagrant@localhost go-to-docker]$ make go-install
go install -v ./
github.com/tangfeixiong/go-to-docker/pkg/ui/data/swagger
github.com/tangfeixiong/go-to-docker/pkg/server
github.com/tangfeixiong/go-to-docker/cmd
github.com/tangfeixiong/go-to-docker
```

Make dockernized
```
[vagrant@localhost go-to-docker]$ make
### snip ###
[vagrant@localhost go-to-docker]$ docker run -d \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  --privileged=true \
  -p 10052:10052 \
  -e DOCKER_CONFIG_JSON='{"auths":{"127.0.0.1:5000":{"auth":"<basicauth-base64-encoded>","email":""}}}' \
  -e REGISTRY_CERTS_JSON='{"127.0.0.1:5000":{"ca_base64":"<cacert-base64-encoded>"}}' \
  --name=go-to-docker docker.io/tangfeixiong/go-to-docker:0.1
24be50de5ed082409ba98560cff37f3ba31e1eda82ace02b98ec83ea8cce680e
```

For `DOCKER_CONFIG_JSON` environment
```
[vagrant@localhost go-to-docker]$ cat ~/.docker/config.json | tr -d '\n\t '
{"auths":{"...":{"auth":"...","email":""},...}}}
```

For `REGISTRY_CERTS_JSON` environment
```
[vagrant@localhost go-to-docker]$ echo "{\"127.0.0.1:5000\":{\"ca_base64\":\"$(base64 -w 0 /etc/docker/certs.d/127.0.0.1:5000/ca.crt)\"}}"
```


### Docker API

Examples with docker client, engine-api, go-dockerclient, and native json api

* docker client - https://github.com/docker/docker/tree/master/client
* go-dockerclient - https://github.com/fsouza/go-dockerclient
* engine-api - https://github.com/docker/engine-api
* docker-registry-client - https://github.com/heroku/docker-registry-client

### Inspired

https://github.com/grpc-ecosystem/grpc-gateway

https://github.com/philips/grpc-gateway-example

https://github.com/jcbsmpsn/golang-https-example

### Deployment

[Docker Compose](https://docs.docker.com/compose/reference/overview/)

Download from [Github Release Repositories](https://github.com/docker/compose/releases)
```
$ curl -jkSL https://github.com/docker/compose/releases/download/1.15.0/docker-compose-Linux-x86_64 -O
$ chmod +x docker-compose-Linux-x86_64
$ ./docker-compose-Linux-x86_64 --help
```

Etcd download from [Github Release Repositories](https://github.com/coreos/etcd/releases)
```
$ curl -jkSL https://github.com/coreos/etcd/releases/download/v3.2.4/etcd-v3.2.4-linux-amd64.tar.gz -O
$ tar -C /home/vagrant/go/bin/ -zxf etcd-v3.2.4-linux-amd64.tar.gz --strip-components=1 etcd-v3.2.4-linux-amd64/etcd etcd-v3.2.4-linux-amd64/etcdctl
$ etcd --version
etcd Version: 3.2.4
Git SHA: c31bec0
Go Version: go1.8.3
Go OS/Arch: linux/amd64
$ etcdctl --version
etcdctl version: 3.2.4
API version: 2
```

Example
```
[vagrant@localhost go-to-docker]$ sudo mkdir -p /srv/etc/data

[vagrant@localhost go-to-docker]$ hostname -I | awk '{print $2}'
172.17.4.50

[vagrant@localhost go-to-docker]$ sed 's/$NODE1/172.17.4.50/g' docker-compose.yml.sed > docker-compose.yml

[vagrant@localhost go-to-docker]$ docker-compose up -d
```
