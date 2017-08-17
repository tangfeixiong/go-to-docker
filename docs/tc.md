
http://backreference.org/2014/06/17/port-mirroring-with-linux-bridges/



https://developer.ibm.com/recipes/tutorials/bridge-the-docker-containers-to-external-network/

 root@docker:~# cat /etc/network/interfaces.d/br0
    auto br0
    iface br0 inet static
      address 10.0.189.109
      netmask 255.255.0.0
      gateway 10.0.0.57
      bridge_ports enp0s1
      bridge_fd 0
      bridge_hello 2
      bridge_maxage 12
      bridge_stp off

docker network create --driver=bridge --ip-range=10.0.190.0/24 --subnet=10.0.0.0/16 --aux-address='ip1=10.0.190.1' --aux-address='ip2=10.0.190.2' --aux-address='ip3=10.0.190.3' -o "com.docker.network.bridge.name=br0" br0
