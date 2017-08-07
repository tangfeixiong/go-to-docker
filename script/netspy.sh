#!/bin/sh

set -e

if [[ $# -lt 1 ]]; then
    echo "container required" > /dev/stderr
	exit 1
fi

if [[ -z $(docker ps -qf name=$1) ]]; then
    echo "container not found" > /dev/stderr
    exit 10
fi

PID=$(docker inspect -f '{{.State.Pid}}' $1)

if [[ -d /proc/$PID ]]; then 
    sudo ln -sf /proc/${PID}/ns/net /var/run/netns/${PID}
	sudo ip netns exec ${PID} ip link list
	echo
	sudo ip netns exec ${PID} ip addr show
	echo
	sudo ip netns exec ${PID} ip route list
    sudo rm /var/run/netns/${PID}
fi