
# https://coreos.com/etcd/docs/latest/op-guide/container.html#docker
# export NODE1=192.168.1.21
#
# docker run \
#   -p 2379:2379 \
#   -p 2380:2380 \
#   --volume=${DATA_DIR}:/etcd-data \
#   --name etcd quay.io/coreos/etcd:latest \
#   /usr/local/bin/etcd \
#   --data-dir=/etcd-data --name node1 \
#   --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://${NODE1}:2380 \
#   --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://${NODE1}:2379 \
#   --initial-cluster node1=http://${NODE1}:2380
#
# etcdctl --endpoints=http://${NODE1}:2379 member list

version: '2'

services:
  etcd:
    command: 
    - /usr/local/bin/etcd
    - --data-dir=/etcd-data
    - --name
    - node1
    - --initial-advertise-peer-urls
    - http://$NODE1:2380
    - --listen-peer-urls
    - http://$NODE1:2380
    - --advertise-client-urls
    - http://$NODE1:2379
    - --listen-client-urls 
    - http://$NODE1:2379
    - --initial-cluster 
    - node1=http://$NODE1:2380
    image: quay.io/coreos/etcd:v3.2.4
    ports:
    - "2379:2379"
    - "2380:2380"
    restart: always
    volumes:
    - /srv/etcd/data:/etcd-data

  go-to-docker:
    depends_on:
    - etcd
    environment:
    - DOCKER_CONFIG_JSON='{"auths":{}}'
    - REGISTRY_CERTS_JSON='{}'
    - ETCD_HOST=etcd
    image: docker.io/tangfeixiong/go-to-docker:0.1.1
    ports:
    - "10052:10052"
    privileged: true
    restart: always
    volumes:
    - /var/run/docker.sock:/var/run/docker.sock:ro
