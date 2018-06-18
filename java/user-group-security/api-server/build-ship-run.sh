#! /bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
echo $DIR

dbinst=$(docker ps -qaf name=mariadb)
if [[ ! -z $dbinst ]]; then
    docker stop $dbinst 2>/dev/null
	docker rm $dbinst
fi

    # -v /my/custom:/etc/mysql/conf.d
docker run -d --name mariadb -p3306:3306 \
    -v ${DIR}/db/mariadb-schema.sql:/docker-entrypoint-initdb.d/mariadb-schema.sql:ro \
    -e MYSQL_ROOT_PASSWORD=password \
    -e MYSQL_DATABASE=example \
    -e MYSQL_USER=dbuser \
    -e MYSQL_PASSWORD=dbpass \
    --restart=always mariadb:10.3

# To create db and init schema, the alternative CLI:
# docker exec -ti mariadb bash
# mysql -h 172.18.0.3 -u root --password=password < /docker-entrypoint-initdb.d/mariadb-schema.sql