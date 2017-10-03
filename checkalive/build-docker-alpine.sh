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

# SRC_DIR=$DIR/../examples/php/...
# DEST_DIR=$CTX/program/examples/php/...
# mkdir -p $DEST_DIR
# cp -r $SRC_DIR/... $DEST_DIR

SRC_DIR=$DIR/../examples/python/check-alive
DEST_DIR=$CTX/program/examples/python/check-alive
mkdir -p $DEST_DIR
cp -r $SRC_DIR/awd10_nothing  $DEST_DIR
cp -r $SRC_DIR/awd11_maccms  $DEST_DIR
cp -r $SRC_DIR/awd12_phpsqllitecms  $DEST_DIR
cp -r $SRC_DIR/awd1_lemon_cms  $DEST_DIR
cp -r $SRC_DIR/awd1_xmanweb2  $DEST_DIR
cp -r $SRC_DIR/awd2_daydayweb  $DEST_DIR
cp -r $SRC_DIR/awd2_dynpage  $DEST_DIR
cp -r $SRC_DIR/awd3_electronics  $DEST_DIR
cp -r $SRC_DIR/awd3_shadow  $DEST_DIR
cp -r $SRC_DIR/awd4_chinaz  $DEST_DIR
cp -r $SRC_DIR/awd4_tomcat  $DEST_DIR
cp -r $SRC_DIR/awd5_babyblog  $DEST_DIR
cp -r $SRC_DIR/awd5_gracer  $DEST_DIR
cp -r $SRC_DIR/awd6_cms  $DEST_DIR
cp -r $SRC_DIR/awd7_upload  $DEST_DIR
cp -r $SRC_DIR/awd8_blog  $DEST_DIR
cp -r $SRC_DIR/awd9_money  $DEST_DIR
cp -r $SRC_DIR/awd1-4  $DEST_DIR
cp -r $SRC_DIR/awd1-8  $DEST_DIR
cp -r $SRC_DIR/web1-1  $DEST_DIR
cp -r $SRC_DIR/web2-1  $DEST_DIR
cp -r $SRC_DIR/web1check  $DEST_DIR
cp -r $SRC_DIR/web2check  $DEST_DIR

if [[ $# -gt 0 ]]; then
case $1 in
  build|--build)
	docker build -t $IMG_NS/$IMG_REPO:$IMG_TAG -f $CTX/Dockerfile $CTX
  ;;
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