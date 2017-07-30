#! /bin/bash
set -e

inst=$(docker ps -qf name=go-to-docker)
if [[ ! -z $inst ]]; then
    docker stop $inst
	docker rm $inst
fi

if [[ ! -z $(docker images -qf dangling=true) ]]; then
	docker rmi $(docker images -qf dangling=true)
fi

conf=$(cat $HOME/.docker/config.json | tr -d '\n\t ')
docker run -d \
	-v /var/run/docker.sock:/var/run/docker.sock:ro \
	--privileged=true \
	-p 10052:10052 \
	-e DOCKER_CONFIG_JSON=${conf} \
	--name=go-to-docker \
	docker.io/tangfeixiong/go-to-docker:0.1

docker exec -ti go-to-docker printenv