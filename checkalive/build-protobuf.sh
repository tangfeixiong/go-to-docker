#!/bin/bash

source `dirname $0`/dirtest.sh
echo "Generate stub at $DIR"
pushd $DIR > /dev/null

GOPATHP=/Users/fanhongling/Downloads/workspace
GOPATHD=/home/vagrant/go

	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
		pb/counselor.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		pb/counselor.proto

popd > /dev/null