# Inspired from https://github.com/philips/grpc-gateway-example

GOPATHP?=/Users/fanhongling/Downloads/workspace
GOPATHD?=/home/vagrant/go

IMG_NS?=docker.io/tangfeixiong
IMG_REPO?=go-to-docker
IMG_TAG?=0.1
GIT_COMMIT=$(shell date +%y%m%d%H%M)-git$(shell git rev-parse --short=7 HEAD)
REGISTRY_HOST?=172.17.4.50:5000

all: protoc-grpc docker-push

protoc-grpc: protoc-moby
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/ \
		-Ivendor/github.com/gogo/googleapis/ \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=Mpb/moby/moby.proto=github.com/tangfeixiong/go-to-docker/pb/moby,\
Mpb/moby/container/container.proto=github.com/tangfeixiong/go-to-docker/pb/moby/container,\
Mpb/moby/network/network.proto=github.com/tangfeixiong/go-to-docker/pb/moby/network,\
Mpb/moby/filters/filters.proto=github.com/tangfeixiong/go-to-docker/pb/moby/filters,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
plugins=grpc:. \
		pb/service.proto
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/ \
		-Ivendor/github.com/gogo/googleapis/ \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=Mpb/moby/moby.proto=github.com/tangfeixiong/go-to-docker/pb/moby,\
Mpb/moby/container/container.proto=github.com/tangfeixiong/go-to-docker/pb/moby/container,\
Mpb/moby/network/network.proto=github.com/tangfeixiong/go-to-docker/pb/moby/network,\
Mpb/moby/filters/filters.proto=github.com/tangfeixiong/go-to-docker/pb/moby/filters,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
logtostderr=true:. \
		pb/service.proto
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway \
		-Ivendor/github.com/gogo/googleapis \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=Mpb/moby/moby.proto=github.com/tangfeixiong/go-to-docker/pb/moby,\
Mpb/moby/container/container.proto=github.com/tangfeixiong/go-to-docker/pb/moby/container,\
Mpb/moby/network/network.proto=github.com/tangfeixiong/go-to-docker/pb/moby/network,\
Mpb/moby/filters/filters.proto=github.com/tangfeixiong/go-to-docker/pb/moby/filters,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
logtostderr=true:. \
		pb/service.proto
	
	# Workaround for https://github.com/grpc-ecosystem/grpc-gateway/issues/229.
	# sed -i.bak "s/empty.Empty/types.Empty/g" pb/service.pb.gw.go && rm pb/service.pb.gw.go.bak

	go generate ./pb

protoc-moby:
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway \
		-Ivendor/github.com/gogo/googleapis \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mpb/moby/blkiodev/blkiodev.proto=github.com/tangfeixiong/go-to-docker/pb/moby/blkiodev:. \
        pb/moby/blkiodev/blkiodev.proto
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway \
		-Ivendor/github.com/gogo/googleapis \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mpb/moby/filters/filters.proto=github.com/tangfeixiong/go-to-docker/pb/moby/filters:. \
        pb/moby/filters/filters.proto
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway \
		-Ivendor/github.com/gogo/googleapis \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mpb/moby/mount/mount.proto=github.com/tangfeixiong/go-to-docker/pb/moby/mount:. \
        pb/moby/mount/mount.proto
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway \
		-Ivendor/github.com/gogo/googleapis \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mpb/moby/nat/nat.proto=github.com/tangfeixiong/go-to-docker/pb/moby/nat:. \
        pb/moby/nat/nat.proto
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway \
		-Ivendor/github.com/gogo/googleapis \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mpb/moby/network/network.proto=github.com/tangfeixiong/go-to-docker/pb/moby/network:. \
        pb/moby/network/network.proto
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway \
		-Ivendor/github.com/gogo/googleapis \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mpb/moby/units/units.proto=github.com/tangfeixiong/go-to-docker/pb/moby/units:. \
        pb/moby/units/units.proto
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway \
		-Ivendor/github.com/gogo/googleapis \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mpb/moby/blkiodev/blkiodev.proto=github.com/tangfeixiong/go-to-docker/pb/moby/blkiodev,\
Mpb/moby/container/container.proto=github.com/tangfeixiong/go-to-docker/pb/moby/container,\
Mpb/moby/filters/filters.proto=github.com/tangfeixiong/go-to-docker/pb/moby/filters,\
Mpb/moby/mount/mount.proto=github.com/tangfeixiong/go-to-docker/pb/moby/mount,\
Mpb/moby/nat/nat.proto=github.com/tangfeixiong/go-to-docker/pb/moby/nat,\
Mpb/moby/network/network.proto=github.com/tangfeixiong/go-to-docker/pb/moby/network,\
Mpb/moby/units/units.proto=github.com/tangfeixiong/go-to-docker/pb/moby/units:. \
        pb/moby/container/container.proto
	protoc -I/usr/local/include -I. \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway \
		-Ivendor/github.com/gogo/googleapis \
		-Ivendor/ \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mpb/moby/moby.proto=github.com/tangfeixiong/go-to-docker/pb/moby,\
Mpb/moby/blkiodev/blkiodev.proto=github.com/tangfeixiong/go-to-docker/pb/moby/blkiodev,\
Mpb/moby/container/container.proto=github.com/tangfeixiong/go-to-docker/pb/moby/container,\
Mpb/moby/filters/filters.proto=github.com/tangfeixiong/go-to-docker/pb/moby/filters,\
Mpb/moby/mount/mount.proto=github.com/tangfeixiong/go-to-docker/pb/moby/mount,\
Mpb/moby/nat/nat.proto=github.com/tangfeixiong/go-to-docker/pb/moby/nat,\
Mpb/moby/network/network.proto=github.com/tangfeixiong/go-to-docker/pb/moby/network,\
Mpb/moby/units/units.proto=github.com/tangfeixiong/go-to-docker/pb/moby/units:. \
        pb/moby/moby.proto

gen-statik:
	@if [[ ! -f $GOBIN/statik ]]; then \
		go install ./vendor/github.com/rakyll/statik; \
	fi
	#@cp ./pb/service.swagger.json ./third_party/OpenAPI/
	@go generate ./pkg/ui

go-install:
	go install -v ./

go-build:
	@CGO_ENABLED=0 go build -a -v -o ./bin/gotodocker --installsuffix cgo -ldflags "-s" ./
	#@CGO_ENABLED=0 go build -v -o ./bin/gotodocker -tags netgo -installsuffix netgo -ldflags "-s" ./

docker-build: go-build
	docker build -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG) ./

docker-push: docker-build
	docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

docker-cgo:
	go build -v -a -o ./bin/gotodocker ./
	docker build -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)-$(GIT_COMMIT) -f Dockerfile.cgo ./
	docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

docker-run:
	$(info $(shell ./bootstrap.sh $(REGISTRY_HOST) "/etc/docker/certs.d/$(REGISTRY_HOST)/ca.crt"))

.PHONY: all protoc-grpc protoc-moby go-install go-build docker-build docker-push docker-cgo docker-run
