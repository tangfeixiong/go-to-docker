#!/bin/bash -e

case $1 in
    test-runcontainer)
	    curl -X POST http://172.17.4.50:10052/v1/containers -d \
'{
  "config":
    {
      "image": "nginx",
      "cmd": [
        "-c", "printenv && ls /usr/share/nginx/html && nginx -g \"daemon off;\""
      ],
      "entrypoint": [
        "/bin/bash"
      ],
      "env": [
        "GOAPTH=/home/vagrant/go",
        "JAVA_HOME=/opt/jdk1.8.0_112"
      ],
      "exposed_ports":
        {
          "value": 
            {
              "80": "webui"
            }
        },
      "volumes":
        {
          "/usr/share/nginx/html/": "usertest"
        }
    },
  "host_config":
    {
      "binds": [
        "/home/vagrant/:/usr/share/nginx/html/:ro"
      ],
      "port_bindings":
        {
          "value":
            {
              "80":
                {
                  "host_port": "80"
                }
            }
        },
      "resources":
        {
          "memory": 300000000
        }
    },
  "network_config":
    {
    },
  "container_name": "nginx"
}'
        ;;
    test-provisions)
        curl -X POST http://172.17.4.50:10052/v1/provisions -d \
'{
    "name": "test1",
    "namespace": "default",
    "metadata":
	  {
        "categroy_name": "default",
        "class_name": "default",
        "field_name": "default"
      },
    "provisionings": [
      {
        "config": {
          "image": "docker.io/tangfeixiong/netcat-hello-http",
          "exposed_ports": {
            "value": {
              "80": "webui"
            }
          }
        },
        "host_config": {
          "port_bindings": {
            "value": {
              "80": {
                "host_port": "80"
              }
            }
          }
        },
        "network_config": {
            
        },
        "container_name": "netcat-hello-http"
      }
    ]
}'
        ;;

    test-instantiation)
        curl -X POST http://172.17.4.50:10052/v1/reap-instantiation -d \
'{
  "name": "test1",
  "namespace": "default",
  "metadata": {
    "categroy_name": "default",
	"class_name": "default",
	"field_name": "default"
  }
}'
        ;;

    test-terminations)
        curl -X POST http://172.17.4.50:10052/v1/terminations -d \
'{
    "name": "test1",
	"namespace": "default",
	"metadata":
	  {
	    "categroy_name": "default",
		"class_name": "default",
		"field_name": "default"
      },
	"provisionings": [
	
	]
}'
        ;;

    test-registry)
	    curl -X POST http://172.17.4.50:10052/v1/reap-registry -d \
'{
    "registries":[
	  {
	    "name": "172.17.4.50:5000"
	  }
	]
}'
        ;;
		
    *)
        echo "test method required!"
        ;;
	
esac



