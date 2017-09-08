#!/bin/bash -e

case $1 in
    create)
	    curl -X POST http://172.17.4.50:10062/v1/checkactions -d \
'{
  "name": "web1check.py",
  "command": [
      "python", "web1check.py"
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
	    curl -X GET http://172.17.4.50:10062/v1/checkactions/web1check.py
		;;
    update)
	    curl -X PUT http://172.17.4.50:10062/v1/checkactions/web1check.py -d \
'{
  "name": "web1check.py",
  "args": [
    ],
  "env": [
  ],
  "conf":
    {
      "hosts.list": "MTI3LjAuMC4xCg=="
    },
  "work_dir": "web1check",
  "periodic": 5
}'
        ;;
    delete)
	    curl -X DELETE http://172.17.4.50:10062/v1/checkactions/web1check.py
		;;
    *)
        echo "Valid test: create, reap, update, delete"
        ;;
	
esac



