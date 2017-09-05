#! /bin/bash
set -e

if [[ $# < 2 ]]; then
	echo "Usage: ${0} <docker registry> <ca/cert file> [<client certificate> <client key>]
	
	<docker registry> For example 192.168.1.99:5000
	<ca/cert file>    For example self signed cert /etc/docker/certs.d/192.168.1.99:5000/server.cert" > /dev/stderr
	exit 1
fi

inst=$(docker ps -qaf name=go-to-docker)
if [[ ! -z $inst ]]; then
    docker stop $inst 2>/dev/null
	docker rm $inst
fi

dimg=$(docker images -qf dangling=true)
if [[ ! -z $dimg ]]; then
	docker rmi $dimg 2>/dev/null
fi

conf=$(cat $HOME/.docker/config.json | tr -d '\n\t ')
certs="{\"${1}\":{\"ca_base64\":\"$(base64 -w 0 ${2})\"}}"

docker run -d \
	-v /var/run/docker.sock:/var/run/docker.sock:ro \
	--privileged=true \
	-p 10052:10052 \
	-e DOCKER_CONFIG_JSON=${conf} \
	-e REGISTRY_CERTS_JSON=${certs} \
	--name=go-to-docker \
	docker.io/tangfeixiong/go-to-docker:0.1

docker exec -ti go-to-docker printenv