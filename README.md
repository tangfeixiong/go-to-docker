# go-to-docker

Working with Docker (http://docs.docker.com/)

Support insecure gRPC and gRPC-Gateway

## Development

How to
```
[vagrant@localhost go-to-docker]$ make go-install
go install -v ./
github.com/tangfeixiong/go-to-docker/pkg/ui/data/swagger
github.com/tangfeixiong/go-to-docker/pkg/server
github.com/tangfeixiong/go-to-docker/cmd
github.com/tangfeixiong/go-to-docker
```

Dockernized
```
[vagrant@localhost go-to-docker]$ make
### snip ###
[vagrant@localhost go-to-docker]$ docker run -d -v /var/run/docker.sock:/var/run/docker.sock:ro --privileged=true -p 10052:10052 -e DOCKER_CONFIG_JSON='{"auths":{"127.0.0.1:5000":{"auth":"<basicauth-base64-encoding>","email":""}}}' --name=go-to-docker docker.io/tangfeixiong/go-to-docker:0.1
24be50de5ed082409ba98560cff37f3ba31e1eda82ace02b98ec83ea8cce680e
```

For `DOCKER_CONFIG_JSON` environment
```
[vagrant@localhost go-to-docker]$ cat ~/.docker/config.json | tr -d '\n\t '
{"auths":{"...":{"auth":"...","email":""},...}}}
```

## Docker Client API

Examples with docker client, engine-api, go-dockerclient, and native json api

* docker client - https://github.com/docker/docker/tree/master/client
* go-dockerclient - https://github.com/fsouza/go-dockerclient
* engine-api - https://github.com/docker/engine-api

## Inspired

https://github.com/grpc-ecosystem/grpc-gateway

https://github.com/philips/grpc-gateway-example