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
        [ $# -gt 1 ] && id=$2
        curl -X POST http://$host/v1/docker-network-rm -H "Content-Type: application/json" -d \
"{
  \"id\": \"$id\"
}"
        ;;		
    docker-container-run)
        if [ $# -eq 1 ]; then
            echo "Usage: runtest_curl.sh docker-container-run tomcat" > /dev/stderr
        elif [ $2 = "tomcat" ]; then
            curl -jkSL https://tomcat.apache.org/tomcat-8.0-doc/appdev/sample/sample.war -o /tmp/sample.war
	        curl -X POST http://$host/v1/docker-container-run -H "Content-Type: application/json" -d \
'{
  "config":
    {
      "image": "tomcat",
      "cmd": [
      ],
      "entrypoint": [
      ],
      "env": [
      ],
      "exposed_ports":
        {
          "internal_map": 
            {
              "8080": {}
            }
        },
      "volumes":
        {
          "/tmp/example": {}
        }
    },
  "host_config":
    {
      "binds": [
        "/tmp/sample.war:/usr/local/tomcat/webapps/sample.war:ro"
      ],
      "network_mode": "stackdocker-brj5lp62cw7",
      "port_bindings":
        {
          "internal_map":
            {
              "8080":
                {
                  "internal_list": [
                    {
                      "host_ip": "",
                      "host_port": ""
                    }
                  ]
                }
            }
        },
      "resources":
        {
          "memory": 300000000
        }
    },
  "networking_config":
    {
      "stackdocker-brj5lp62cw7":
        {
           "ipam_config": {},
           "network_id": "cd1526477a81d27a2b3751903dd268cc92c1bad453f495ac974407ba61f0eb04" 
        }
    },
  "name": "tomcat-example",
  "image_pull_options":
    {
      "registry_auth": ""  
    }
}'
        else 
            echo "If you should test docker image other than tomcat, please let me know" > /dev/stdout
        fi
        ;;
    docker-container-rm)
        id=""
        [ $# > 1 ] && id=$2
        curl -X POST http://$host/v1/docker-container-rm -H "Content-Type: application/json" -d \
"{
  \"id\": \"$id\",
  \"container_remove_options\":
    {
      \"remove_volumes\": true,
      \"remove_links\": false,
      \"force\": true
    }
}"
        ;;		
    docker-image-build)
        curl -X POST http://$host/v1/docker-image-build -H "Content-Type: application/json" -d \
'{
  "build_context": "CkZST00gYWxwaW5lClJVTiBhcGsgYWRkIC0tdXBkYXRlIG5ldGNhdC1vcGVuYnNkICYmIHJtIC1yZiAvdmFyL2NhY2hlL2Fway8qClJVTiBlY2hvIC1lICIjIS9iaW4vc2hcblwgCnNldCAtZVxuXAp3aGlsZSB0cnVlOyBkbyBlY2hvIC1lIFwiSFRUUC8xLjEgMjAwIE9LXG5cbiBcJChkYXRlKSBIZWxsbyB3b3JsZFwiIHwgbmMgLWwgODA7IGRvbmUiID4gL2VudHJ5cG9pbnQuc2ggXAogICAgJiYgY2htb2QgK3ggL2VudHJ5cG9pbnQuc2gKIyBSVU4gdG91Y2ggL2VudHJ5cG9pbnQuc2ggJiYgY2htb2QgK3ggL2VudHJ5cG9pbnQuc2ggJiYgZWNobyAtZSAiIyEvYmluL3NoXG5zZXQgLWVcbndoaWxlIHRydWU7IGRvIG5jIC1sIDgwIDwgaW5kZXguaHRtbDsgZG9uZSIgPiAvZW50cnlwb2ludC5zaApSVU4gZWNobyAtZSAiXG5cCjxodG1sPlwKICAgICAgICA8aGVhZD5cCiAgICAgICAgICAgICAgICA8dGl0bGU+SGVsbG8gUGFnZTwvdGl0bGU+XAogICAgICAgIDwvaGVhZD5cCiAgICAgICAgPGJvZHk+XAogICAgICAgICAgICAgICAgPGgxPkhlbGxvPC9oMT5cCiAgICAgICAgICAgICAgICA8aDI+Q29udGFpbmVyPC9oMj5cCiAgICAgICAgICAgICAgICA8cD5Qb3dlcmVkIGJ5IG5jPC9wPlwKICAgICAgICA8L2JvZHk+XAo8L2h0bWw+XAoiID4gL2luZGV4Lmh0bWwKCkVOVFJZUE9JTlQgWyIvZW50cnlwb2ludC5zaCJdCkVYUE9TRSA4MAo=",
  "image_build_options":
    {
        "tags": ["tangfeixiong/hello-world:netcat-http"],
        "no_cache": true,
        "suppress_output": true,
        "remove": true,
        "force_remove": true,
        "pull_parent": true
	}
}'
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



