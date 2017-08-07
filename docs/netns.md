
### For docker (Fedora23)

Refer to  https://docs.docker.com/v1.7/articles/networking/

A networking container
```
[vagrant@localhost kopos]$ docker inspect -f '{{.State.Pid}}' kopos
17020

[vagrant@localhost ~]$ sudo ls -l /proc/2019/ns
总用量 0
lrwxrwxrwx. 1 root root 0 7月  12 21:36 ipc -> ipc:[4026532204]
lrwxrwxrwx. 1 root root 0 7月  12 21:36 mnt -> mnt:[4026532202]
lrwxrwxrwx. 1 root root 0 7月  12 21:36 net -> net:[4026532207]
lrwxrwxrwx. 1 root root 0 7月  12 21:36 pid -> pid:[4026532205]
lrwxrwxrwx. 1 root root 0 7月  12 21:36 user -> user:[4026531837]
lrwxrwxrwx. 1 root root 0 7月  12 21:36 uts -> uts:[4026532203]

[vagrant@localhost kopos]$ sudo ls /proc/2019/ns/net
/proc/17020/ns/net

[vagrant@localhost kopos]$ sudo ln -s /proc/2019/ns/net /var/run/netns/2019

[vagrant@localhost ~]$ sudo ip netns exec 2019 ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
51: eth0@if52: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 3e:4e:52:68:43:ff brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.4/22 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::3c4e:52ff:fe68:43ff/64 scope link 
       valid_lft forever preferred_lft forever
```

### Network Virtualization

Device
```
[vagrant@localhost echopb]$ lspci | egrep -i --color 'network|ethernet'
00:03.0 Ethernet controller: Intel Corporation 82540EM Gigabit Ethernet Controller (rev 02)
00:08.0 Ethernet controller: Intel Corporation 82540EM Gigabit Ethernet Controller (rev 02)

[vagrant@localhost echopb]$ lshw -class network
  *-network DISABLED      
       description: Wireless interface
       product: Ultimate N WiFi Link 5300
       vendor: Intel Corporation
       physical id: 0
       bus info: pci@0000:0c:00.0
       logical name: wlan0
       version: 00
       serial: 00:21:6a:ca:9b:10
       width: 64 bits
       clock: 33MHz
       capabilities: pm msi pciexpress bus_master cap_list ethernet physical wireless
       configuration: broadcast=yes driver=iwlwifi driverversion=3.2.0-0.bpo.1-amd64 firmware=8.83.5.1 build 33692 latency=0 link=no multicast=yes wireless=IEEE 802.11abgn
       resources: irq:46 memory:f1ffe000-f1ffffff
  *-network
       description: Ethernet interface
       product: NetXtreme BCM5761e Gigabit Ethernet PCIe
       vendor: Broadcom Corporation
       physical id: 0
       bus info: pci@0000:09:00.0
       logical name: eth0
       version: 10
       serial: b8:ac:6f:65:31:e5
       size: 1GB/s
       capacity: 1GB/s
       width: 64 bits
       clock: 33MHz
       capabilities: pm vpd msi pciexpress bus_master cap_list ethernet physical tp 10bt 10bt-fd 100bt 100bt-fd 1000bt 1000bt-fd autonegotiation
       configuration: autonegotiation=on broadcast=yes driver=tg3 driverversion=3.121 duplex=full firmware=5761e-v3.71 ip=192.168.1.5 latency=0 link=yes multicast=yes port=twisted pair speed=1GB/s
       resources: irq:48 memory:f1be0000-f1beffff memory:f1bf0000-f1bfffff

[vagrant@localhost echopb]$ ip link show
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP mode DEFAULT group default qlen 1000
    link/ether 08:00:27:46:54:e7 brd ff:ff:ff:ff:ff:ff
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP mode DEFAULT group default qlen 1000
    link/ether 08:00:27:5a:1a:a4 brd ff:ff:ff:ff:ff:ff
4: br-3ba1465afaf1: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN mode DEFAULT group default 
    link/ether 02:42:be:c7:20:bb brd ff:ff:ff:ff:ff:ff
5: br-cee016075c51: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default 
    link/ether 02:42:06:68:b8:d7 brd ff:ff:ff:ff:ff:ff
6: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default 
    link/ether 02:42:fa:1d:73:42 brd ff:ff:ff:ff:ff:ff
8: veth3856bd8@if7: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP mode DEFAULT group default 
    link/ether ba:f5:da:bd:43:1b brd ff:ff:ff:ff:ff:ff link-netnsid 4
10: veth2dfea95@if9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master br-cee016075c51 state UP mode DEFAULT group default 
    link/ether 22:3c:91:4c:df:ed brd ff:ff:ff:ff:ff:ff link-netnsid 1
12: veth6ab2a92@if11: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP mode DEFAULT group default 
    link/ether 46:88:70:32:67:fa brd ff:ff:ff:ff:ff:ff link-netnsid 2
14: veth294ceb0@if13: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master br-cee016075c51 state UP mode DEFAULT group default 
    link/ether 72:4d:f1:9d:99:99 brd ff:ff:ff:ff:ff:ff link-netnsid 3
16: vethc79d386@if15: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master br-cee016075c51 state UP mode DEFAULT group default 
    link/ether 7a:30:e9:b7:13:7f brd ff:ff:ff:ff:ff:ff link-netnsid 6
18: veth97b96c2@if17: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP mode DEFAULT group default 
    link/ether 42:9b:4b:52:25:b5 brd ff:ff:ff:ff:ff:ff link-netnsid 0
20: veth38bdfc6@if19: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP mode DEFAULT group default 
    link/ether 32:fc:21:62:df:87 brd ff:ff:ff:ff:ff:ff link-netnsid 5
22: veth08409b0@if21: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP mode DEFAULT group default 
    link/ether 7a:91:4d:95:89:14 brd ff:ff:ff:ff:ff:ff link-netnsid 7
62: vethc5d9544@if61: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP mode DEFAULT group default 
    link/ether 82:b1:10:61:6e:7e brd ff:ff:ff:ff:ff:ff link-netnsid 8

[vagrant@localhost echopb]$ cat /proc/net/dev
Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
br-3ba1465afaf1:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0
  eth0: 355574235  406658    0    0    0     0          0         0 147658138  230132    0    0    0     0       0          0
docker0: 251568253  386194    0    0    0     0          0         0 1120552678  437038    0    0    0     0       0          0
  eth1: 8216279   58333    0    0    0     0          0         0 15157352   43161    0    0    0     0       0          0
veth6ab2a92:   27298     100    0    0    0     0          0         0   150236     933    0    0    0     0       0          0
vethc79d386: 6945457031 8429888    0    0    0     0          0         0 4018308869 10111049    0    0    0     0       0          0
br-cee016075c51:  969320     588    0    0    0     0          0         0   213453     796    0    0    0     0       0          0
veth38bdfc6: 6394075   39743    0    0    0     0          0         0 276473090   63919    0    0    0     0       0          0
vethc5d9544:    8676      62    0    0    0     0          0         0    11669      78    0    0    0     0       0          0
    lo: 7081739   19879    0    0    0     0          0         0  7081739   19879    0    0    0     0       0          0
veth2dfea95: 3979792681 10059753    0    0    0     0          0         0 6931477472 8376040    0    0    0     0       0          0
veth3856bd8:  933807   12607    0    0    0     0          0         0 115348018   15425    0    0    0     0       0          0
veth294ceb0: 38304821   50519    0    0    0     0          0         0 13026847   53559    0    0    0     0       0          0
veth08409b0: 172307690    6879    0    0    0     0          0         0   485059    6009    0    0    0     0       0          0
veth97b96c2: 72097301  231387    0    0    0     0          0         0 514015370  245170    0    0    0     0       0          0

[vagrant@localhost ~]$ ip addr list
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:46:54:e7 brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global dynamic eth0
       valid_lft 86243sec preferred_lft 86243sec
    inet6 fe80::a00:27ff:fe46:54e7/64 scope link 
       valid_lft forever preferred_lft forever
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:5a:1a:a4 brd ff:ff:ff:ff:ff:ff
    inet 172.17.4.50/24 brd 172.17.4.255 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe5a:1aa4/64 scope link 
       valid_lft forever preferred_lft forever
4: br-3ba1465afaf1: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 02:42:d9:cb:55:d9 brd ff:ff:ff:ff:ff:ff
    inet 10.1.0.1/24 scope global br-3ba1465afaf1
       valid_lft forever preferred_lft forever
5: br-cee016075c51: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:98:c6:78:30 brd ff:ff:ff:ff:ff:ff
    inet 172.18.0.1/16 scope global br-cee016075c51
       valid_lft forever preferred_lft forever
    inet6 fe80::42:98ff:fec6:7830/64 scope link 
       valid_lft forever preferred_lft forever
6: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:01:74:cc:7e brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/22 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:1ff:fe74:cc7e/64 scope link 
       valid_lft forever preferred_lft forever
24: veth548d6f6@if23: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master br-cee016075c51 state UP group default 
    link/ether ce:68:9a:f1:bf:7c brd ff:ff:ff:ff:ff:ff link-netnsid 1
26: veth54ebb65@if25: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master br-cee016075c51 state UP group default 
    link/ether c6:7d:22:72:3c:de brd ff:ff:ff:ff:ff:ff link-netnsid 3
28: veth7d8aff1@if27: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether 0e:f2:07:82:b3:fb brd ff:ff:ff:ff:ff:ff link-netnsid 4
30: vethcb2298a@if29: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master br-cee016075c51 state UP group default 
    link/ether 06:4f:88:29:82:48 brd ff:ff:ff:ff:ff:ff link-netnsid 5
32: veth6f8afba@if31: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether a6:4d:70:6c:0d:c5 brd ff:ff:ff:ff:ff:ff link-netnsid 6
36: vethb9eeaf6@if35: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether 7a:18:bf:02:67:48 brd ff:ff:ff:ff:ff:ff link-netnsid 9
38: vethd92b9d8@if37: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether da:74:c8:f7:5a:fc brd ff:ff:ff:ff:ff:ff link-netnsid 10
40: vethda26728@if39: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether 7a:ac:68:99:66:c3 brd ff:ff:ff:ff:ff:ff link-netnsid 11
52: vethedb1a2c@if51: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether 76:78:1c:8a:d1:db brd ff:ff:ff:ff:ff:ff link-netnsid 0
60: veth915b98d@if59: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether d2:a4:82:17:3c:36 brd ff:ff:ff:ff:ff:ff link-netnsid 2
```

Docker inspect
```
[vagrant@localhost ~]$ docker inspect kopos
[
    {
        "Id": "044bddd75e205c0eeba94ca32af92827793ce26f0891420776dcb5a75412dd22",
        "Created": "2017-07-11T06:43:25.909008618Z",
        "Path": "/kopos",
        "Args": [
            "serve",
            "-k",
            "/Users/fanhongling/Downloads/tmp/vagrant",
            "-u",
            "vagrant",
            "--loglevel",
            "2"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 2019,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2017-07-11T06:43:27.000279365Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:e2458293ee21f032316acbce27045f800062597ef36494774cc6306b61441e18",
        "ResolvConfPath": "/var/lib/docker/containers/044bddd75e205c0eeba94ca32af92827793ce26f0891420776dcb5a75412dd22/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/044bddd75e205c0eeba94ca32af92827793ce26f0891420776dcb5a75412dd22/hostname",
        "HostsPath": "/var/lib/docker/containers/044bddd75e205c0eeba94ca32af92827793ce26f0891420776dcb5a75412dd22/hosts",
        "LogPath": "",
        "Name": "/kopos",
        "RestartCount": 0,
        "Driver": "devicemapper",
        "MountLabel": "system_u:object_r:svirt_sandbox_file_t:s0:c409,c842",
        "ProcessLabel": "system_u:system_r:svirt_lxc_net_t:s0:c409,c842",
        "AppArmorProfile": "",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "journald",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {
                "10001/tcp": [
                    {
                        "HostIp": "",
                        "HostPort": "10001"
                    }
                ]
            },
            "RestartPolicy": {
                "Name": "always",
                "MaximumRetryCount": 0
            },
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "ShmSize": 67108864,
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": null,
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "KernelMemory": 0,
            "Memory": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": -1,
            "OomKillDisable": false,
            "PidsLimit": 0,
            "Ulimits": null
        },
        "GraphDriver": {
            "Name": "devicemapper",
            "Data": {
                "DeviceId": "13117",
                "DeviceName": "docker-253:0-1310998-56b610aa85cd13fbf62469c41b905f46342df04049d362a8c45dc6b5f8583880",
                "DeviceSize": "107374182400"
            }
        },
        "Mounts": [],
        "Config": {
            "Hostname": "044bddd75e20",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "10000/tcp": {},
                "10001/tcp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "OS_PASSWORD=7uj8ik",
                "MOCK_HOST=172.17.4.50",
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "OS_TENANT_NAME=admin",
                "OS_USERNAME=admin",
                "OS_AUTH_URL=http://10.121.198.2:5000/v2.0/",
                "OS_AUTH_STRATEGY=keystone",
                "OS_REGION_NAME=openstack",
                "HV_USERNAME=fake user",
                "HV_PASSWORD=fake secret",
                "HV_KEYPATH=/Users/fanhongling/Downloads/tmp/vagrant",
                "HV_AGENT=false",
                "HV_GRPC=false"
            ],
            "Cmd": [
                "-k",
                "/Users/fanhongling/Downloads/tmp/vagrant",
                "-u",
                "vagrant",
                "--loglevel",
                "2"
            ],
            "Image": "docker.io/tangfeixiong/go-to-openstack-bootcamp:0.1",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": [
                "/kopos",
                "serve"
            ],
            "OnBuild": null,
            "Labels": {
                "maintainer": "tangfeixiong \u003ctangfx128@gmail.com\u003e",
                "name": "kopos",
                "project": "https://github.com/tangfeixiong/go-to-openstack-bootcamp",
                "version": "0.1"
            }
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "203254cad31c889075a790eb7b8e1aa275ada64efefec3d613f8e7430a29566a",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "10000/tcp": null,
                "10001/tcp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "10001"
                    }
                ]
            },
            "SandboxKey": "/var/run/docker/netns/203254cad31c",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "a986028219cf22c91d8c4c2dae168433985d660d2ad98c36a9dc85108738ae58",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.4",
            "IPPrefixLen": 22,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:04",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "203b62e09f95d1c865c8a00ad3572b7187d2507ab8fe2e8b71acb951f631c42a",
                    "EndpointID": "a986028219cf22c91d8c4c2dae168433985d660d2ad98c36a9dc85108738ae58",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.4",
                    "IPPrefixLen": 22,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:04"
                }
            }
        }
    }
]

[vagrant@localhost ~]$ docker network inspect bridge
[
    {
        "Name": "bridge",
        "Id": "203b62e09f95d1c865c8a00ad3572b7187d2507ab8fe2e8b71acb951f631c42a",
        "Scope": "local",
        "Driver": "bridge",
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.17.0.1/22",
                    "Gateway": "172.17.0.1"
                }
            ]
        },
        "Containers": {
            "044bddd75e205c0eeba94ca32af92827793ce26f0891420776dcb5a75412dd22": {
                "Name": "kopos",
                "EndpointID": "a986028219cf22c91d8c4c2dae168433985d660d2ad98c36a9dc85108738ae58",
                "MacAddress": "02:42:ac:11:00:04",
                "IPv4Address": "172.17.0.4/22",
                "IPv6Address": ""
            },
            "5f5980fb661e5a7309423c5d25e12538eae3237012be8815025e8f132f9c2144": {
                "Name": "nexus",
                "EndpointID": "c68b092bc139e0475164cfccdd2905d16a6454dded2e8e7676babdc9da6ac2f7",
                "MacAddress": "02:42:ac:11:00:05",
                "IPv4Address": "172.17.0.5/22",
                "IPv6Address": ""
            },
            "87e7cacd5279180909d2b2c58ec6574818c7f9c13f18d8464c00b42dc0e3923b": {
                "Name": "registry2_registry_1",
                "EndpointID": "d89e7e62e0024ea4195f29edf461711d9114f1e043c3d5864a790d521013f107",
                "MacAddress": "02:42:ac:11:00:03",
                "IPv4Address": "172.17.0.3/22",
                "IPv6Address": ""
            },
            "b1e178fa584eea379eb60a2f0947ec21247aba96378e8e073d82ed5d393a301c": {
                "Name": "gofileserver",
                "EndpointID": "553c201248e153699ba66c97eb74737bfa2fe5e0e87fed0ac748fed6541fe31a",
                "MacAddress": "02:42:ac:11:00:07",
                "IPv4Address": "172.17.0.7/22",
                "IPv6Address": ""
            },
            "b86cbb644ab9a2039824db8dab48ca444edabf7c2a0ee71bb48bb731776e3d30": {
                "Name": "tomcat",
                "EndpointID": "62b278db83f7da45cb8deaf3a139dd1beaecb75390546b2a9fc7b995897d47e2",
                "MacAddress": "02:42:ac:11:00:08",
                "IPv4Address": "172.17.0.8/22",
                "IPv6Address": ""
            },
            "d3c9244ac592677d2c961cc556fd2122eeb52c7654d4f3645577d2c148484f40": {
                "Name": "registry2_dockerauth_1",
                "EndpointID": "ad8d966bad61a014d9559c34ca4fa9c71e78c8e7ba9d4aab27336371fbf9ca29",
                "MacAddress": "02:42:ac:11:00:02",
                "IPv4Address": "172.17.0.2/22",
                "IPv6Address": ""
            },
            "d7127155027542082be86c2b843e8a7c7f40fecaf8df8e7ed1300e86116831f9": {
                "Name": "jenkins",
                "EndpointID": "0cf5f5c4879f66da4c80601bdf7d15c7f91b6b680ae509aa28888fa90dc6fd2f",
                "MacAddress": "02:42:ac:11:00:06",
                "IPv4Address": "172.17.0.6/22",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        }
    }
]

```

Linux bridge
```
[vagrant@localhost ~]$ brctl show
bridge name	bridge id		STP enabled	interfaces
br-3ba1465afaf1		8000.0242d9cb55d9	no		
br-cee016075c51		8000.024298c67830	no		veth548d6f6
							veth54ebb65
							vethcb2298a
docker0		8000.02420174cc7e	no		veth6f8afba
							veth7d8aff1
							veth915b98d
							vethb9eeaf6
							vethd92b9d8
							vethda26728
							vethedb1a2c
[vagrant@localhost ~]$ brctl showmacs docker0
port no	mac addr		is local?	ageing timer
  1	0e:f2:07:82:b3:fb	yes		   0.00
  1	0e:f2:07:82:b3:fb	yes		   0.00
  3	76:78:1c:8a:d1:db	yes		   0.00
  3	76:78:1c:8a:d1:db	yes		   0.00
  4	7a:18:bf:02:67:48	yes		   0.00
  4	7a:18:bf:02:67:48	yes		   0.00
  6	7a:ac:68:99:66:c3	yes		   0.00
  6	7a:ac:68:99:66:c3	yes		   0.00
  2	a6:4d:70:6c:0d:c5	yes		   0.00
  2	a6:4d:70:6c:0d:c5	yes		   0.00
  7	d2:a4:82:17:3c:36	yes		   0.00
  7	d2:a4:82:17:3c:36	yes		   0.00
  5	da:74:c8:f7:5a:fc	yes		   0.00
  5	da:74:c8:f7:5a:fc	yes		   0.00
[vagrant@localhost ~]$ brctl showstp docker0
docker0
 bridge id		8000.02420174cc7e
 designated root	8000.02420174cc7e
 root port		   0			path cost		   0
 max age		  20.00			bridge max age		  20.00
 hello time		   2.00			bridge hello time	   2.00
 forward delay		  15.00			bridge forward delay	  15.00
 ageing time		 300.00
 hello timer		   0.00			tcn timer		   0.00
 topology change timer	   0.00			gc timer		  48.73
 flags			


veth6f8afba (2)
 port id		8002			state		     forwarding
 designated root	8000.02420174cc7e	path cost		   2
 designated bridge	8000.02420174cc7e	message age timer	   0.00
 designated port	8002			forward delay timer	   0.00
 designated cost	   0			hold timer		   0.00
 flags			

veth7d8aff1 (1)
 port id		8001			state		     forwarding
 designated root	8000.02420174cc7e	path cost		   2
 designated bridge	8000.02420174cc7e	message age timer	   0.00
 designated port	8001			forward delay timer	   0.00
 designated cost	   0			hold timer		   0.00
 flags			

veth915b98d (7)
 port id		8007			state		     forwarding
 designated root	8000.02420174cc7e	path cost		   2
 designated bridge	8000.02420174cc7e	message age timer	   0.00
 designated port	8007			forward delay timer	   0.00
 designated cost	   0			hold timer		   0.00
 flags			

vethb9eeaf6 (4)
 port id		8004			state		     forwarding
 designated root	8000.02420174cc7e	path cost		   2
 designated bridge	8000.02420174cc7e	message age timer	   0.00
 designated port	8004			forward delay timer	   0.00
 designated cost	   0			hold timer		   0.00
 flags			

vethd92b9d8 (5)
 port id		8005			state		     forwarding
 designated root	8000.02420174cc7e	path cost		   2
 designated bridge	8000.02420174cc7e	message age timer	   0.00
 designated port	8005			forward delay timer	   0.00
 designated cost	   0			hold timer		   0.00
 flags			

vethda26728 (6)
 port id		8006			state		     forwarding
 designated root	8000.02420174cc7e	path cost		   2
 designated bridge	8000.02420174cc7e	message age timer	   0.00
 designated port	8006			forward delay timer	   0.00
 designated cost	   0			hold timer		   0.00
 flags			

vethedb1a2c (3)
 port id		8003			state		     forwarding
 designated root	8000.02420174cc7e	path cost		   2
 designated bridge	8000.02420174cc7e	message age timer	   0.00
 designated port	8003			forward delay timer	   0.00
 designated cost	   0			hold timer		   0.00
 flags			

```