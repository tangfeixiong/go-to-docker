# Inspired from https://github.com/philips/grpc-gateway-example

GOPATHP:=/Users/fanhongling/Downloads/workspace
GOPATHD:=/home/vagrant/go

all: build-bin build-docker

protoc: moby
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mpb/moby/api.proto=github.com/tangfeixiong/go-to-docker/pb/moby,Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
		pb/service.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=Mpb/moby/api.proto=github.com/tangfeixiong/go-to-docker/pb/moby,logtostderr=true:. \
		pb/service.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		pb/service.proto
	go generate ./pb

moby:
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mpb/moby/api.proto=github.com/tangfeixiong/go-to-docker/pb/moby:. \
        pb/moby/api.proto

build-bin: protoc
	@CGO_ENABLED=0 go build -v -o ./bin/gotodocker --installsuffix ./ 

build-docker:
	docker build -t docker.io/tangfeixiong/go-to-docker:0.1 ./

install-bin:
	go install -v ./

.PHONY: all protoc moby build-bin build-docker install-bin
