# go-to-docker

A Micro-Service working around Docker (http://docs.docker.com/) Remote API (version 1.26+, docker 1.11+) & Registry API (v2)

Feature: based on gRPC and gRPC-Gateway (insecure mode)

## Install

On Linux, docker server installed (minimal api version is 1.23), for example
```
[vagrant@kubedev-172-17-4-59 go-to-docker]$ docker version
Client:
 Version:         1.13.1
 API version:     1.26
 Package version: <unknown>
 Go version:      go1.8.4
 Git commit:      584d391/1.13.1
 Built:           Thu Nov 23 21:40:58 2017
 OS/Arch:         linux/amd64

Server:
 Version:         1.13.1
 API version:     1.26 (minimum version 1.12)
 Package version: <unknown>
 Go version:      go1.8.4
 Git commit:      584d391/1.13.1
 Built:           Thu Nov 23 21:40:58 2017
 OS/Arch:         linux/amd64
 Experimental:    false
```

Get docker image
```
$ docker pull docker.io/tangfeixiong/go-to-docker:0.2
```

Run (note: Docker container name must be unique, that's saying if already install before, destroy it firstly like `docker stop go-to-docket` then `docker rm go-to-docker`)
```
[vagrant@kubedev-172-17-4-59 go-to-docker]$ docker run -d --name=go-to-docker -p 10053:10053 -v /var/run/docker.sock:/var/run/docker.sock:ro --privileged=true docker.io/tangfeixiong/go-to-docker:0.2
74223a63ddd98d71a9213e772a653f712b83dde3363fae4a6eabbc75d2833075
```

Log console
```
[vagrant@kubedev-172-17-4-59 go-to-docker]$ docker logs go-to-docker
I0912 21:26:13.970335       1 client.go:93] Start docker client with request timeout=1m59s
```

### Test

To create docker network (bridge)
```
[vagrant@kubedev-172-17-4-59 go-to-docker]$ ./runtests_curl.sh --addr=172.17.4.59:10053 docker-network-create
{"name":"stackdocker-brj5lp62cw7","network_create":{"check_duplicate":false,"driver":"","scope":"","enable_ipv6":false,"ipam":{"driver":"","options":{},"config":[]},"internal":false,"attachable":false,"ingress":false,"config_only":false,"config_from":{"network":""},"options":{},"labels":{}},"network_create_response":{"id":"cd1526477a81d27a2b3751903dd268cc92c1bad453f495ac974407ba61f0eb04","warning":""},"state_code":0,"state_message":"created"}
```

Verified via `docker`
```
[vagrant@kubedev-172-17-4-59 go-to-docker]$ docker network ls
NETWORK ID          NAME                                  DRIVER              SCOPE
ff440256465f        bridge                                bridge              local
05f9661b589a        host                                  host                local
dce70daba980        none                                  null                local
cd1526477a81        stackdocker-brj5lp62cw7               bridge              local
```

and
```
[vagrant@kubedev-172-17-4-59 go-to-docker]$ docker logs go-to-docker
I0912 21:26:13.970335       1 client.go:93] Start docker client with request timeout=1m59s
W0912 21:33:31.249346       1 helper.go:482] Docker network name not specified
W0912 21:33:31.249678       1 helper.go:497] Networking IPAM arguments not specified
I0912 21:33:46.401586       1 network.go:38] docker network cd1526477a81d27a2b3751903dd268cc92c1bad453f495ac974407ba61f0eb04 created
```

### OpanAPI

swagger 2.0
[屏幕快照 2018-09-12 下午2.46.33.png](./docs/屏幕快照%202018-09-12%20下午2.46.33.png)


### More

For example, process status
```
[vagrant@kubedev-172-17-4-59 go-to-docker]$ sudo ps -ef | grep go-to-docker
vagrant  26238  3803  0 21:45 pts/0    00:00:00 grep --color=auto go-to-docker
```

and such as find listening port
```
[vagrant@kubedev-172-17-4-59 go-to-docker]$ sudo netstat -tpnl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 127.0.0.1:10248         0.0.0.0:*               LISTEN      1526/kubelet        
tcp        0      0 127.0.0.1:10249         0.0.0.0:*               LISTEN      3344/kube-proxy     
tcp        0      0 127.0.0.1:10251         0.0.0.0:*               LISTEN      9720/kube-scheduler 
tcp        0      0 127.0.0.1:2379          0.0.0.0:*               LISTEN      2927/etcd           
tcp        0      0 127.0.0.1:10252         0.0.0.0:*               LISTEN      9774/kube-controlle 
tcp        0      0 127.0.0.1:2380          0.0.0.0:*               LISTEN      2927/etcd           
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      472/sshd            
tcp6       0      0 :::10080                :::*                    LISTEN      1059/docker-proxy-c 
tcp6       0      0 :::10053                :::*                    LISTEN      22531/docker-proxy- 
tcp6       0      0 :::10022                :::*                    LISTEN      1070/docker-proxy-c 
tcp6       0      0 :::30538                :::*                    LISTEN      3344/kube-proxy     
tcp6       0      0 :::10250                :::*                    LISTEN      1526/kubelet        
tcp6       0      0 :::6443                 :::*                    LISTEN      9747/kube-apiserver 
tcp6       0      0 :::31471                :::*                    LISTEN      3344/kube-proxy     
tcp6       0      0 :::10255                :::*                    LISTEN      1526/kubelet        
tcp6       0      0 :::10256                :::*                    LISTEN      3344/kube-proxy     
tcp6       0      0 :::48080                :::*                    LISTEN      1047/docker-proxy-c 
tcp6       0      0 :::22                   :::*                    LISTEN      472/sshd            
tcp6       0      0 :::13306                :::*                    LISTEN      1081/docker-proxy-c 
```

## Development

Make binary
```
[vagrant@localhost go-to-docker]$ make go-install
go install -v ./
github.com/tangfeixiong/go-to-docker/pkg/ui/data/swagger
github.com/tangfeixiong/go-to-docker/pkg/server
github.com/tangfeixiong/go-to-docker/cmd
github.com/tangfeixiong/go-to-docker
```

Make dockernized
```
[vagrant@localhost go-to-docker]$ make
### snip ###
[vagrant@localhost go-to-docker]$ docker run -d \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  --privileged=true \
  -p 10052:10052 \
  -e DOCKER_CONFIG_JSON='{"auths":{"127.0.0.1:5000":{"auth":"<basicauth-base64-encoded>","email":""}}}' \
  -e REGISTRY_CERTS_JSON='{"127.0.0.1:5000":{"ca_base64":"<cacert-base64-encoded>"}}' \
  --name=go-to-docker docker.io/tangfeixiong/go-to-docker:0.1
24be50de5ed082409ba98560cff37f3ba31e1eda82ace02b98ec83ea8cce680e
```

For `DOCKER_CONFIG_JSON` environment
```
[vagrant@localhost go-to-docker]$ cat ~/.docker/config.json | tr -d '\n\t '
{"auths":{"...":{"auth":"...","email":""},...}}}
```

For `REGISTRY_CERTS_JSON` environment
```
[vagrant@localhost go-to-docker]$ echo "{\"127.0.0.1:5000\":{\"ca_base64\":\"$(base64 -w 0 /etc/docker/certs.d/127.0.0.1:5000/ca.crt)\"}}"
```


### Docker API

Examples with docker client, engine-api, go-dockerclient, and native json api

* docker client - https://github.com/docker/docker/tree/master/client
* go-dockerclient - https://github.com/fsouza/go-dockerclient
* engine-api - https://github.com/docker/engine-api
* docker-registry-client - https://github.com/heroku/docker-registry-client

### Inspired

https://github.com/grpc-ecosystem/grpc-gateway

https://github.com/philips/grpc-gateway-example

https://github.com/jcbsmpsn/golang-https-example

### Deployment

[Docker Compose](https://docs.docker.com/compose/reference/overview/)

Download from [Github Release Repositories](https://github.com/docker/compose/releases)
```
$ curl -jkSL https://github.com/docker/compose/releases/download/1.15.0/docker-compose-Linux-x86_64 -O
$ chmod +x docker-compose-Linux-x86_64
$ ./docker-compose-Linux-x86_64 --help
```

Etcd download from [Github Release Repositories](https://github.com/coreos/etcd/releases)
```
$ curl -jkSL https://github.com/coreos/etcd/releases/download/v3.2.4/etcd-v3.2.4-linux-amd64.tar.gz -O
$ tar -C /home/vagrant/go/bin/ -zxf etcd-v3.2.4-linux-amd64.tar.gz --strip-components=1 etcd-v3.2.4-linux-amd64/etcd etcd-v3.2.4-linux-amd64/etcdctl
$ etcd --version
etcd Version: 3.2.4
Git SHA: c31bec0
Go Version: go1.8.3
Go OS/Arch: linux/amd64
$ etcdctl --version
etcdctl version: 3.2.4
API version: 2
```

Example
```
[vagrant@localhost go-to-docker]$ sudo mkdir -p /srv/etc/data

[vagrant@localhost go-to-docker]$ hostname -I | awk '{print $2}'
172.17.4.50

[vagrant@localhost go-to-docker]$ sed 's/$NODE1/172.17.4.50/g' docker-compose.yml.sed > docker-compose.yml

[vagrant@localhost go-to-docker]$ docker-compose up -d
```
