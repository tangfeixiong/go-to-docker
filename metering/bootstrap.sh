#! /bin/bash
set -e

if [[ $# < 2 ]]; then
	echo "Usage: ${0} [<meter driver> <elasticserach endpoint>]
	
	<meter driver> For example cadvisor=http://172.17.0.10:8080
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
	--name=metering-collector \
	docker.io/tangfeixiong/metering \
    collect --storage=elasticsearch=http://${addrelasticsearch}:9200

docker exec -ti metering-collector printenv

addrcadvisor=$(docker inspect -f {{.NetworkSettings.IPAddress}} cadvisor)
addrcollector=$(docker inspect -f {{.NetworkSettings.IPAddress}} metering-collector)

docker run -d \
	-p 12315:12315 \
    -p 12316:12316 \
	--name=metering-exporter \
	docker.io/tangfeixiong/metering \
    export --meter=cadvisor=http://${addrcadvisor}:8080 --collector=grpc=${addrcollector}:12305

docker exec -ti metering-exporter printenv