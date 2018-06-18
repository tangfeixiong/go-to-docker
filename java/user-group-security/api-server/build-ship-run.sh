#! /bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
# echo $DIR

tgt=$(ls $DIR/target/api*.jar)
if [[ -z $tgt ]]; then
    echo "Binary not found (e.g. api-xxx.jar), build with Maven first" >/dev/stderr
    exit 1;
fi

img=$(docker images -aq docker.io/tangfeixiong/ugs)

docker build --force-rm --no-cache \
    -t docker.io/tangfeixiong/ugs \
    --build-arg jarTgt=/target/${tgt} \
    -f $DIR/Dockerfile.centos7 $DIR

if [[ ! -z $dbinst ]]; then
    docker rmi $img
fi

dbinst=$(docker ps -qaf name=mariadb)
if [[ ! -z $dbinst ]]; then
    docker stop $dbinst 2>/dev/null
	docker rm $dbinst
fi

vnet=$(docker network ls --filter=name=user_organization_security --quiet)
if [[ -z $vnet ]]; then
    docker network create user_organization_security
fi

    # -v /my/custom:/etc/mysql/conf.d
docker run -d --network=user_organization_security \
    --name mariadb -p3306:3306 \
    -v ${DIR}/db/mariadb-schema.sql:/docker-entrypoint-initdb.d/mariadb-schema.sql:ro \
    -e MYSQL_ROOT_PASSWORD=password \
    -e MYSQL_DATABASE=example \
    -e MYSQL_USER=dbuser \
    -e MYSQL_PASSWORD=dbpass \
    --restart=always mariadb:10.3

# To create db and init schema, the alternative CLI:
# docker exec -ti mariadb bash
# mysql -h 172.18.0.3 -u root --password=password < /docker-entrypoint-initdb.d/mariadb-schema.sql

apic=$(docker ps -qaf name=ugs)
if [[ ! -z $apic ]]; then
    docker stop $apic 2>/dev/null
	docker rm $apic
fi

docker run -d --network=user_organization_security \
    --name ugs -p18080:8080 \
    -e SPRING_DATASOURCE_URL=jdbc:mysql://mariadb:3306/example
    -e SPRING_DATASOURCE_USERNAME=dbuser
    -e SPRING_DATASOURCE_PASSWORD=dbpass
    --restart=no docker.io/tangfeixong/ugs

