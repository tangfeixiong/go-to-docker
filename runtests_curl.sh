#!/bin/bash -e

host="127.0.0.1:10053"
if [[ ${1} =~ --addr=.* ]]; then
    host=${1#*=}
    shift
fi

case $1 in
    docker-network-create)
        curl -X POST http://$host/v1/docker-network-create -H "Content-Type: application/json" -d \
'{
  "name": "",
  "network_create":
    {
	}
}'
        ;;		
    docker-network-rm)
        id=""
        [ $# > 1 ] && id=$2
        curl -X POST http://$host/v1/docker-network-rm -H "Content-Type: application/json" -d \
"{
  \"id\": \"$id\"
}"
        ;;		
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
    psContainers)
        curl -X POST http://172.17.4.50:10052/v1/process-statuses -d \
'{
    "options":
	  {
		"filter":
		  {
			"fields":
			  {
			    "name":
				  {
					"value":
					  {
				        "go-to-docker": true
				      }
			      }	
			  }
		  }
	  }	
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
    network-ls)
        curl -X GET http://172.17.4.50:10052/v1/networks
        ;;		
    network-create)
        curl -X POST http://172.17.4.50:10052/v1/network-creation -d \
'{
  "network_create_request":
    {
      "network_create":
        {
          "ipam":
            {
              "config":[
                {
                  "subnet": "172.25.0.0/16",
                  "ip_range": "172.25.1.0/24",
                  "gateway": "172.25.1.1"
                }
              ]
            },
          "options":
            {
			  "com.docker.network.driver.mtu": "1500"
            },
          "labels":
            {
            }
        },
      "name": "brtest0"
	}
}'
        ;;		
    *)
        echo "test method required!"
        ;;
	
esac



