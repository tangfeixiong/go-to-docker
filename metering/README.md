# Instruction

## Developement

Refer to [Makefile](./Makefile), for example:

Make _bin_
```
[vagrant@bogon metering]$ make go-build
### snip ###
github.com/tangfeixiong/go-to-docker/metering/pb
github.com/tangfeixiong/go-to-docker/metering/pkg/exporter
github.com/tangfeixiong/go-to-docker/metering/pkg/server
github.com/tangfeixiong/go-to-docker/metering/cmd
github.com/tangfeixiong/go-to-docker/metering
```

Build _exporter_
```
[vagrant@bogon metering]$ make docker-build-exporter
Sending build context to Docker daemon 7.434 MB
Step 1 : FROM busybox
 ---> d20ae45477cb
Step 2 : LABEL "maintainer" "tangfeixiong <tangfx128@gmail.com>" "project" "https://github.com/tangfeixiong/go-to-docker" "name" "metering" "version" "alpha" "namespace" "stackdocker" "annotation" '{"stackdocker.io/created-by":"n/a"}' "tag" "busybox golang"
 ---> Running in e4aa47eaec34
 ---> bd08285c80ad
Removing intermediate container e4aa47eaec34
Step 3 : COPY bin/metering /
 ---> 192c370de0bd
Removing intermediate container 49d453214004
Step 4 : ENV METERING_NAME_URLS cadvisor=http://localhost:8080 METRICS_COLLECTOR_RPC grpc=localhost:12305 METRICS_STORAGE_DRIVER elasticsearch=http://localhost:9200 DOCKER_CONFIG_JSON '{"auths": {"localhost:5000": {"auth": "","email": ""}}}' REGISTRY_CERTS_JSON '{"localhost:5000": {"ca_base64": "", "crt_base64": "", "key_base64": ""}}'
 ---> Running in fa6be91b2e2e
 ---> 75bbd5826ebf
Removing intermediate container fa6be91b2e2e
Step 5 : EXPOSE 12315 12316
 ---> Running in 0d7302fc0bd7
 ---> 6cc74e86a302
Removing intermediate container 0d7302fc0bd7
Step 6 : ENTRYPOINT /metering --v=2 --logtostderr
 ---> Running in a74bec09310a
 ---> 37a30c1a035d
Removing intermediate container a74bec09310a
Step 7 : CMD export
 ---> Running in 5bdf0296fe15
 ---> 997873386184
Removing intermediate container 5bdf0296fe15
Successfully built 997873386184
```

Run _exporter_
```
meter_driver=cadvisor=http://$(docker inspect -f {{.NetworkSettings.IPAddress}} cadvisor):8080  docker run -ti -p 12316:12316 -e METERING_DRIVER_LIST=${meter_driver} --rm --name=metric-exporter docker.io/tangfeixiong/metering
Start gRPC on host [::]:12315
So gRPC running
Start gRPC Gateway into host :12315
http on host: [::]:12316
Tick at 2017-10-06 05:20:44.610224173 +0000 UTC
"meter_url:\"http://172.17.0.10:8080\" timestamp_nanosec:1507267244613425429 meter_response:<key:0 value:\"{\\\"num_cores\\\":2,\\\"cpu_frequency_khz\\\":2699969,\\\"memory_capacity\\\":5201813504,\\\"hugepages\\\":null,\\\"machine_id\\\":\\\"5c949bb3146241e09f7e671b0704d4fb\\\",\\\"system_uuid\\\":\\\"284705D4-2625-4A04-8B45-675B98E6E998\\\",\\\"boot_id\\\":\\\"dc1f609e-002d-42a6-9c92-20654a483a47\\\",\\\"filesystems\\\": -------- snip --------
2017/10/06 05:20:47 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: dial tcp [::1]:12305: getsockopt: connection refused"; Reconnecting to {localhost:12305 <nil>}
2017/10/06 05:20:47 Failed to dial localhost:12305: context canceled; please retry.
### snip ###
```

### Deploy

Via _docker-compose_


### Check elasticsearch

Via _curl_
```
[vagrant@localhost elastic.v5]$ docker exec -ti metering_elasticsearch_1 curl -u elastic:changeme http://172.17.4.50:39200/_nodes/http?pretty
{
  "_nodes" : {
    "total" : 1,
    "successful" : 1,
    "failed" : 0
  },
  "cluster_name" : "docker-cluster",
  "nodes" : {
    "nGVvgYeATV2HC_gDoQIiJA" : {
      "name" : "nGVvgYe",
      "transport_address" : "172.19.0.2:9300",
      "host" : "172.19.0.2",
      "ip" : "172.19.0.2",
      "version" : "5.6.2",
      "build_hash" : "57e20f3",
      "roles" : [
        "master",
        "data",
        "ingest"
      ],
      "attributes" : {
        "ml.max_open_jobs" : "10",
        "ml.enabled" : "true"
      },
      "http" : {
        "bound_address" : [
          "[::]:9200"
        ],
        "publish_address" : "172.19.0.2:9200",
        "max_content_length_in_bytes" : 104857600
      }
    }
  }
}
```

Refer to https://github.com/olivere/elastic/wiki/Connection-Problems#how-to-figure-out-connection-problems