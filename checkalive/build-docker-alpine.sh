#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

source `dirname $0`/dirtest.sh

IMG_NS=docker.io/tangfeixiong
IMG_REPO=target-cm
IMG_TAG=0.1
GIT_COMMIT=$(date +%y%m%d%H%M)-git_$(git rev-parse --short=7 HEAD)

CTX=$(mktemp -d)
cp -r $DIR/Dockerfile $DIR/bin $CTX

# mkdir -p $CTX/program/examples/php/...
# cp -r $DIR/../examples/php/... $CTX/program/examples/...

mkdir -p $CTX/program/examples/python/checkalive
cp -r $DIR/../examples/python/checkalive/awd1-4 $CTX/program/examples/python/checkalive
cp -r $DIR/../examples/python/checkalive/awd1-8 $CTX/program/examples/python/checkalive
cp -r $DIR/../examples/python/checkalive/web1-2 $CTX/program/examples/python/checkalive
cp -r $DIR/../examples/python/checkalive/web1check $CTX/program/examples/python/checkalive
cp -r $DIR/../examples/python/checkalive/web2check $CTX/program/examples/python/checkalive

if [[ $# -gt 0 ]]; then
case $1 in
  push|--push)
	docker build -t $IMG_NS/$IMG_REPO:$IMG_TAG -f $CTX/Dockerfile $CTX
	docker push $IMG_NS/$IMG_REPO:$IMG_TAG
  ;;
  *)
    echo "Usage: $0 [push|--push]"
	docker build -t $IMG_NS/$IMG_REPO:$IMG_TAG-$GIT_COMMIT -f $CTX/Dockerfile $CTX
  ;;
esac	
else
	docker build -t $IMG_NS/$IMG_REPO:$IMG_TAG-$GIT_COMMIT -f $CTX/Dockerfile $CTX
fi

rm -rf $CTX