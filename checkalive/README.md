# Instruction

Wrapper server working like crond

## Table of content
* [Development](#development)
* [Test](#test)

## Development

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

### Docker

Build
```
[vagrant@bogon go-to-docker]$ ./checkalive/build-docker-alpine.sh --build
Sending build context to Docker daemon 10.19 MB
Step 1 : FROM alpine
 ---> 7328f6f8b418
Step 2 : LABEL maintainer 'tangfeixiong <tangfx128@gmail.com>' project "https://github.com/tangfeixiong/go-to-docker" name "tickerjobcm" namespace "stackdocker" annotation '{"stackdocker.io/created-by":"n/a"}' tag "alpine python"
 ---> Running in 0dd87420592e
 ---> cee892f056ef
Removing intermediate container 0dd87420592e
Step 3 : RUN set -x     && apk add --update         bash 		ca-certificates         curl         wget         git         mysql-client 		openssl-dev 		libxml2-dev 		libxslt-dev 		libffi-dev         python         py-pip         py-mysqldb 		python-dev 		build-base     && rm -rf /var/cache/apk/*     && pip install 		requests 	    MySQL-python redis 		MultipartPostHandler 		scp 		pyquery 		multipart_encode 		poster     && echo
 ---> Running in 0d43475556cb
+ apk add --update bash ca-certificates curl wget git mysql-client openssl-dev libxml2-dev libxslt-dev libffi-dev python py-pip py-mysqldb python-dev build-base
fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/community/x86_64/APKINDEX.tar.gz
(1/57) Upgrading musl (1.1.16-r10 -> 1.1.16-r14)
(2/57) Installing ncurses-terminfo-base (6.0_p20170930-r0)
(3/57) Installing ncurses-terminfo (6.0_p20170930-r0)
(4/57) Installing ncurses-libs (6.0_p20170930-r0)
(5/57) Installing readline (6.3.008-r5)
(6/57) Installing bash (4.3.48-r1)
Executing bash-4.3.48-r1.post-install
(7/57) Installing binutils-libs (2.28-r3)
(8/57) Installing binutils (2.28-r3)
(9/57) Installing gmp (6.1.2-r0)
(10/57) Installing isl (0.17.1-r0)
(11/57) Installing libgomp (6.3.0-r4)
(12/57) Installing libatomic (6.3.0-r4)
(13/57) Installing pkgconf (1.3.7-r0)
(14/57) Installing libgcc (6.3.0-r4)
(15/57) Installing mpfr3 (3.1.5-r0)
(16/57) Installing mpc1 (1.0.3-r0)
(17/57) Installing libstdc++ (6.3.0-r4)
(18/57) Installing gcc (6.3.0-r4)
(19/57) Installing musl-dev (1.1.16-r14)
(20/57) Installing libc-dev (0.7.1-r0)
(21/57) Installing g++ (6.3.0-r4)
(22/57) Installing make (4.2.1-r0)
(23/57) Installing fortify-headers (0.8-r0)
(24/57) Installing build-base (0.5-r0)
(25/57) Installing ca-certificates (20161130-r2)
(26/57) Installing libssh2 (1.8.0-r1)
(27/57) Installing libcurl (7.56.1-r0)
(28/57) Installing curl (7.56.1-r0)
(29/57) Installing expat (2.2.0-r1)
(30/57) Installing pcre (8.41-r0)
(31/57) Installing git (2.13.5-r0)
(32/57) Upgrading musl-utils (1.1.16-r10 -> 1.1.16-r14)
(33/57) Installing libffi (3.2.1-r3)
(34/57) Installing libffi-dev (3.2.1-r3)
(35/57) Installing zlib-dev (1.2.11-r0)
(36/57) Installing libxml2 (2.9.4-r4)
(37/57) Installing libxml2-dev (2.9.4-r4)
(38/57) Installing libgpg-error (1.27-r0)
(39/57) Installing libgcrypt (1.7.9-r0)
(40/57) Installing libxslt (1.1.29-r3)
(41/57) Installing libxslt-dev (1.1.29-r3)
(42/57) Installing mariadb-common (10.1.26-r0)
(43/57) Installing mariadb-client (10.1.26-r0)
(44/57) Installing mysql-client (10.1.26-r0)
(45/57) Installing libcrypto1.0 (1.0.2k-r0)
(46/57) Installing libssl1.0 (1.0.2k-r0)
(47/57) Installing openssl-dev (1.0.2k-r0)
(48/57) Installing libbz2 (1.0.6-r5)
(49/57) Installing gdbm (1.12-r0)
(50/57) Installing sqlite-libs (3.20.1-r0)
(51/57) Installing python2 (2.7.13-r1)
(52/57) Installing mariadb-client-libs (10.1.26-r0)
(53/57) Installing py-mysqldb (1.2.5-r0)
(54/57) Installing py-setuptools (33.1.1-r1)
(55/57) Installing py2-pip (9.0.1-r1)
(56/57) Installing python2-dev (2.7.13-r1)
(57/57) Installing wget (1.19.1-r2)
Executing busybox-1.26.2-r5.trigger
Executing ca-certificates-20161130-r2.trigger
OK: 304 MiB in 66 packages
+ rm -rf /var/cache/apk/APKINDEX.24d64ab1.tar.gz /var/cache/apk/APKINDEX.84815163.tar.gz
+ pip install requests MySQL-python redis MultipartPostHandler scp pyquery multipart_encode poster
Collecting requests
  Downloading requests-2.18.4-py2.py3-none-any.whl (88kB)
Requirement already satisfied: MySQL-python in /usr/lib/python2.7/site-packages
Collecting redis
  Downloading redis-2.10.6-py2.py3-none-any.whl (64kB)
Collecting MultipartPostHandler
  Downloading MultipartPostHandler-0.1.0.tar.gz
Collecting scp
  Downloading scp-0.10.2-py2.py3-none-any.whl
Collecting pyquery
  Downloading pyquery-1.3.0-py2.py3-none-any.whl
Collecting multipart_encode
  Downloading multipart-encode-0.1.0.tar.gz
Collecting poster
  Downloading poster-0.8.1.tar.gz
Collecting chardet<3.1.0,>=3.0.2 (from requests)
  Downloading chardet-3.0.4-py2.py3-none-any.whl (133kB)
Collecting certifi>=2017.4.17 (from requests)
  Downloading certifi-2017.11.5-py2.py3-none-any.whl (330kB)
Collecting urllib3<1.23,>=1.21.1 (from requests)
  Downloading urllib3-1.22-py2.py3-none-any.whl (132kB)
Collecting idna<2.7,>=2.5 (from requests)
  Downloading idna-2.6-py2.py3-none-any.whl (56kB)
Collecting paramiko (from scp)
  Downloading paramiko-2.3.1-py2.py3-none-any.whl (182kB)
Collecting lxml>=2.1 (from pyquery)
  Downloading lxml-4.1.1.tar.gz (2.4MB)
Collecting cssselect>0.7.9 (from pyquery)
  Downloading cssselect-1.0.1-py2.py3-none-any.whl
Collecting pyasn1>=0.1.7 (from paramiko->scp)
  Downloading pyasn1-0.3.7-py2.py3-none-any.whl (63kB)
Collecting bcrypt>=3.1.3 (from paramiko->scp)
  Downloading bcrypt-3.1.4.tar.gz (42kB)
Collecting cryptography>=1.5 (from paramiko->scp)
  Downloading cryptography-2.1.3.tar.gz (441kB)
Collecting pynacl>=1.0.1 (from paramiko->scp)
  Downloading PyNaCl-1.2.0.tar.gz (3.3MB)
Collecting cffi>=1.1 (from bcrypt>=3.1.3->paramiko->scp)
  Downloading cffi-1.11.2.tar.gz (435kB)
Collecting six>=1.4.1 (from bcrypt>=3.1.3->paramiko->scp)
  Downloading six-1.11.0-py2.py3-none-any.whl
Collecting asn1crypto>=0.21.0 (from cryptography>=1.5->paramiko->scp)
  Downloading asn1crypto-0.23.0-py2.py3-none-any.whl (99kB)
Collecting enum34 (from cryptography>=1.5->paramiko->scp)
  Downloading enum34-1.1.6-py2-none-any.whl
Collecting ipaddress (from cryptography>=1.5->paramiko->scp)
  Downloading ipaddress-1.0.18-py2-none-any.whl
Collecting pycparser (from cffi>=1.1->bcrypt>=3.1.3->paramiko->scp)
  Downloading pycparser-2.18.tar.gz (245kB)
Installing collected packages: chardet, certifi, urllib3, idna, requests, redis, MultipartPostHandler, pyasn1, pycparser, cffi, six, bcrypt, asn1crypto, enum34, ipaddress, cryptography, pynacl, paramiko, scp, lxml, cssselect, pyquery, multipart-encode, poster
  Running setup.py install for MultipartPostHandler: started
    Running setup.py install for MultipartPostHandler: finished with status 'done'
  Running setup.py install for pycparser: started
    Running setup.py install for pycparser: finished with status 'done'
  Running setup.py install for cffi: started
    Running setup.py install for cffi: finished with status 'done'
  Running setup.py install for bcrypt: started
    Running setup.py install for bcrypt: finished with status 'done'
  Running setup.py install for cryptography: started
    Running setup.py install for cryptography: finished with status 'done'
  Running setup.py install for pynacl: started
    Running setup.py install for pynacl: still running...
    Running setup.py install for pynacl: finished with status 'done'
  Running setup.py install for lxml: started
    Running setup.py install for lxml: still running...
    Running setup.py install for lxml: finished with status 'done'
  Running setup.py install for multipart-encode: started
    Running setup.py install for multipart-encode: finished with status 'done'
  Running setup.py install for poster: started
    Running setup.py install for poster: finished with status 'done'
Successfully installed MultipartPostHandler-0.1.0 asn1crypto-0.23.0 bcrypt-3.1.4 certifi-2017.11.5 cffi-1.11.2 chardet-3.0.4 cryptography-2.1.3 cssselect-1.0.1 enum34-1.1.6 idna-2.6 ipaddress-1.0.18 lxml-4.1.1 multipart-encode-0.1.0 paramiko-2.3.1 poster-0.8.1 pyasn1-0.3.7 pycparser-2.18 pynacl-1.2.0 pyquery-1.3.0 redis-2.10.6 requests-2.18.4 scp-0.10.2 six-1.11.0 urllib3-1.22
+ echo

 ---> 5af7b70ec708
Removing intermediate container 0d43475556cb
Step 4 : COPY bin/checkalive program/ /
 ---> 0d3ce8638f50
Removing intermediate container 5a414f5c7bc0
Step 5 : EXPOSE 10061 10062
 ---> Running in c694aa11576d
 ---> 53760f0f842f
Removing intermediate container c694aa11576d
Step 6 : ENTRYPOINT /checkalive serve
 ---> Running in 2afbb2c03872
 ---> 834b41c52664
Removing intermediate container 2afbb2c03872
Step 7 : CMD --v=2 --logtostderr=true
 ---> Running in 4e8df7db49c6
 ---> e09378469ac9
Removing intermediate container 4e8df7db49c6
Successfully built e09378469ac9
```

Image
```
[vagrant@localhost go-to-docker]$ docker images tangfeixiong/target-cm
REPOSITORY                         TAG                 IMAGE ID            CREATED             SIZE
docker.io/tangfeixiong/target-cm   0.1                 e09378469ac9        5 minutes ago       334.8 MB
```

Inside
```
[vagrant@localhost go-to-docker]$ docker run -ti --rm --entrypoint="/bin/ash" tangfeixiong/target-cm:0.1 -c "ls /"
bin         dev         examples    lib         mnt         root        sbin        sys         usr
checkalive  etc         home        media       proc        run         srv         tmp         var
```

## Test

### Via curl

Refering to run [service](#server) at first.

Start with [runtests_curl.sh](./runtests_curl.sh)
```
[vagrant@localhost go-to-docker]$ ./checkalive/runtests_curl.sh start-awd01-lemon-cms-check
{"name":"awd01_lemon_cms_check.py","command":["python","awd01_lemon_cms_check.py"],"args":["--host=$(ip)","--port=$(port)"],"periodic":3,"duration":10,"dest_configurations":{"team1":{"name":"container1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd01_lemon_cms/awd01_lemon_cms_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd10-nothing
{"name":"awd10_nothing_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd11-maccms
{"error":"Dispatcher exists, delete or update first. name=awd10_nothing_check.py","code":2}[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd11-maccms
{"name":"awd11_maccms_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd11_maccms/awd11_maccms_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd12-phpsqllitecms
{"name":"awd12_phpsqllitecms_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd12_phpsqllitecms/awd12_phpsqllitecms_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd1-lemon-cms
{"name":"awd1_lemon_cms_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd1_lemon_cms/awd1_lemon_cms_check.py"}[vagrant@localhost go-to-docker]
$ checkalive/runtests_curl.sh start-awd1-xmanweb2
{"name":"awd1_xmanweb2_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd1_xmanweb2/awd1_xmanweb2_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd2-daydayweb-check
{"error":"Command file awd1_daydayweb_check.py not found","code":2}[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd2-daydayweb-check
{"name":"awd2_daydayweb_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd2_daydayweb/awd2_daydayweb_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd2-dynpage-check
{"name":"awd2_dynpage_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd2_dynpage/awd2_dynpage_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd3-electronics-check
{"name":"awd3_electronics_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd3_electronics/awd3_electronics_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd3-shadow-check
{"name":"awd3_shadow_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd3_shadow/awd3_shadow_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd4-chinaz-check
{"name":"awd4_chinaz_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd4_chinaz/awd4_chinaz_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd4-tomcat-check
{"name":"awd4_tomcat_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd4_tomcat/awd4_tomcat_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd5-babyblog-check
{"name":"awd5_babyblog_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd5_babyblog/awd5_babyblog_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd5-gracer-check
{"name":"awd5_gracer_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd5_gracer/awd5_gracer_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd6-cms-check
{"name":"awd6_cms_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd6_cms/awd6_cms_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd7-upload-check
{"name":"awd7_upload_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd7_upload/awd7_upload_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd8-blog-check
{"name":"awd8_blog_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd8_blog/awd8_blog_check.py"}
[vagrant@localhost go-to-docker]$ checkalive/runtests_curl.sh start-awd9-money-check
{"name":"awd9_money_check.py","args":["--host=localhost","--port=80"],"periodic":3,"duration":10,"dest_configurations":{"env1":{"name":"team1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"}}},"destination_path":"examples/python/check-alive/awd9_money/awd9_money_check.py"}
```

Reap
```
fanhonglingdeMacBook-Pro:checkalive fanhongling$ ./runtests_curl.sh reap
{"name":"awd10_nothing_check.py","command":["python","awd10_nothing_check.py"],"args":["--host=localhost","--port=80"],"work_dir":"awd10_nothing","periodic":6,"duration":12,"dest_configurations":{"team1":{"name":"container1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"},"state_code":3,"state_message":"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\u003curllib3.connection.HTTPConnection object at 0x7f85c0acaa90\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\u003curllib3.connection.HTTPConnection object at 0x7fc19a4c7a90\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\u003curllib3.connection.HTTPConnection object at 0x7f75b80a9a90\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\n\n","timestamp":"2017-10-16T07:54:34Z"}},"destination_path":"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py"}
```

Update
```
fanhonglingdeMacBook-Pro:checkalive fanhongling$ ./runtests_curl.sh update
{"name":"awd10_nothing_check.py","command":["python","awd10_nothing_check.py"],"args":["--host=localhost","--port=80"],"work_dir":"awd10_nothing","periodic":6,"duration":12,"dest_configurations":{"team1":{"name":"container1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"},"state_code":1,"state_message":"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\u003curllib3.connection.HTTPConnection object at 0x7f85c0acaa90\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\n\n","timestamp":"2017-10-16T07:54:21Z"}},"destination_path":"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py"}
```

Delete
```
fanhonglingdeMacBook-Pro:checkalive fanhongling$ ./runtests_curl.sh delete
{"name":"awd10_nothing_check.py","command":["python","awd10_nothing_check.py"],"args":["--host=localhost","--port=80"],"work_dir":"awd10_nothing","periodic":6,"duration":12,"dest_configurations":{"team1":{"name":"container1","args":["--host=localhost","--port=80"],"tpl":{"ip":"localhost","port":"80"},"state_code":3,"state_message":"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\u003curllib3.connection.HTTPConnection object at 0x7f85c0acaa90\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\u003curllib3.connection.HTTPConnection object at 0x7fc19a4c7a90\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\u003curllib3.connection.HTTPConnection object at 0x7f75b80a9a90\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\n\n","timestamp":"2017-10-16T07:54:34Z"}},"destination_path":"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py"}
```

### Server

Note: server must be tested at repository home dir
```
[vagrant@localhost go-to-docker]$ ./checkalive/bin/checkalive serve --v=2 --logtostderr
Start gRPC Gateway into host :10061
http on host: [::]:10062
Start gRPC on host [::]:10061
I1105 15:46:09.113042   16403 server.go:267] go to create: "name:\"awd01_lemon_cms_check.py\" command:\"python\" command:\"awd01_lemon_cms_check.py\" args:\"--host=$(ip)\" args:\"--port=$(port)\" work_dir:\"awd10_lemon_cms\" periodic:3 duration:10 dest_configurations:<key:\"team1\" value:<name:\"container1\" tpl:<key:\"ip\" value:\"localhost\" > tpl:<key:\"port\" value:\"80\" > > > "
filepath.Walk() returned <nil>
I1105 15:46:09.184147   16403 checkctl.go:208] Command file not found: awd01_lemon_cms_check.py <nil>
proto: no coders for int
I1105 15:46:55.168002   16403 server.go:267] go to create: "name:\"awd01_lemon_cms_check.py\" command:\"python\" command:\"awd01_lemon_cms_check.py\" args:\"--host=$(ip)\" args:\"--port=$(port)\" work_dir:\"awd01_lemon_cms\" periodic:3 duration:10 dest_configurations:<key:\"team1\" value:<name:\"container1\" tpl:<key:\"ip\" value:\"localhost\" > tpl:<key:\"port\" value:\"80\" > > > "
Visited: examples/python/check-alive/awd01_lemon_cms/awd01_lemon_cms_check.py
filepath.Walk() returned Stop recursive searching
command: awd01_lemon_cms_check.py
Dest config of team1: name:"container1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > name:"awd01_lemon_cms_check.py" command:"python" command:"awd01_lemon_cms_check.py" args:"--host=$(ip)" args:"--port=$(port)" work_dir:"awd01_lemon_cms" periodic:3 duration:10 dest_configurations:<key:"team1" value:<name:"container1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > > > destination_path:"examples/python/check-alive/awd01_lemon_cms/awd01_lemon_cms_check.py" 
Tick at 2017-11-05 15:46:58.181767447 +0000 UTC m=+247.612239078 -> key: team1
HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /index.php (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7fb60d94cf90>: Failed to establish a new connection: [Errno 111] Connection refused',))
(False, 'check index page exception')

name:"awd01_lemon_cms_check.py" command:"python" command:"awd01_lemon_cms_check.py" args:"--host=$(ip)" args:"--port=$(port)" work_dir:"awd01_lemon_cms" periodic:3 duration:10 dest_configurations:<key:"team1" value:<name:"container1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > state_code:1 state_message:"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /index.php (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7fb60d94cf90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check index page exception')\n\n" timestamp:"2017-11-05T15:46:58Z" > > destination_path:"examples/python/check-alive/awd01_lemon_cms/awd01_lemon_cms_check.py" 
Tick at 2017-11-05 15:47:01.181766797 +0000 UTC m=+250.612237503 -> key: team1
HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /index.php (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f6e8b975f90>: Failed to establish a new connection: [Errno 111] Connection refused',))
(False, 'check index page exception')

name:"awd01_lemon_cms_check.py" command:"python" command:"awd01_lemon_cms_check.py" args:"--host=$(ip)" args:"--port=$(port)" work_dir:"awd01_lemon_cms" periodic:3 duration:10 dest_configurations:<key:"team1" value:<name:"container1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > state_code:2 state_message:"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /index.php (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7fb60d94cf90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check index page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /index.php (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f6e8b975f90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check index page exception')\n\n" timestamp:"2017-11-05T15:47:01Z" > > destination_path:"examples/python/check-alive/awd01_lemon_cms/awd01_lemon_cms_check.py" 
Tick at 2017-11-05 15:47:04.181135933 +0000 UTC m=+253.611606490 -> key: team1
HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /index.php (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f3b70257f90>: Failed to establish a new connection: [Errno 111] Connection refused',))
(False, 'check index page exception')

name:"awd01_lemon_cms_check.py" command:"python" command:"awd01_lemon_cms_check.py" args:"--host=$(ip)" args:"--port=$(port)" work_dir:"awd01_lemon_cms" periodic:3 duration:10 dest_configurations:<key:"team1" value:<name:"container1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > state_code:3 state_message:"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /index.php (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7fb60d94cf90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check index page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /index.php (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f6e8b975f90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check index page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /index.php (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f3b70257f90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check index page exception')\n\n" timestamp:"2017-11-05T15:47:04Z" > > destination_path:"examples/python/check-alive/awd01_lemon_cms/awd01_lemon_cms_check.py" 
Stopped
```

### Via docker

Refer to `docker-compose.yml` or `Dockerfile` for details
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

or docker-compose logging
```
checker-cm_1  | I1016 07:54:18.576679       1 server.go:267] go to create: "name:\"awd10_nothing_check.py\" command:\"python\" command:\"awd10_nothing_check.py\" args:\"--host=$(ip)\" args:\"--port=$(port)\" work_dir:\"awd10_nothing\" periodic:3 duration:10 dest_configurations:<key:\"team1\" value:<name:\"container1\" tpl:<key:\"ip\" value:\"localhost\" > tpl:<key:\"port\" value:\"80\" > > > "
checker-cm_1  | Visited: examples/python/check-alive/awd10_nothing/awd10_nothing_check.py
checker-cm_1  | filepath.Walk() returned Stop recursive searching
checker-cm_1  | command: awd10_nothing_check.py
checker-cm_1  | Dest tpl: team1
checker-cm_1  | name:"awd10_nothing_check.py" command:"python" command:"awd10_nothing_check.py" args:"--host=localhost" args:"--port=80" work_dir:"awd10_nothing" periodic:3 duration:10 dest_configurations:<key:"team1" value:<name:"container1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > > > destination_path:"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py" 
checker-cm_1  | Tick at 2017-10-16 07:54:21.580857208 +0000 UTC -> key: team1
checker-cm_1  | HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f85c0acaa90>: Failed to establish a new connection: [Errno 111] Connection refused',))
checker-cm_1  | (False, 'check login page exception')
checker-cm_1  | 
checker-cm_1  | I1016 07:54:21.668922       1 server.go:566] write cm into cache
checker-cm_1  | I1016 07:54:21.910763       1 server.go:594] Set CM checkalive.awd10_nothing_check.py: OK
checker-cm_1  | I1016 07:54:21.911660       1 server.go:646] publish check...
checker-cm_1  | I1016 07:54:22.253355       1 server.go:667] Published subject checkalive.awd10_nothing_check.py: 0
checker-cm_1  | I1016 07:54:22.631953       1 server.go:478] go to update: "name:\"awd10_nothing_check.py\" periodic:6 duration:12 "
checker-cm_1  | Ticker stopped
checker-cm_1  | name:"awd10_nothing_check.py" command:"python" command:"awd10_nothing_check.py" args:"--host=localhost" args:"--port=80" work_dir:"awd10_nothing" periodic:6 duration:12 dest_configurations:<key:"team1" value:<name:"container1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > state_code:1 state_message:"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f85c0acaa90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\n\n" timestamp:"2017-10-16T07:54:21Z" > > destination_path:"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py" 
checker-cm_1  | Tick at 2017-10-16 07:54:28.635395148 +0000 UTC -> key: team1
checker-cm_1  | HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7fc19a4c7a90>: Failed to establish a new connection: [Errno 111] Connection refused',))
checker-cm_1  | (False, 'check login page exception')
checker-cm_1  | 
checker-cm_1  | I1016 07:54:28.741270       1 server.go:566] write cm into cache
checker-cm_1  | I1016 07:54:28.966716       1 server.go:594] Set CM checkalive.awd10_nothing_check.py: OK
checker-cm_1  | I1016 07:54:28.966770       1 server.go:646] publish check...
checker-cm_1  | I1016 07:54:29.196228       1 server.go:667] Published subject checkalive.awd10_nothing_check.py: 0
checker-cm_1  | name:"awd10_nothing_check.py" command:"python" command:"awd10_nothing_check.py" args:"--host=localhost" args:"--port=80" work_dir:"awd10_nothing" periodic:6 duration:12 dest_configurations:<key:"team1" value:<name:"container1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > state_code:2 state_message:"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f85c0acaa90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7fc19a4c7a90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\n\n" timestamp:"2017-10-16T07:54:28Z" > > destination_path:"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py" 
checker-cm_1  | Tick at 2017-10-16 07:54:34.635623057 +0000 UTC -> key: team1
checker-cm_1  | HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f75b80a9a90>: Failed to establish a new connection: [Errno 111] Connection refused',))
checker-cm_1  | (False, 'check login page exception')
checker-cm_1  | 
checker-cm_1  | I1016 07:54:34.730975       1 server.go:566] write cm into cache
checker-cm_1  | I1016 07:54:34.958241       1 server.go:594] Set CM checkalive.awd10_nothing_check.py: OK
checker-cm_1  | I1016 07:54:34.958468       1 server.go:646] publish check...
checker-cm_1  | I1016 07:54:35.184389       1 server.go:667] Published subject checkalive.awd10_nothing_check.py: 0
checker-cm_1  | name:"awd10_nothing_check.py" command:"python" command:"awd10_nothing_check.py" args:"--host=localhost" args:"--port=80" work_dir:"awd10_nothing" periodic:6 duration:12 dest_configurations:<key:"team1" value:<name:"container1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > state_code:3 state_message:"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f85c0acaa90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7fc19a4c7a90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f75b80a9a90>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\n\n" timestamp:"2017-10-16T07:54:34Z" > > destination_path:"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py" 
checker-cm_1  | Stopped
checker-cm_1  | I1016 07:58:35.547362       1 server.go:540] go to reap: "name:\"awd10_nothing_check.py\" "
checker-cm_1  | Ticker stopped
checker-cm_1  | I1016 07:59:09.681736       1 server.go:552] go to delete: "name:\"awd10_nothing_check.py\" "
```

### Redis pub/sub

Container
```
[vagrant@bogon checkalive]$ docker inspect -f {{.NetworkSettings.IPAddress}} redis
172.17.0.9
```

Or export environment to execute
```
[vagrant@bogon go-to-docker]$ export DATABUS_REDIS_HOST=172.17.0.9:6379
[vagrant@bogon go-to-docker]$ export DATABUS_REDIS_DB=15
```

Or execute server with process-only environment  (Note: By default, server must be tested at repository home dir)
```
[vagrant@bogon go-to-docker]$ DATABUS_REDIS_HOST=172.17.0.9:6379 DATABUS_REDIS_DB=15 ./checkalive/bin/checkalive serve --v=2 --logtostderr
Start gRPC Gateway into host :10061
Start gRPC on host [::]:10061
http on host: [::]:10062
I1002 09:18:53.465940   22062 server.go:267] go to create check: "name:\"awd10_nothing_check.py\" command:\"python\" command:\"awd10_nothing_check.py\" args:\"--host=$(ip)\" args:\"--port=$(port)\" work_dir:\"awd10_nothing\" periodic:3 duration:10 dest_configurations:<key:\"env1\" value:<name:\"team1\" tpl:<key:\"ip\" value:\"localhost\" > tpl:<key:\"port\" value:\"80\" > > > "
Visited: examples/python/check-alive/awd10_nothing/awd10_nothing_check.py
filepath.Walk() returned Stop recursive searching
command: awd10_nothing_check.py
Dest tpl: env1
name:"awd10_nothing_check.py" command:"python" command:"awd10_nothing_check.py" args:"--host=localhost" args:"--port=80" work_dir:"awd10_nothing" periodic:3 duration:10 dest_configurations:<key:"env1" value:<name:"team1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > > > destination_path:"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py" 
Tick at 2017-10-02 09:18:56.480039322 +0000 UTC -> key: env1
HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f343dcba490>: Failed to establish a new connection: [Errno 111] Connection refused',))
(False, 'check login page exception')

I1002 09:18:56.939705   22062 server.go:556] write cm into cache
I1002 09:18:56.943085   22062 server.go:584] Set CM checkalive.awd10_nothing_check.py: OK
I1002 09:18:56.943873   22062 server.go:636] publish check...
I1002 09:18:56.944706   22062 server.go:657] Published subject checkalive.awd10_nothing_check.py: 0
name:"awd10_nothing_check.py" command:"python" command:"awd10_nothing_check.py" args:"--host=localhost" args:"--port=80" work_dir:"awd10_nothing" periodic:3 duration:10 dest_configurations:<key:"env1" value:<name:"team1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > state_code:1 state_message:"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f343dcba490>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\n\n" timestamp:"2017-10-02T09:18:56Z" > > destination_path:"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py" 
Tick at 2017-10-02 09:18:59.479471649 +0000 UTC -> key: env1
HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f222f1e5490>: Failed to establish a new connection: [Errno 111] Connection refused',))
(False, 'check login page exception')

I1002 09:18:59.942978   22062 server.go:556] write cm into cache
I1002 09:18:59.944269   22062 server.go:584] Set CM checkalive.awd10_nothing_check.py: OK
I1002 09:18:59.944922   22062 server.go:636] publish check...
I1002 09:18:59.945772   22062 server.go:657] Published subject checkalive.awd10_nothing_check.py: 0
name:"awd10_nothing_check.py" command:"python" command:"awd10_nothing_check.py" args:"--host=localhost" args:"--port=80" work_dir:"awd10_nothing" periodic:3 duration:10 dest_configurations:<key:"env1" value:<name:"team1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > state_code:2 state_message:"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f343dcba490>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f222f1e5490>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\n\n" timestamp:"2017-10-02T09:18:59Z" > > destination_path:"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py" 
Tick at 2017-10-02 09:19:02.479418088 +0000 UTC -> key: env1
HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f7160112490>: Failed to establish a new connection: [Errno 111] Connection refused',))
(False, 'check login page exception')

I1002 09:19:03.203510   22062 server.go:556] write cm into cache
I1002 09:19:03.204326   22062 server.go:584] Set CM checkalive.awd10_nothing_check.py: OK
I1002 09:19:03.204915   22062 server.go:636] publish check...
I1002 09:19:03.205872   22062 server.go:657] Published subject checkalive.awd10_nothing_check.py: 0
name:"awd10_nothing_check.py" command:"python" command:"awd10_nothing_check.py" args:"--host=localhost" args:"--port=80" work_dir:"awd10_nothing" periodic:3 duration:10 dest_configurations:<key:"env1" value:<name:"team1" args:"--host=localhost" args:"--port=80" tpl:<key:"ip" value:"localhost" > tpl:<key:"port" value:"80" > state_code:3 state_message:"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f343dcba490>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f222f1e5490>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('<requests.packages.urllib3.connection.HTTPConnection object at 0x7f7160112490>: Failed to establish a new connection: [Errno 111] Connection refused',))\n(False, 'check login page exception')\n\n" timestamp:"2017-10-02T09:19:02Z" > > destination_path:"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py" 
Stopped
^C
```

Or
```
[vagrant@bogon go-to-docker]$ docker run --rm -ti --name check-alive -p 10062:10062 -e DATABUS_REDIS_HOST=172.17.0.9:6379 -e DATABUS_REDIS_DB=15 docker.io/tangfeixiong/target-cm:0.1
```

Redis cli
```
[vagrant@localhost go-to-docker]$ docker exec -ti redis bash
root@fe0ae27989c4:/# redis-cli 
127.0.0.1:6379> select 15
OK
127.0.0.1:6379[15]> keys *
1) "checkalive.awd10_nothing_check.py"
127.0.0.1:6379[15]> get "checkalive.awd10_nothing_check.py"
"{\"name\":\"awd10_nothing_check.py\",\"command\":[\"python\",\"awd10_nothing_check.py\"],\"args\":[\"--host=localhost\",\"--port=80\"],\"work_dir\":\"awd10_nothing\",\"periodic\":3,\"duration\":10,\"dest_configurations\":{\"env1\":{\"name\":\"team1\",\"args\":[\"--host=localhost\",\"--port=80\"],\"tpl\":{\"ip\":\"localhost\",\"port\":\"80\"},\"state_code\":3,\"state_message\":\"HTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\\u003crequests.packages.urllib3.connection.HTTPConnection object at 0x7f343dcba490\\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\\n(False, 'check login page exception')\\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\\u003crequests.packages.urllib3.connection.HTTPConnection object at 0x7f222f1e5490\\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\\n(False, 'check login page exception')\\nHTTPConnectionPool(host='localhost', port=80): Max retries exceeded with url: /login.html (Caused by NewConnectionError('\\u003crequests.packages.urllib3.connection.HTTPConnection object at 0x7f7160112490\\u003e: Failed to establish a new connection: [Errno 111] Connection refused',))\\n(False, 'check login page exception')\\n\\n\",\"timestamp\":\"2017-10-02T09:19:02Z\"}},\"destination_path\":\"examples/python/check-alive/awd10_nothing/awd10_nothing_check.py\"}"
127.0.0.1:6379[15]> quit
root@fe0ae27989c4:/# exit
exit
```