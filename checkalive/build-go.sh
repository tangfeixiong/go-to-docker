#!/bin/sh

DIR=`dirname $0`
source ${DIR}/dirtest.sh
pushd $DIR

#if [[ 0 -lt $# ]]; then
case $1 in
  cgo|--cgo)
	go build -v -o ./bin/target-cm ./
	echo "build with CGO_ENABLED=1"
  ;;
  *)
	CGO_ENABLED=0 go build -v -a -o ./bin/target-cm --installsuffix cgo -ldflags "-s" ./
	echo "build with CGO_ENABLED=0"
  ;;
esac
#fi

popd