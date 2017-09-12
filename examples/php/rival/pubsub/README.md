

### Pub/Sub

Run Redis

Run Server
```
[vagrant@bogon go-to-docker]$ docker run -ti --rm --net ec9fc4fb4445 -p 10062:10062 -e DATABUS_REDIS_HOST=172.18.0.3:6379 -e DATABUS_REDIS_DB=15 docker.io/tangfeixiong/target-cm:0.1
Start gRPC Gateway into host :10061
Start gRPC on host [::]:10061
http on host: [::]:10062
```
Note: In this demo, the `net` argument is used because of Redis working at it
```
[vagrant@bogon go-to-docker]$ docker network inspect -f {{.Containers.fe0ae27989c4006c821e336e52f14817872b4483495d661dcf3f3a275475ab02.IPv4Address}} ec9fc4fb4445
172.18.0.3/16
```

Or cli
```
[vagrant@bogon go-to-docker]$ ./checkalive/bin/target-cm serve --logtostderr
Start gRPC Gateway into host :10061
Start gRPC on host [::]:10061
http on host: [::]:10062
```

Start publisher
```
fanhonglingdeMacBook-Pro:go-to-authnz fanhongling$ ~/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/checkalive/runtests_curl.sh create
{"name":"web1check.py","command":["python","web1check.py"],"conf":{"hosts.list":"bG9jYWxob3N0Cg=="},"work_dir":"web1check","periodic":3,"destination_path":"examples/python/checkalive/web1check/web1check.py"}
```

Server publish
```
I0909 05:25:13.507876       1 daemon.go:270] go to create check: "name:\"web1check.py\" command:\"python\" command:\"web1check.py\" conf:<key:\"hosts.list\" value:\"bG9jYWxob3N0Cg==\" > work_dir:\"web1check\" periodic:3 "
I0909 05:25:13.508021       1 daemon.go:295] path: web1check.py
Visited: examples/python/checkalive/web1check/web1check.py
filepath.Walk() returned Stop recursive searching
config file
Tick at 2017-09-09 05:25:16.511477717 +0000 UTC
---------------------------------------------------------------
checking host: localhost
global name 'headers' is not defined
Host: localhost seems down

I0909 05:25:16.693792       1 daemon.go:483] write cm into cache
I0909 05:25:16.695607       1 daemon.go:515] Set CM .web1check.py: OK
I0909 05:25:16.695670       1 daemon.go:567] public check...
I0909 05:25:16.696964       1 daemon.go:592] Published subject checkalive.web1check.py: 1
Tick at 2017-09-09 05:25:19.510938039 +0000 UTC
---------------------------------------------------------------
checking host: localhost
global name 'headers' is not defined
Host: localhost seems down

I0909 05:25:19.550553       1 daemon.go:483] write cm into cache
I0909 05:25:19.552402       1 daemon.go:515] Set CM .web1check.py: OK
I0909 05:25:19.552892       1 daemon.go:567] public check...
I0909 05:25:19.554019       1 daemon.go:592] Published subject checkalive.web1check.py: 1
Tick at 2017-09-09 05:25:22.511158553 +0000 UTC
```

Run subscriber
```
[vagrant@localhost go-to-docker]$ php examples/php/rival/pubsubredis_consumer.php 
Subscribed to control_channel
Subscribed to notifications
Received the following message from checkalive.web1check.py:
  {"name":"web1check.py","command":["python","web1check.py"],"conf":{"hosts.list":"bG9jYWxob3N0Cg=="},"work_dir":"web1check","periodic":3,"state_message":"---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n","timestamp":"2017-09-09T05:25:16Z","destination_path":"examples/python/checkalive/web1check/web1check.py"}

Received the following message from checkalive.web1check.py:
  {"name":"web1check.py","command":["python","web1check.py"],"conf":{"hosts.list":"bG9jYWxob3N0Cg=="},"work_dir":"web1check","periodic":3,"state_message":"---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n\n---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n","timestamp":"2017-09-09T05:25:19Z","destination_path":"examples/python/checkalive/web1check/web1check.py"}

Received the following message from checkalive.web1check.py:
  {"name":"web1check.py","command":["python","web1check.py"],"conf":{"hosts.list":"bG9jYWxob3N0Cg=="},"work_dir":"web1check","periodic":3,"state_message":"---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n\n---------------------------------------------------------------\nchecking host: localhost\nglobal name 'headers' is not defined\nHost: localhost seems down\n","timestamp":"2017-09-09T05:25:22Z","destination_path":"examples/python/checkalive/web1check/web1check.py"}

```