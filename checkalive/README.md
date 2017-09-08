# Wrapper server working like crond

## Build

Protobuf
```
[vagrant@bogon go-to-docker]$ ./checkalive/build-Protobuf.sh
```

Bin
```
[vagrant@bogon go-to-docker]$ ./checkalive/build-go.sh
```

Alternative bin
```
[vagrant@bogon go-to-docker]$ ./checkalive/build-go.sh --cgo
```

Docker
```
[vagrant@bogon go-to-docker]$ ./checkalive/build-docker-alpine.sh 
Sending build context to Docker daemon 9.309 MB
Step 1 : FROM alpine
 ---> 7328f6f8b418
Step 2 : LABEL maintainer 'tangfeixiong <tangfx128@gmail.com>' project "https://github.com/tangfeixiong/go-to-docker" name "tickerjobcm" namespace "stackdocker" annotation '{"stackdocker.io/created-by":"n/a"}' tag "alpine php python"
 ---> Using cache
 ---> 84039785f7c7
Step 3 : RUN set -x     && apk add --update         bash         curl         wget         git         mysql-client         php5         php5-cli         php5-common         php5-gd         php5-phar         php5-curl         php5-mysql         php5-openssl         php5-json         php5-dom         python         py-pip         py-mysqldb     && rm -rf /var/cache/apk/*     && pip install MySQL-python redis     && ln -sf /usr/bin/php5 /usr/bin/php     && wget https://raw.githubusercontent.com/composer/getcomposer.org/master/web/installer -O - -q | php -- --quiet     && echo
 ---> Running in 288c510e8e06
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
+ wget https://raw.githubusercontent.com/composer/getcomposer.org/master/web/installer -O - -q
+ php -- --quiet
Some settings on your machine may cause stability issues with Composer.
If you encounter issues, try to change the following:

The zlib extension is not loaded, this can slow down Composer a lot.
If possible, install it or recompile php with --with-zlib

The php.ini used by your command-line PHP is: /etc/php5/php.ini
If you can not modify the ini file, you can also run `php -d option=value` to modify ini values on the fly. You can use -d multiple times.

+ echo

 ---> 7fa7e9ddb00b
Removing intermediate container 288c510e8e06
Step 4 : COPY bin/ program/ /
 ---> 0f56ed79331a
Removing intermediate container 732545d123db
Step 5 : VOLUME /examples
 ---> Running in 76aa72449097
 ---> 6b273bae4db4
Removing intermediate container 76aa72449097
Step 6 : EXPOSE 10061 10062
 ---> Running in 7059f79de59f
 ---> 7f620583301d
Removing intermediate container 7059f79de59f
Step 7 : ENTRYPOINT /target-cm serve
 ---> Running in ec06533fff32
 ---> e333179d9852
Removing intermediate container ec06533fff32
Step 8 : CMD --v=2 --logtostderr=true
 ---> Running in 8bec407b9e90
 ---> 15826af34826
Removing intermediate container 8bec407b9e90
Successfully built 15826af34826
```

Image
```
[vagrant@bogon go-to-docker]$ docker images tangfeixiong/target-cm
REPOSITORY                                             TAG                                     IMAGE ID            CREATED              SIZE
docker.io/tangfeixiong/target-cm                       0.1-1709080032-git_0c88f84              15826af34826        About a minute ago   132.9 MB
```

## Test

Create
```
{"name":"web1check.py","command":["python","web1check.py"],"conf":{"hosts.list":"bG9jYWxob3N0Cg=="},"periodic":3,"destination_path":"examples/python/checkalive/web1check.py"}
```

Reap
```
fanhonglingdeMacBook-Pro:checkalive fanhongling$ ./runtests_curl.sh reap
{"name":"web1check.py","command":["python","web1check.py"],"conf":{"hosts.list":"bG9jYWxob3N0Cg=="},"periodic":3,"state_message":"---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n\n---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n","timestamp":"2017-09-07T23:13:56Z","destination_path":"examples/python/checkalive/web1check.py"}
```

Update
```
fanhonglingdeMacBook-Pro:checkalive fanhongling$ ./runtests_curl.sh update
{"name":"web1check.py","command":["python","web1check.py"],"conf":{"hosts.list":"MTI3LjAuMC4xCg=="},"periodic":5,"state_message":"---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n\n---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n","timestamp":"2017-09-07T23:13:59Z","destination_path":"examples/python/checkalive/web1check.py"}
```

Delete
```
fanhonglingdeMacBook-Pro:checkalive fanhongling$ ./runtests_curl.sh delete
{"name":"web1check.py","command":["python","web1check.py"],"conf":{"hosts.list":"MTI3LjAuMC4xCg=="},"periodic":5,"state_message":"---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n\n---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n","timestamp":"2017-09-07T23:13:59Z","destination_path":"examples/python/checkalive/web1check.py"}
```


Server output
```
[vagrant@bogon go-to-docker]$ ./checkalive/bin/target-cm serve --v=2 --logtostderr
Start gRPC Gateway into host :10061
http on host: [::]:10062
Start gRPC on host [::]:10061
go to create check
check path
destination path
Visited: examples/python/checkalive/web1check.py
filepath.Walk() returned Stop recursive searching
config file
Tick at 2017-09-07 23:13:53.42750264 +0000 UTC
---------------------------------------------------------------
checking host: localhost
global name 'headers' is not defined
Host: localhost seems down

Tick at 2017-09-07 23:13:56.427504027 +0000 UTC
---------------------------------------------------------------
checking host: localhost
global name 'headers' is not defined
Host: localhost seems down

go to reap check
Tick at 2017-09-07 23:13:59.430029891 +0000 UTC
---------------------------------------------------------------
checking host: localhost
global name 'headers' is not defined
Host: localhost seems down

go to update check
workdir
config file
Ticker stopped
go to delete check
Ticker stopped


```

### Docker run

For example
```
[vagrant@bogon go-to-docker]$ docker run --rm -ti --name target-cm -p 10062:10062 tangfeixiong/target-cm:0.1-1709080032-git_0c88f84
Start gRPC Gateway into host :10061
Start gRPC on host [::]:10061
http on host: [::]:10062
I0908 00:41:42.491245       1 daemon.go:250] go to create check: "name:\"web1check.py\" command:\"python\" command:\"web1check.py\" conf:<key:\"hosts.list\" value:\"bG9jYWxob3N0Cg==\" > work_dir:\"web1check\" periodic:3 "
I0908 00:41:42.491405       1 daemon.go:275] path: web1check.py
filepath.Walk() returned Stop recursive searching
config file
Tick at 2017-09-08 00:41:45.587638257 +0000 UTC
---------------------------------------------------------------
checking host: localhost
global name 'headers' is not defined
Host: localhost seems down

Tick at 2017-09-08 00:41:48.507896347 +0000 UTC
---------------------------------------------------------------
checking host: localhost
global name 'headers' is not defined
Host: localhost seems down

Tick at 2017-09-08 00:41:51.493957559 +0000 UTC
---------------------------------------------------------------
checking host: localhost
global name 'headers' is not defined
Host: localhost seems down

Tick at 2017-09-08 00:41:54.508434942 +0000 UTC
---------------------------------------------------------------
checking host: localhost
global name 'headers' is not defined
Host: localhost seems down

go to update check
workdir
config file
Ticker stopped
Tick at 2017-09-08 00:41:59.667602647 +0000 UTC
---------------------------------------------------------------
checking host: 127.0.0.1
global name 'headers' is not defined
Host: 127.0.0.1 seems down

Tick at 2017-09-08 00:42:04.6677189 +0000 UTC
---------------------------------------------------------------
checking host: 127.0.0.1
global name 'headers' is not defined
Host: 127.0.0.1 seems down

go to delete check
Ticker stopped
^C
```