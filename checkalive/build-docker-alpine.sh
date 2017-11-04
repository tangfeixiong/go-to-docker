#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

source `dirname $0`/dirtest.sh

IMG_NS=docker.io/tangfeixiong
IMG_REPO=target-cm
IMG_TAG=0.1
GIT_COMMIT=$(date +%y%m%dT%H%M%SZ)-git_$(git rev-parse --short=7 HEAD)

CTX=$(mktemp -d)
cp -r $DIR/Dockerfile $DIR/bin $CTX

# SRC_DIR=$DIR/../examples/php/...
# DEST_DIR=$CTX/program/examples/php/...
# mkdir -p $DEST_DIR
# cp -r $SRC_DIR/... $DEST_DIR

SRC_DIR=$DIR/../examples/python/check-alive
DEST_DIR=$CTX/program/examples/python/check-alive
mkdir -p $DEST_DIR
cp -r $SRC_DIR/awd*  $DEST_DIR
cp -r $SRC_DIR/web*  $DEST_DIR

if [[ $# -gt 0 ]]; then
case $1 in
  build|--build)
	docker build --no-cache --force-rm -t $IMG_NS/$IMG_REPO:$IMG_TAG -f $CTX/Dockerfile $CTX
  ;;
  push|--push)
	docker build --no-cache --force-rm -t $IMG_NS/$IMG_REPO:$IMG_TAG -f $CTX/Dockerfile $CTX
	docker push $IMG_NS/$IMG_REPO:$IMG_TAG
  ;;
  *)
    echo "Usage: $0 [build|--build|push|--push]
        Warning: using cache..."
	docker build -t $IMG_NS/$IMG_REPO:$IMG_TAG -f $CTX/Dockerfile $CTX
  ;;
esac	
else
	docker build -t $IMG_NS/$IMG_REPO:$IMG_TAG-$GIT_COMMIT -f $CTX/Dockerfile $CTX
fi

rm -rf $CTX