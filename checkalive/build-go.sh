#!/bin/sh

DIR=`dirname $0`
source ${DIR}/dirtest.sh
pushd $DIR

#if [[ 0 -lt $# ]]; then
case $1 in
  cgo|--cgo)
	echo "CGO_ENABLED"
	go build -v -o ./bin/checkalive ./
  ;;
  *)
	echo "build with static link (CGO_ENABLED=0)"
	CGO_ENABLED=0 go build -v -a -o ./bin/checkalive --installsuffix cgo -ldflags "-s" ./
  ;;
esac
#fi

popd