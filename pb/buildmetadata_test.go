/*
  [vagrant@kubedev-172-17-4-59 go-to-docker]$ GOPATH=/Users/fanhongling/Downloads/workspace/:/Users/fanhongling/go go test -test.v -run BuildMetadata ./pb/
*/

package pb

import (
	"testing"
)

var (
	dockerfile []byte = []byte(`
#
FROM scratch
# FROM alpine:3.4
# FROM docker.io/gliderlabs/alpine
### MAINTAINER tangfeixiong <tangfx128@gmail.com>
LABEL maintainer="tangfeixiong" mailto="tangfx128@gmail.com"

# RUN apk add ospd-netstat --update --repository http://dl-6.alpinelinux.org/alpine/edge/community && rm -rf /var/cache/apk/*
COPY gofileserver /
WORKDIR /mnt
VOLUME ["/mnt"]
CMD ["/gofileserver"]
EXPOSE 48080
`)

	gitMetadata []byte = []byte(`# ** metadata:git **
#

https://github.com/tangfeixiong/nta#:docs
#
`)
)

func TestBuildMetadataDockerfile(t *testing.T) {
	metadata := NewImageBuildMetadata(dockerfile)
	ok := metadata.IsDockerfile()
	if ok {
		t.Log("OK")
	} else {
		t.FailNow()
	}
}

func TestBuildMetadataGit(t *testing.T) {
	metadata := NewImageBuildMetadata([]byte(gitMetadata))
	ok := metadata.IsGitReopsitory()
	if ok {
		t.Log("OK")
	} else {
		t.FailNow()
	}
}
