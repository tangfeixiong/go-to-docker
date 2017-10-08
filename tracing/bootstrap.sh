#! /bin/bash
set -e

if [[ $# < 2 ]]; then
	echo "Usage: ${0} [<elasticserach endpoint>]
	
	<elasticserach endpoint>    For example http://localhost:9200
" > /dev/stderr
	# exit 1
fi

inst=$(docker ps -qaf label=name=metering)
if [[ ! -z $inst ]]; then
    docker stop $inst 2>/dev/null
	docker rm $inst
fi

dimg=$(docker images -qf dangling=true)
if [[ ! -z $dimg ]]; then
	docker rmi $dimg 2>/dev/null
fi

addrelasticsearch=$(docker inspect -f {{.NetworkSettings.IPAddress}} elasticsearch)

docker run -d \
	-p 12305:12305 \
    -p 12306:12306 \
	--name=tracing-collector \
	docker.io/tangfeixiong/tracing \
    jaeger-collector \
    --dependency-storage.type=elasticsearch \
    --span-storage.type=elasticsearch \
    --es.server-urls=http://${addrelasticsearch}:9200 \
    --es.username=elastic \
    --es.password=changeme

docker exec -ti tracing-collector printenv

addrcollector=$(docker inspect -f {{.NetworkSettings.IPAddress}} tracing-collector)

docker run -d \
	-p 12315:12315 \
    -p 12316:12316 \
	--name=tracing-agent \
	docker.io/tangfeixiong/tracing \
    jaeger-agent \
    --collector.host-port=${addrcollector}:14267

docker exec -ti tracing-agent printenv