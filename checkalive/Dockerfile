# Inspired by
#   https://gist.github.com/mhubig/a01276e17496e9fd6648cf426d9ceeec
#   https://github.com/gliderlabs/docker-alpine/issues/120
#   https://github.com/xordiv/docker-alpine-cron
#   https://github.com/dcharbonnier/alpine-docker-cron   
#   https://github.com/renskiy/cron-docker-image/tree/master/alpine
#   https://forums.docker.com/t/how-to-run-a-cron-job-inside-a-container-alpine/7759
#   https://github.com/matriphe/docker-alpine-php/blob/master/5.6/CLI/Dockerfile
#   https://github.com/wodby/php
#   https://getcomposer.org/doc/faqs/how-to-install-composer-programmatically.md
FROM alpine
LABEL maintainer='tangfeixiong <tangfx128@gmail.com>' \
      project="https://github.com/tangfeixiong/go-to-docker" \
      name="tickerjobcm" \
      namespace="stackdocker" \
      annotation='{"stackdocker.io/created-by":"n/a"}' \
      tag="alpine python"

RUN set -x \
    && apk add --update \
        bash \
		ca-certificates \
        curl \
        wget \
        git \
        mysql-client \
		openssl-dev \
		libxml2-dev \
		libxslt-dev \
		libffi-dev \
        python \
        py-pip \
        py-mysqldb \
		python-dev \
		build-base \
    && rm -rf /var/cache/apk/* \
# python 
    && pip install \
		requests \
	    MySQL-python \
		redis \
		MultipartPostHandler \
		scp \
		pyquery \
		multipart_encode \
		poster \
# finally
    && echo    

COPY bin/checkalive program/ /

# default settings
#ENV CHECKALIVE_GRPC_PORT="10061" \
#    CHECKALIVE_HTTP_PORT="10062" \
#    CHECKALIVE_PACKAGE_HOME="examples/python/checkalive" \
#    DATABUS_REDIS_SENTINEL_HOSTS="127.0.0.1:26379" \
#    DATABUS_REDIS_HOST="127.0.0.1:6379" \
#    DATABUS_REDIS_DB="15" \
#    DATABUS_ETCD_HOSTS="127.0.0.1:2379" \
#    DATABUS_MYSQL_HOST="127.0.0.1:3306" \
#    DATABUS_GNATSD_HOSTS="127.0.0.1:4222" \
#    DATABUS_KAFKA_HOSTS="127.0.0.1:9092" \
#    DATABUS_ZOOKEEPER_HOSTS="127.0.0.1:2181" \
#    DATABUS_RABBIT_HOST= "127.0.0.1:5672" \  
#    CHECK_MESSAGE_PUBSUB="checkalive" \
#    CHECK_CM_CACHE="checkalive" \
#    DATABUS_PRIORITY_CMDB="sentinel redis etcd mysql" \
#    DATABUS_PRIORITY_MQ="sentinel redis gnatsd kafka rabbit"

VOLUME ["/examples"]

EXPOSE 10061 10062

ENTRYPOINT ["/checkalive", "serve"]
CMD ["--v=2", "--logtostderr=true"]
