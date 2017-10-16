#!/bin/bash -e

case $1 in
    start-awd10-nothing-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd10_nothing_check.py",
  "command": [
      "python", "awd10_nothing_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd10_nothing",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd11-maccms-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd11_maccms_check.py",
  "command": [
      "python", "awd11_maccms_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd11_maccms",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd12-phpsqllitecms-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd12_phpsqllitecms_check.py",
  "command": [
      "python", "awd12_phpsqllitecms_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd12_phpsqllitecms",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd1-lemon-cms-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd1_lemon_cms_check.py",
  "command": [
      "python", "awd1_lemon_cms_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd1_lemon_cms",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd1-xmanweb2-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd1_xmanweb2_check.py",
  "command": [
      "python", "awd1_xmanweb2_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd1_xmanweb2",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd2-daydayweb-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd2_daydayweb_check.py",
  "command": [
      "python", "awd2_daydayweb_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd2_daydayweb",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd2-dynpage-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd2_dynpage_check.py",
  "command": [
      "python", "awd2_dynpage_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd2_dynpage",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd3-electronics-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd3_electronics_check.py",
  "command": [
      "python", "awd3_electronics_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd3_electronics",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd3-shadow-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd3_shadow_check.py",
  "command": [
      "python", "awd3_shadow_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd3_shadow",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd4-chinaz-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd4_chinaz_check.py",
  "command": [
      "python", "awd4_chinaz_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd4_chinaz",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd4-tomcat-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd4_tomcat_check.py",
  "command": [
      "python", "awd4_tomcat_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd4_tomcat",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd5-babyblog-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd5_babyblog_check.py",
  "command": [
      "python", "awd5_babyblog_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd5_babyblog",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd5-gracer-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd5_gracer_check.py",
  "command": [
      "python", "awd5_gracer_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd5_gracer",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd6-cms-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd6_cms_check.py",
  "command": [
      "python", "awd6_cms_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd6_cms",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd7-upload-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd7_upload_check.py",
  "command": [
      "python", "awd7_upload_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd7_upload",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd8-blog-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd8_blog_check.py",
  "command": [
      "python", "awd8_blog_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd8_blog",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-awd9-money-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "awd9_money_check.py",
  "command": [
      "python", "awd9_money_check.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "awd9_money",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-web1-1-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "web1-checker.py",
  "command": [
      "python", "web1-checker.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "web1-1",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    start-web2-1-check)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "web2-checker.py",
  "command": [
      "python", "web2-checker.py"
    ],
  "args": [
      "--host=$(ip)",
      "--port=$(port)"
    ],
  "env": [
  ],
  "conf":
    {
    },
  "work_dir": "web2-1",
  "periodic": 3,
  "duration": 10,
  "dest_configurations":
    {
      "team1":
        {
          "name": "container1",
          "tpl":
            {
              "ip": "localhost",
              "port": "80"
            }
        }
    }
}'
        ;;
    create)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "web1-check.py",
  "command": [
      "python", "web1-check.py"
    ],
  "args": [
    ],
  "env": [
  ],
  "conf":
    {
      "hosts.list": "bG9jYWxob3N0Cg=="
    },
  "work_dir": "web1check",
  "periodic": 3
}'
        ;;
    reap)
	    curl -X GET http://172.17.4.50:10062/v1/checkactions/awd10_nothing_check.py
		;;
    update)
	    curl -X PUT http://172.17.4.50:10062/v1/checkactions/awd10_nothing_check.py -d \
'{
  "name": "awd10_nothing_check.py",
  "periodic": 6,
  "duration": 12
}'
        ;;
    delete)
	    curl -X DELETE http://172.17.4.50:10062/v1/checkactions/awd10_nothing_check.py
		;;
    *)
        echo "Valid test: create, reap, update, delete"
        ;;
	
esac



