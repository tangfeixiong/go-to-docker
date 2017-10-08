
GOPATHP?=/Users/fanhongling/Downloads/workspace
GOPATHD?=/home/vagrant/go

IMG_NS?=docker.io/tangfeixiong
IMG_REPO?=tracing
IMG_TAG?=latest

GIT_COMMIT=$(shell date +%y%m%d%H%M)-git$(shell git rev-parse --short=7 HEAD)
DOCKER_FILE?=Dockerfile.alpine

REGISTRY_HOST?=172.17.4.50:5000

all: protoc-grpc docker-push

protoc-grpc:
	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/gogo/protobuf/protobuf \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/protobuf/google/protobuf,Mpb/metric.proto=github.com/tangfeixiong/go-to-docker/tracing/pb,plugins=grpc:. \
		pb/agent.proto pb/span.proto
	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		pb/agent.proto

go-install:
	@go install -v ./

go-build:
	@CGO_ENABLED=0 go build -a -v -o ./bin/tracing --installsuffix cgo -ldflags "-s" ./
	#@CGO_ENABLED=0 go build -v -o ./bin/tracing -tags netgo -installsuffix netgo -ldflags "-s" ./

docker-build: go-build
	#@docker build --no-cache --force-rm -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG) ./
	@docker build --no-cache --force-rm -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG) -f $(DOCKER_FILE) ./

docker-push: docker-build
	@docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

docker-build-agent:
	@sed 's/^# \(CMD \["jaeger-agent".*\]\)/\1/; s/^\(CMD \["jaeger-collector".*\]\)/# \1/; s/^# \(EXPOSE 5775.*\)/\1/; s/^\(EXPOSE 14267.*\)/# \1/' $(DOCKER_FILE) > Dockerfile.agent
	@docker build --no-cache --force-rm -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG) -f Dockerfile.agent ./

docker-build-collector:
	@sed 's/^\(CMD \["jaeger-agent".*\]\)/# \1/; s/^# \(CMD \["jaeger-collector".*\]\)/\1/; s/^\(EXPOSE 5775.*\)/# \1/; s/^# \(EXPOSE 14267.*\)/\1/' $(DOCKER_FILE) > Dockerfile.collector
	@docker build --no-cache --force-rm -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG) -f Dockerfile.collector ./ 

docker-cgo:
	@go build -v -a -o ./bin/tracing ./
	@docker build -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)-$(GIT_COMMIT) -f Dockerfile.cgo ./
	@docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

docker-run:
	$(info $(shell ./bootstrap.sh cadvisor=http://localhost:8080 elasticsearch=http://localhost:9200))

docker-compose:
	@docker-compose --project-name=tracing up --force-recreate

.PHONY: all protoc-grpc go-install go-build docker-build docker-push docker-build-agent docker-build-collector docker-cgo docker-run
