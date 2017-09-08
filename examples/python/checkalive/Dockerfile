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
      name="cron" \
      namespace="stackdocker" \
      annotation='{"stackdocker.io/created-by":""}' \
      tag="alpine cron python"

RUN set -x \
    && apk add --update \
        bash \
        curl \
        wget \
        git \
        mysql-client \
        php5 \
        php5-cli \
        php5-common \
        php5-gd \
        php5-phar \
        php5-curl \
        php5-mysql \
        php5-openssl \
        php5-json \
        php5-dom \
        python \
        py-pip \
        py-mysqldb \
    && rm -rf /var/cache/apk/* \
# python 
    && pip install MySQL-python redis \
# php
    && ln -sf /usr/bin/php5 /usr/bin/php \
# php composer
    #curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/bin --filename=composer 
    && wget https://raw.githubusercontent.com/composer/getcomposer.org/master/web/installer -O - -q | php -- --quiet \
# finally
    && echo    

COPY docker-entrypoint.sh web1check/web1check.py web2checker/web2checker.py app/
WORKDIR app/

RUN touch crontab.tmp \
    && if [[ ! -e hosts.list ]]; then \
        echo "Add localhost into hosts list"; \
	    echo "127.0.0.1" > hosts.list; \
    fi \
    && mkdir -p /output \
    && touch /output/web1check.out \
# crontab into var pool
    && echo '*/2 * * * * python web1check.py>/output/web1check.out' > crontab.tmp \
    #echo '* */6 * * * /usr/bin/php /var/www/partkeepr/app/console partkeepr:cron:run' > crontab.tmp 
    #echo '0 2   * * * /usr/bin/sql_backup' >> crontab.tmp 
    && crontab crontab.tmp \
    && rm -rf crontab.tmp

# CMD ["/usr/sbin/crond", "-f", "-d", "0"]
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["-m", "2", "-a", "web1check.py", "127.0.0.1 localhost"]