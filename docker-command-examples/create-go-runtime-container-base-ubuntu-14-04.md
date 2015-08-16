#Create a basic go 1.4.2 runtime container from Ubuntu 14.04

the container is most skeleton to meet with go application production

its Dockerfile is from [Docker-library golang repo resided in Github](https://github.com/docker-library/golang/tree/master/1.4)

##Get Dockerfile

*Use __git__ client to clone into local repository

First create local repository dir, for example in my lab

>#cd /opt/tfx/docker

>#mkdir -p github.com/docker-library

>#cd github.com/docker-library

Then run `git clone` command

>#git clone https://github.com/docker-library/golang.git

*Use edit tool to create a file copy

Like following

>#mkdir -p golang-Dockerfile-exercise/1.4/onbuild

>#cd golang-Dockerfile-exercise/1.4

On broswer, copy content from golang/1.4/onbuild/Dockerfile, open edit tool to paste into

>#vi onbuild/Dockerfile

>in editor, press __i__ switching into edit mode, press __ctrl-insert__ to copy, press __Esc__ switching into command mode, press __:wq__ save and quit

As well as to

>#vi go-wrapper

*Use download tool 

For example, asume current work dir is golang-Dockerfile-exercise/1.4 according previous command

>#wget -O onbuild/Dockerfile https://raw.githubusercontent.com/docker-library/golang/master/1.4/onbuild/Dockerfile

>#wget https://raw.githubusercontent.com/docker-library/golang/master/1.4/go-wrapper

##Build image

Please refer to [Dockerfile reference](https://docs.docker.com/reference/builder/) for full knowledge

For example, in my lab

>#docker build onbuild/

>or omit tag

>#docker build -t tangfeixiong/ubuntu-go-runtime-exercise onbuild/

>the image can be tagged laterly, here the tag is composed with my Docker Hub account and expected repository name, for how to sign up a Docker Hub account, show in this link [Your docker Hub account](https://docs.docker.com/docker-hub/accounts/)

*Waiting builder download base image from Hub

As first line in Dockfile show __FROM golang:1.4.2__
the golang repository is created as Docker official, the tag 1.4.2 means go binary version

*Then initialize container

The second __RUN__ instruction line to create a go source dir, because the base image has already setup $GOPATH, e.g /go

The third __WORKDIR__ instruction line to change working directory into /go/src/app while running this image for container

The fourth __CMD__ instruction line to set a default exection environment for fire a app or command onto container, the exec env is called __go-wrapper__ bash, the download source code show how it work

*The __onbuild__ instruction

##Image information

After created, use docker command to understand new image

*Get image id

Like following

>#docker images

>or if already have more, use pipe (grep, head, less ...)

>#docker images | head

>REPOSITORY                                               TAG                 IMAGE ID            CREATED             VIRTUAL SIZE

>tangfeixiong/exercise-golang-1-4-onbuild                 latest              da07533394f8        28 hours ago        517.3 MB


*Get build history

Like

>#docker history 

>IMAGE               CREATED             CREATED BY                                      SIZE                COMMENT

>da07533394f8        28 hours ago        /bin/sh -c #(nop) ONBUILD RUN go-wrapper inst   0 B

>204200b31c9c        28 hours ago        /bin/sh -c #(nop) ONBUILD RUN go-wrapper down   0 B

>96b6ac0eb149        28 hours ago        /bin/sh -c #(nop) ONBUILD COPY . /go/src/app    0 B

>420b1ebc2b16        28 hours ago        /bin/sh -c #(nop) CMD ["go-wrapper" "run"]      0 B

>45988f2f31cb        28 hours ago        /bin/sh -c #(nop) WORKDIR /go/src/app           0 B

>805bde027917        28 hours ago        /bin/sh -c mkdir -p /go/src/app                 0 B

>124e2127157f        4 weeks ago         /bin/sh -c #(nop) COPY file:56695ddefe9b0bd83   2.481 kB

>69c177f0c117        4 weeks ago         /bin/sh -c #(nop) WORKDIR /go                   0 B

>141b650c3281        4 weeks ago         /bin/sh -c #(nop) ENV PATH=/go/bin:/usr/src/g   0 B

>8fb45e60e014        4 weeks ago         /bin/sh -c #(nop) ENV GOPATH=/go                0 B

>63e9d2557cd7        4 weeks ago         /bin/sh -c mkdir -p /go/src /go/bin && chmod    0 B

>b279b4aae826        4 weeks ago         /bin/sh -c #(nop) ENV PATH=/usr/src/go/bin:/u   0 B

>d86979befb72        4 weeks ago         /bin/sh -c cd /usr/src/go/src && ./make.bash    97.4 MB

>8ddc08289e1a        4 weeks ago         /bin/sh -c curl -sSL https://golang.org/dl/go   39.69 MB

>8d38711ccc0d        4 weeks ago         /bin/sh -c #(nop) ENV GOLANG_VERSION=1.4.2      0 B

>0f5121dd42a6        4 weeks ago         /bin/sh -c apt-get update && apt-get install    88.32 MB

>607e965985c1        4 weeks ago         /bin/sh -c apt-get update && apt-get install    122.3 MB

>1ff9f26f09fb        4 weeks ago         /bin/sh -c apt-get update && apt-get install    44.36 MB

>9a61b6b1315e        4 weeks ago         /bin/sh -c #(nop) CMD ["/bin/bash"]             0 B

>902b87aaaec9        4 weeks ago         /bin/sh -c #(nop) ADD file:e1dd18493a216ecd0c   125.2 MB

*Watch image contruct data

Like

>#docker inspect da07533394f8

## Run container with image

Like

>#docker run -ti da07533394f8 /bin/bash

Prompt indicate now is in container 

>root@ae56fe85c270:/go/src/app# ls /usr/local/bin

>go-wrapper

The base image has already installed go-wrapper, as well go runtime 

>root@ae56fe85c270:/go/src/app# go version

>go version go1.4.2 linux/amd64

>root@ae56fe85c270:/go/src/app# ls /usr/src/go/bin/

>go  gofmt

Also, defined $GOPATH environment variable

>root@ae56fe85c270:/go/src/app# echo $GOPATH

>/go

Return host

>root@ae56fe85c270:/go/src/app# exit

>exit

## Re-start container

The container stopped when exit, because it has not been running a daemon

Want to access again, first get container id

>#docker ps -l

>CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                      PORTS               NAMES

>ae56fe85c270        da07533394f8        "/bin/bash"         2 minutes ago       Exited (0) 12 seconds ago                       trusting_mcclintock

Start it

>#docker start ae56fe85c270

Execute bash for enter container

>#docker exec -ti ae56fe85c270 /bin/bash
