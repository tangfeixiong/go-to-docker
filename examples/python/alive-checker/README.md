Test
```
[vagrant@bogon alive-checker]$ docker build -t tangfeixiong/target-alive-checker .
Sending build context to Docker daemon 17.41 kB
Step 1 : FROM alpine
 ---> 7328f6f8b418
Step 2 : LABEL maintainer 'tangfeixiong <tangfx128@gmail.com>' project "https://github.com/tangfeixiong/go-to-docker" name "cron" namespace "stackdocker" annotation '{"stackdocker.io/created-by":""}' tag "alpine cron python"
 ---> Running in b88a0af951b9
 ---> 84039785f7c7
Removing intermediate container b88a0af951b9
Step 3 : RUN set -x     && apk add --update         bash         curl         wget         git         mysql-client         php5         php5-cli         php5-common         php5-gd         php5-phar         php5-curl         php5-mysql         php5-openssl         php5-json         php5-dom         python         py-pip         py-mysqldb     && rm -rf /var/cache/apk/*     && pip install MySQL-python redis     && ln -sf /usr/bin/php5 /usr/bin/php     && wget https://raw.githubusercontent.com/composer/getcomposer.org/master/web/installer -O - -q | php -- --quiet     && echo
 ---> Running in 64950711cc4e
+ apk add --update bash curl wget git mysql-client php5 php5-cli php5-common php5-gd php5-phar php5-curl php5-mysql php5-openssl php5-json php5-dom python py-pip py-mysqldb
fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/community/x86_64/APKINDEX.tar.gz
(1/39) Installing ncurses-terminfo-base (6.0-r8)
(2/39) Installing ncurses-terminfo (6.0-r8)
(3/39) Installing ncurses-libs (6.0-r8)
(4/39) Installing readline (6.3.008-r5)
(5/39) Installing bash (4.3.48-r1)
Executing bash-4.3.48-r1.post-install
(6/39) Installing ca-certificates (20161130-r2)
(7/39) Installing libssh2 (1.8.0-r1)
(8/39) Installing libcurl (7.55.0-r0)
(9/39) Installing curl (7.55.0-r0)
(10/39) Installing expat (2.2.0-r1)
(11/39) Installing pcre (8.41-r0)
(12/39) Installing git (2.13.5-r0)
(13/39) Installing mariadb-common (10.1.26-r0)
(14/39) Installing mariadb-client (10.1.26-r0)
(15/39) Installing mysql-client (10.1.26-r0)
(16/39) Installing php5-common (5.6.31-r0)
(17/39) Installing libxml2 (2.9.4-r4)
(18/39) Installing php5-cli (5.6.31-r0)
(19/39) Installing php5 (5.6.31-r0)
(20/39) Installing php5-curl (5.6.31-r0)
(21/39) Installing php5-dom (5.6.31-r0)
(22/39) Installing libbz2 (1.0.6-r5)
(23/39) Installing libpng (1.6.29-r1)
(24/39) Installing freetype (2.7.1-r1)
(25/39) Installing libjpeg-turbo (1.5.1-r0)
(26/39) Installing php5-gd (5.6.31-r0)
(27/39) Installing php5-json (5.6.31-r0)
(28/39) Installing php5-mysql (5.6.31-r0)
(29/39) Installing php5-openssl (5.6.31-r0)
(30/39) Installing php5-phar (5.6.31-r0)
(31/39) Installing libffi (3.2.1-r3)
(32/39) Installing gdbm (1.12-r0)
(33/39) Installing sqlite-libs (3.18.0-r0)
(34/39) Installing python2 (2.7.13-r1)
(35/39) Installing mariadb-client-libs (10.1.26-r0)
(36/39) Installing py-mysqldb (1.2.5-r0)
(37/39) Installing py-setuptools (33.1.1-r1)
(38/39) Installing py2-pip (9.0.1-r1)
(39/39) Installing wget (1.19.1-r2)
Executing busybox-1.26.2-r5.trigger
Executing ca-certificates-20161130-r2.trigger
OK: 129 MiB in 50 packages
+ rm -rf /var/cache/apk/APKINDEX.24d64ab1.tar.gz /var/cache/apk/APKINDEX.84815163.tar.gz
+ pip install MySQL-python redis
Requirement already satisfied: MySQL-python in /usr/lib/python2.7/site-packages
Collecting redis
  Downloading redis-2.10.6-py2.py3-none-any.whl (64kB)
Installing collected packages: redis
Successfully installed redis-2.10.6
+ ln -sf /usr/bin/php5 /usr/bin/php
+ php -- --quiet
+ wget https://raw.githubusercontent.com/composer/getcomposer.org/master/web/installer -O - -q
Signature mismatch, could not verify the phar file integrity
Signature mismatch, could not verify the phar file integrity
Some settings on your machine may cause stability issues with Composer.
If you encounter issues, try to change the following:

The zlib extension is not loaded, this can slow down Composer a lot.
If possible, install it or recompile php with --with-zlib

The php.ini used by your command-line PHP is: /etc/php5/php.ini
If you can not modify the ini file, you can also run `php -d option=value` to modify ini values on the fly. You can use -d multiple times.

+ echo

 ---> 5f51addeb2cb
Removing intermediate container 64950711cc4e
Step 4 : COPY docker-entrypoint.sh *.py app/
 ---> 484428cfc763
Removing intermediate container b88c8b94c5fa
Step 5 : WORKDIR app/
 ---> Running in 49c72caddee4
 ---> 8fb78831d831
Removing intermediate container 49c72caddee4
Step 6 : RUN touch crontab.tmp     && if [[ ! -e hosts.list ]]; then         echo "Add localhost into hosts list"; 	    echo "127.0.0.1" > hosts.list;     fi     && mkdir -p /output     && touch /output/web1check.out     && echo '*/2 * * * * python web1check.py>/output/web1check.out' > crontab.tmp     && crontab crontab.tmp     && rm -rf crontab.tmp
 ---> Running in 34db98048ff4
Add localhost into hosts list
 ---> 138fa4818ad2
Removing intermediate container 34db98048ff4
Step 7 : ENTRYPOINT docker-entrypoint.sh
 ---> Running in d61b4105ce69
 ---> 4ca2358472be
Removing intermediate container d61b4105ce69
Step 8 : CMD -m 2 -a web1check.py 127.0.0.1 localhost
 ---> Running in 9ff504e3e45d
 ---> e2b7fe745d62
Removing intermediate container 9ff504e3e45d
Successfully built e2b7fe745d62
```

Image
```
[vagrant@bogon alive-checker]$ docker images tangfeixiong/target-alive-checker
REPOSITORY                          TAG                 IMAGE ID            CREATED             SIZE
tangfeixiong/target-alive-checker   latest              e2b7fe745d62        35 seconds ago      123.6 MB
```
