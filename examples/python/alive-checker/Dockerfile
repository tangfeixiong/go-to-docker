# Inspired by
#   https://gist.github.com/mhubig/a01276e17496e9fd6648cf426d9ceeec
#   https://github.com/gliderlabs/docker-alpine/issues/120
#   https://github.com/xordiv/docker-alpine-cron
#   https://github.com/dcharbonnier/alpine-docker-cron   
#   https://github.com/renskiy/cron-docker-image/tree/master/alpine
#   https://forums.docker.com/t/how-to-run-a-cron-job-inside-a-container-alpine/7759
#   https://getcomposer.org/doc/faqs/how-to-install-composer-programmatically.md
FROM alpine
LABEL maintainer='tangfeixiong <tangfx128@gmail.com>' \
      project="https://github.com/tangfeixiong/go-to-docker" \
      name="cron" \
      namespace="stackdocker" \
      annotation='{"stackdocker.io/created-by":""}' \
      tag = "alpine cron python"

RUN set -x \
    && apk add --update bash wget curl git php php-curl php-openssl php-json php-phar php-dom python py-pip mysql-client \
    && rm -rf /var/cache/apk/* \
# python 
    && pip install MySQL-python redis \
# php composer
    #curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/bin --filename=composer 
    && wget https://raw.githubusercontent.com/composer/getcomposer.org/master/web/installer -O - -q | php -- --quiet \
# finally
    && echo    

COPY app.sh *.py app/
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