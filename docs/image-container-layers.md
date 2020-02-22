# Docker image and container store

Table of contents
- Storage Driver
- Docker image
    1. ImageDB
    1. LayerDB
- Docker container
    1. r/w layer
- Mounts and Volumes

Reference
- https://docs.docker.com/storage/storagedriver/
- https://docs.docker.com/storage/storagedriver/overlayfs-driver
- https://docs.docker.com/v17.09/engine/userguide/storagedriver/imagesandcontainers/#images-and-layers
- https://github.com/wagoodman/dive
- https://windsock.io/explaining-docker-image-ids/

__Tools__

jq
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo apt install -y jq
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  linux-headers-4.15.0-74 linux-headers-4.15.0-74-generic linux-image-4.15.0-74-generic linux-modules-4.15.0-74-generic
Use 'sudo apt autoremove' to remove them.
The following additional packages will be installed:
  libjq1 libonig4
The following NEW packages will be installed:
  jq libjq1 libonig4
0 upgraded, 3 newly installed, 0 to remove and 155 not upgraded.
Need to get 276 kB of archives.
After this operation, 930 kB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu bionic/universe amd64 libonig4 amd64 6.7.0-1 [119 kB]
Get:2 http://archive.ubuntu.com/ubuntu bionic/universe amd64 libjq1 amd64 1.5+dfsg-2 [111 kB]
Get:3 http://archive.ubuntu.com/ubuntu bionic/universe amd64 jq amd64 1.5+dfsg-2 [45.6 kB]
Fetched 276 kB in 3s (96.2 kB/s)
Selecting previously unselected package libonig4:amd64.
(Reading database ... 169731 files and directories currently installed.)
Preparing to unpack .../libonig4_6.7.0-1_amd64.deb ...
Unpacking libonig4:amd64 (6.7.0-1) ...
Selecting previously unselected package libjq1:amd64.
Preparing to unpack .../libjq1_1.5+dfsg-2_amd64.deb ...
Unpacking libjq1:amd64 (1.5+dfsg-2) ...
Selecting previously unselected package jq.
Preparing to unpack .../jq_1.5+dfsg-2_amd64.deb ...
Unpacking jq (1.5+dfsg-2) ...
Setting up libonig4:amd64 (6.7.0-1) ...
Setting up libjq1:amd64 (1.5+dfsg-2) ...
Processing triggers for libc-bin (2.27-3ubuntu1) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
Setting up jq (1.5+dfsg-2) ...
```

yq
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo pip3 install yq
The directory '/home/vagrant/.cache/pip/http' or its parent directory is not owned by the current user and the cache has been disabled. Please check the permissions and owner of that directory. If executing pip with sudo, you may want sudo's -H flag.
The directory '/home/vagrant/.cache/pip' or its parent directory is not owned by the current user and caching wheels has been disabled. check the permissions and owner of that directory. If executing pip with sudo, you may want sudo's -H flag.
Collecting yq
  Downloading https://files.pythonhosted.org/packages/c3/68/8994d78dcbb92cbfb70800bb6873f2694705516f4ae14ff08c05111fefeb/yq-2.10.0-py2.py3-none-any.whl
Collecting argcomplete>=1.8.1 (from yq)
  Downloading https://files.pythonhosted.org/packages/82/7d/455e149c28c320044cb763c23af375bd77d52baca041f611f5c2b4865cf4/argcomplete-1.11.1-py2.py3-none-any.whl
Requirement already satisfied: setuptools in /usr/lib/python3/dist-packages (from yq)
Collecting xmltodict>=0.11.0 (from yq)
  Downloading https://files.pythonhosted.org/packages/28/fd/30d5c1d3ac29ce229f6bdc40bbc20b28f716e8b363140c26eff19122d8a5/xmltodict-0.12.0-py2.py3-none-any.whl
Requirement already satisfied: PyYAML>=3.11 in /usr/lib/python3/dist-packages (from yq)
Collecting importlib-metadata<2,>=0.23; python_version == "3.6" (from argcomplete>=1.8.1->yq)
  Downloading https://files.pythonhosted.org/packages/8b/03/a00d504808808912751e64ccf414be53c29cad620e3de2421135fcae3025/importlib_metadata-1.5.0-py2.py3-none-any.whl
Collecting zipp>=0.5 (from importlib-metadata<2,>=0.23; python_version == "3.6"->argcomplete>=1.8.1->yq)
  Downloading https://files.pythonhosted.org/packages/6f/6d/a55f6e81ac213942b9a19cbc05b560c726c3e16f8fb17555f059c17d65f2/zipp-3.0.0-py3-none-any.whl
Installing collected packages: zipp, importlib-metadata, argcomplete, xmltodict, yq
Successfully installed argcomplete-1.11.1 importlib-metadata-1.5.0 xmltodict-0.12.0 yq-2.10.0 zipp-3.0.0
```

## Storage Driver

overlay2
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker info | grep 'Storage Driver'
Storage Driver: overlay2
WARNING: No swap limit support
```

## Docker image

### ImageDB

Repo
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo cat /var/lib/docker/image/overlay2/repositories.json | yq -y .
Repositories:
  gcr.io/google-containers/cadvisor:
    gcr.io/google-containers/cadvisor:latest: sha256:d24b7db72c9963ec20eb4b6e7b74fcfe0ea427418ffd9476b29ededfcca35981
    gcr.io/google-containers/cadvisor@sha256:46d4d730ef886aaece9e0a65a912564cab0303cf88718d82b3df84d3add6885c: sha256:d24b7db72c9963ec20eb4b6e7b74fcfe0ea427418ffd9476b29ededfcca35981
  k8s.gcr.io/coredns:
    k8s.gcr.io/coredns:1.6.5: sha256:70f311871ae12c14bd0e02028f249f933f925e4370744e4e35f706da773a8f61
    k8s.gcr.io/coredns@sha256:7ec975f167d815311a7136c32e70735f0d00b73781365df1befd46ed35bd4fe7: sha256:70f311871ae12c14bd0e02028f249f933f925e4370744e4e35f706da773a8f61
  k8s.gcr.io/etcd:
    k8s.gcr.io/etcd:3.4.3-0: sha256:303ce5db0e90dab1c5728ec70d21091201a23cdf8aeca70ab54943bbaaf0833f
    k8s.gcr.io/etcd@sha256:4afb99b4690b418ffc2ceb67e1a17376457e441c1f09ab55447f0aaf992fa646: sha256:303ce5db0e90dab1c5728ec70d21091201a23cdf8aeca70ab54943bbaaf0833f
  k8s.gcr.io/kube-apiserver:
    k8s.gcr.io/kube-apiserver:v1.17.2: sha256:41ef50a5f06a7fca2d93c6ea62b61ea894a13be3daf85d4c3e2d8e9840d76dae
    k8s.gcr.io/kube-apiserver@sha256:b22f7be5165a0022d282815067bda22f0282922f5ee65151e64cf3b54be09543: sha256:41ef50a5f06a7fca2d93c6ea62b61ea894a13be3daf85d4c3e2d8e9840d76dae
  k8s.gcr.io/kube-controller-manager:
    k8s.gcr.io/kube-controller-manager:v1.17.2: sha256:da5fd66c4068cf446a74d1e66273ec608fae896124e18e2287f52452cd2eea3f
    k8s.gcr.io/kube-controller-manager@sha256:146581abaa098200e65026303ba9e55805e34fd61a1e26c3a3358c4e57a18916: sha256:da5fd66c4068cf446a74d1e66273ec608fae896124e18e2287f52452cd2eea3f
  k8s.gcr.io/kube-proxy:
    k8s.gcr.io/kube-proxy:v1.17.2: sha256:cba2a99699bdf6829a4ec27388024f29ccbf7a758d862a311cb4215635d33e4b
    k8s.gcr.io/kube-proxy@sha256:9a2940bd7718e1b7060b96bb33b0af44ebb4e3de0a0f80b1e6f1622118e4f430: sha256:cba2a99699bdf6829a4ec27388024f29ccbf7a758d862a311cb4215635d33e4b
  k8s.gcr.io/kube-scheduler:
    k8s.gcr.io/kube-scheduler:v1.17.2: sha256:f52d4c527ef2f6919eaa8c1e9d2ed25134b2f8e6d47223f75ee2d2962c6c6fd9
    k8s.gcr.io/kube-scheduler@sha256:1530d01bdc70dfecd49f35454c51ce64665811ce736ad1c98cfc21f030e325c9: sha256:f52d4c527ef2f6919eaa8c1e9d2ed25134b2f8e6d47223f75ee2d2962c6c6fd9
  k8s.gcr.io/kubernetes-dashboard-amd64:
    k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1: sha256:f9aed6605b814b69e92dece6a50ed1e4e730144eb1cc971389dde9cb3820d124
    k8s.gcr.io/kubernetes-dashboard-amd64@sha256:0ae6b69432e78069c5ce2bcde0fe409c5c4d6f0f4d9cd50a17974fea38898747: sha256:f9aed6605b814b69e92dece6a50ed1e4e730144eb1cc971389dde9cb3820d124
  k8s.gcr.io/pause:
    k8s.gcr.io/pause:3.1: sha256:da86e6ba6ca197bf6bc5e9d900febd906b133eaa4750e6bed647b0fbe50ed43e
    k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea: sha256:da86e6ba6ca197bf6bc5e9d900febd906b133eaa4750e6bed647b0fbe50ed43e
  nginx:
    nginx:latest: sha256:231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145
    nginx@sha256:50cf965a6e08ec5784009d0fccb380fc479826b6e0e65684d9879170a9df8566: sha256:231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145
  quay.io/coreos/flannel:
    quay.io/coreos/flannel:v0.11.0-amd64: sha256:ff281650a721f46bbe2169292c91031c66411554739c88c861ba78475c1df894
    quay.io/coreos/flannel@sha256:7806805c93b20a168d0bbbd25c6a213f00ac58a511c47e8fa6409543528a204e: sha256:ff281650a721f46bbe2169292c91031c66411554739c88c861ba78475c1df894
  tangfeixiong/go-to-docker:
    tangfeixiong/go-to-docker:0.2: sha256:6d7a6961bf9c852a09fbda54d3e2f312739ec406b670ec23500da9233a38303f
    tangfeixiong/go-to-docker@sha256:7c6fd3fffb1ad6a4b6032cf795d6251fdf07cfbd21fe345bd31ee8abb4af31b7: sha256:6d7a6961bf9c852a09fbda54d3e2f312739ec406b670ec23500da9233a38303f
  tangfeixiong/gofileserver:
    tangfeixiong/gofileserver:latest: sha256:f07338e49481359f125f74f08665cb2e09722228b41834040e681034899c6159
    tangfeixiong/gofileserver@sha256:d7cb2c1256d8c67f93548975d2ecf03d658a5e8819c5cb58c5a7d90554b28479: sha256:f07338e49481359f125f74f08665cb2e09722228b41834040e681034899c6159
```

FileInfo
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/image/overlay2/imagedb/content/sha256/
231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145  da5fd66c4068cf446a74d1e66273ec608fae896124e18e2287f52452cd2eea3f
303ce5db0e90dab1c5728ec70d21091201a23cdf8aeca70ab54943bbaaf0833f  da86e6ba6ca197bf6bc5e9d900febd906b133eaa4750e6bed647b0fbe50ed43e
41ef50a5f06a7fca2d93c6ea62b61ea894a13be3daf85d4c3e2d8e9840d76dae  f07338e49481359f125f74f08665cb2e09722228b41834040e681034899c6159
6d7a6961bf9c852a09fbda54d3e2f312739ec406b670ec23500da9233a38303f  f52d4c527ef2f6919eaa8c1e9d2ed25134b2f8e6d47223f75ee2d2962c6c6fd9
70f311871ae12c14bd0e02028f249f933f925e4370744e4e35f706da773a8f61  f9aed6605b814b69e92dece6a50ed1e4e730144eb1cc971389dde9cb3820d124
cba2a99699bdf6829a4ec27388024f29ccbf7a758d862a311cb4215635d33e4b  ff281650a721f46bbe2169292c91031c66411554739c88c861ba78475c1df894
d24b7db72c9963ec20eb4b6e7b74fcfe0ea427418ffd9476b29ededfcca35981
```

Via CLI
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker images --no-trunc --format='table {{.ID}}\t{{.Repository}}:{{.Tag}}\t{{.Size}}'
IMAGE ID                                                                  REPOSITORY:TAG                                  SIZE
sha256:cba2a99699bdf6829a4ec27388024f29ccbf7a758d862a311cb4215635d33e4b   k8s.gcr.io/kube-proxy:v1.17.2                   116MB
sha256:da5fd66c4068cf446a74d1e66273ec608fae896124e18e2287f52452cd2eea3f   k8s.gcr.io/kube-controller-manager:v1.17.2      161MB
sha256:41ef50a5f06a7fca2d93c6ea62b61ea894a13be3daf85d4c3e2d8e9840d76dae   k8s.gcr.io/kube-apiserver:v1.17.2               171MB
sha256:f52d4c527ef2f6919eaa8c1e9d2ed25134b2f8e6d47223f75ee2d2962c6c6fd9   k8s.gcr.io/kube-scheduler:v1.17.2               94.4MB
sha256:231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145   nginx:latest                                    126MB
sha256:70f311871ae12c14bd0e02028f249f933f925e4370744e4e35f706da773a8f61   k8s.gcr.io/coredns:1.6.5                        41.6MB
sha256:303ce5db0e90dab1c5728ec70d21091201a23cdf8aeca70ab54943bbaaf0833f   k8s.gcr.io/etcd:3.4.3-0                         288MB
sha256:d24b7db72c9963ec20eb4b6e7b74fcfe0ea427418ffd9476b29ededfcca35981   gcr.io/google-containers/cadvisor:latest        185MB
sha256:ff281650a721f46bbe2169292c91031c66411554739c88c861ba78475c1df894   quay.io/coreos/flannel:v0.11.0-amd64            52.6MB
sha256:f9aed6605b814b69e92dece6a50ed1e4e730144eb1cc971389dde9cb3820d124   k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1   122MB
sha256:6d7a6961bf9c852a09fbda54d3e2f312739ec406b670ec23500da9233a38303f   tangfeixiong/go-to-docker:0.2                   37.3MB
sha256:da86e6ba6ca197bf6bc5e9d900febd906b133eaa4750e6bed647b0fbe50ed43e   k8s.gcr.io/pause:3.1                            742kB
sha256:f07338e49481359f125f74f08665cb2e09722228b41834040e681034899c6159   tangfeixiong/gofileserver:latest                12.6MB
```

__Inspect__

Example image is `nginx:lastest`

FileInfo
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo cat /var/lib/docker/image/overlay2/imagedb/content/sha256/231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145 | jq
{
  "architecture": "amd64",
  "config": {
    "Hostname": "",
    "Domainname": "",
    "User": "",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "ExposedPorts": {
      "80/tcp": {}
    },
    "Tty": false,
    "OpenStdin": false,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
      "NGINX_VERSION=1.17.6",
      "NJS_VERSION=0.3.7",
      "PKG_RELEASE=1~buster"
    ],
    "Cmd": [
      "nginx",
      "-g",
      "daemon off;"
    ],
    "ArgsEscaped": true,
    "Image": "sha256:f96d70a1d708239afa79b86f1e005c033864d22dabe94b466acba087d5bbc722",
    "Volumes": null,
    "WorkingDir": "",
    "Entrypoint": null,
    "OnBuild": null,
    "Labels": {
      "maintainer": "NGINX Docker Maintainers <docker-maint@nginx.com>"
    },
    "StopSignal": "SIGTERM"
  },
  "container": "806a0a78bcfee5212b2530e6f2a7e3f8eec5b51cc55d7a28935f5f8c8bd45826",
  "container_config": {
    "Hostname": "806a0a78bcfe",
    "Domainname": "",
    "User": "",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "ExposedPorts": {
      "80/tcp": {}
    },
    "Tty": false,
    "OpenStdin": false,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
      "NGINX_VERSION=1.17.6",
      "NJS_VERSION=0.3.7",
      "PKG_RELEASE=1~buster"
    ],
    "Cmd": [
      "/bin/sh",
      "-c",
      "#(nop) ",
      "CMD [\"nginx\" \"-g\" \"daemon off;\"]"
    ],
    "ArgsEscaped": true,
    "Image": "sha256:f96d70a1d708239afa79b86f1e005c033864d22dabe94b466acba087d5bbc722",
    "Volumes": null,
    "WorkingDir": "",
    "Entrypoint": null,
    "OnBuild": null,
    "Labels": {
      "maintainer": "NGINX Docker Maintainers <docker-maint@nginx.com>"
    },
    "StopSignal": "SIGTERM"
  },
  "created": "2019-11-23T01:12:31.219881158Z",
  "docker_version": "18.06.1-ce",
  "history": [
    {
      "created": "2019-11-22T14:55:09.912242636Z",
      "created_by": "/bin/sh -c #(nop) ADD file:bc8179c87c8dbb3d962bed1801f99e7c860ff03797cde6ad19b107d43b973ada in / "
    },
    {
      "created": "2019-11-22T14:55:10.253859615Z",
      "created_by": "/bin/sh -c #(nop)  CMD [\"bash\"]",
      "empty_layer": true
    },
    {
      "created": "2019-11-23T01:12:03.028307814Z",
      "created_by": "/bin/sh -c #(nop)  LABEL maintainer=NGINX Docker Maintainers <docker-maint@nginx.com>",
      "empty_layer": true
    },
    {
      "created": "2019-11-23T01:12:03.191740003Z",
      "created_by": "/bin/sh -c #(nop)  ENV NGINX_VERSION=1.17.6",
      "empty_layer": true
    },
    {
      "created": "2019-11-23T01:12:03.343142292Z",
      "created_by": "/bin/sh -c #(nop)  ENV NJS_VERSION=0.3.7",
      "empty_layer": true
    },
    {
      "created": "2019-11-23T01:12:03.499278285Z",
      "created_by": "/bin/sh -c #(nop)  ENV PKG_RELEASE=1~buster",
      "empty_layer": true
    },
    {
      "created": "2019-11-23T01:12:29.888960546Z",
      "created_by": "/bin/sh -c set -x     && addgroup --system --gid 101 nginx     && adduser --system --disabled-login --ingroup nginx --no-create-home --home /nonexistent --gecos \"nginx user\" --shell /bin/false --uid 101 nginx     && apt-get update     && apt-get install --no-install-recommends --no-install-suggests -y gnupg1 ca-certificates     &&     NGINX_GPGKEY=573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62;     found='';     for server in         ha.pool.sks-keyservers.net         hkp://keyserver.ubuntu.com:80         hkp://p80.pool.sks-keyservers.net:80         pgp.mit.edu     ; do         echo \"Fetching GPG key $NGINX_GPGKEY from $server\";         apt-key adv --keyserver \"$server\" --keyserver-options timeout=10 --recv-keys \"$NGINX_GPGKEY\" && found=yes && break;     done;     test -z \"$found\" && echo >&2 \"error: failed to fetch GPG key $NGINX_GPGKEY\" && exit 1;     apt-get remove --purge --auto-remove -y gnupg1 && rm -rf /var/lib/apt/lists/*     && dpkgArch=\"$(dpkg --print-architecture)\"     && nginxPackages=\"         nginx=${NGINX_VERSION}-${PKG_RELEASE}         nginx-module-xslt=${NGINX_VERSION}-${PKG_RELEASE}         nginx-module-geoip=${NGINX_VERSION}-${PKG_RELEASE}         nginx-module-image-filter=${NGINX_VERSION}-${PKG_RELEASE}         nginx-module-njs=${NGINX_VERSION}.${NJS_VERSION}-${PKG_RELEASE}     \"     && case \"$dpkgArch\" in         amd64|i386)             echo \"deb https://nginx.org/packages/mainline/debian/ buster nginx\" >> /etc/apt/sources.list.d/nginx.list             && apt-get update             ;;         *)             echo \"deb-src https://nginx.org/packages/mainline/debian/ buster nginx\" >> /etc/apt/sources.list.d/nginx.list                         && tempDir=\"$(mktemp -d)\"             && chmod 777 \"$tempDir\"                         && savedAptMark=\"$(apt-mark showmanual)\"                         && apt-get update             && apt-get build-dep -y $nginxPackages             && (                 cd \"$tempDir\"                 && DEB_BUILD_OPTIONS=\"nocheck parallel=$(nproc)\"                     apt-get source --compile $nginxPackages             )                         && apt-mark showmanual | xargs apt-mark auto > /dev/null             && { [ -z \"$savedAptMark\" ] || apt-mark manual $savedAptMark; }                         && ls -lAFh \"$tempDir\"             && ( cd \"$tempDir\" && dpkg-scanpackages . > Packages )             && grep '^Package: ' \"$tempDir/Packages\"             && echo \"deb [ trusted=yes ] file://$tempDir ./\" > /etc/apt/sources.list.d/temp.list             && apt-get -o Acquire::GzipIndexes=false update             ;;     esac         && apt-get install --no-install-recommends --no-install-suggests -y                         $nginxPackages                         gettext-base     && apt-get remove --purge --auto-remove -y ca-certificates && rm -rf /var/lib/apt/lists/* /etc/apt/sources.list.d/nginx.list         && if [ -n \"$tempDir\" ]; then         apt-get purge -y --auto-remove         && rm -rf \"$tempDir\" /etc/apt/sources.list.d/temp.list;     fi"
    },
    {
      "created": "2019-11-23T01:12:30.644605868Z",
      "created_by": "/bin/sh -c ln -sf /dev/stdout /var/log/nginx/access.log     && ln -sf /dev/stderr /var/log/nginx/error.log"
    },
    {
      "created": "2019-11-23T01:12:30.827542342Z",
      "created_by": "/bin/sh -c #(nop)  EXPOSE 80",
      "empty_layer": true
    },
    {
      "created": "2019-11-23T01:12:31.018531148Z",
      "created_by": "/bin/sh -c #(nop)  STOPSIGNAL SIGTERM",
      "empty_layer": true
    },
    {
      "created": "2019-11-23T01:12:31.219881158Z",
      "created_by": "/bin/sh -c #(nop)  CMD [\"nginx\" \"-g\" \"daemon off;\"]",
      "empty_layer": true
    }
  ],
  "os": "linux",
  "rootfs": {
    "type": "layers",
    "diff_ids": [
      "sha256:831c5620387fb9efec59fc82a42b948546c6be601e3ab34a87108ecf852aa15f",
      "sha256:5fb987d2e54d85820d95d6c31f3fe4cd95bf71fe6d9d9e4684082cb551b728b0",
      "sha256:4fc1aa8003a3d0d2481f10d17773869cbff12c1008df30e0bab8259086a0311c"
    ]
  }
}
```

__CLI__

image inspect
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker image inspect 231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145
[
    {
        "Id": "sha256:231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145",
        "RepoTags": [
            "nginx:latest"
        ],
        "RepoDigests": [
            "nginx@sha256:50cf965a6e08ec5784009d0fccb380fc479826b6e0e65684d9879170a9df8566"
        ],
        "Parent": "",
        "Comment": "",
        "Created": "2019-11-23T01:12:31.219881158Z",
        "Container": "806a0a78bcfee5212b2530e6f2a7e3f8eec5b51cc55d7a28935f5f8c8bd45826",
        "ContainerConfig": {
            "Hostname": "806a0a78bcfe",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "80/tcp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "NGINX_VERSION=1.17.6",
                "NJS_VERSION=0.3.7",
                "PKG_RELEASE=1~buster"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "#(nop) ",
                "CMD [\"nginx\" \"-g\" \"daemon off;\"]"
            ],
            "ArgsEscaped": true,
            "Image": "sha256:f96d70a1d708239afa79b86f1e005c033864d22dabe94b466acba087d5bbc722",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {
                "maintainer": "NGINX Docker Maintainers <docker-maint@nginx.com>"
            },
            "StopSignal": "SIGTERM"
        },
        "DockerVersion": "18.06.1-ce",
        "Author": "",
        "Config": {
            "Hostname": "",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "80/tcp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "NGINX_VERSION=1.17.6",
                "NJS_VERSION=0.3.7",
                "PKG_RELEASE=1~buster"
            ],
            "Cmd": [
                "nginx",
                "-g",
                "daemon off;"
            ],
            "ArgsEscaped": true,
            "Image": "sha256:f96d70a1d708239afa79b86f1e005c033864d22dabe94b466acba087d5bbc722",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {
                "maintainer": "NGINX Docker Maintainers <docker-maint@nginx.com>"
            },
            "StopSignal": "SIGTERM"
        },
        "Architecture": "amd64",
        "Os": "linux",
        "Size": 126323486,
        "VirtualSize": 126323486,
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/06f2a5aae7c610b7e566d62ed6c75ccf26af6439daaf710a6becb27d9a602873/diff:/var/lib/docker/overlay2/332335ea2f7bb50941bf495f1332f59413b5477708502c130f8e3d60201260f7/diff",
                "MergedDir": "/var/lib/docker/overlay2/d80f3d7b32fd6217f213ecfbc6c4826fbccc8a37fc8b5aff60faff50654670bd/merged",
                "UpperDir": "/var/lib/docker/overlay2/d80f3d7b32fd6217f213ecfbc6c4826fbccc8a37fc8b5aff60faff50654670bd/diff",
                "WorkDir": "/var/lib/docker/overlay2/d80f3d7b32fd6217f213ecfbc6c4826fbccc8a37fc8b5aff60faff50654670bd/work"
            },
            "Name": "overlay2"
        },
        "RootFS": {
            "Type": "layers",
            "Layers": [
                "sha256:831c5620387fb9efec59fc82a42b948546c6be601e3ab34a87108ecf852aa15f",
                "sha256:5fb987d2e54d85820d95d6c31f3fe4cd95bf71fe6d9d9e4684082cb551b728b0",
                "sha256:4fc1aa8003a3d0d2481f10d17773869cbff12c1008df30e0bab8259086a0311c"
            ]
        },
        "Metadata": {
            "LastTagTime": "0001-01-01T00:00:00Z"
        }
    }
]
```

and history
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker image history 231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145
IMAGE               CREATED             CREATED BY                                      SIZE                COMMENT
231d40e811cd        3 months ago        /bin/sh -c #(nop)  CMD ["nginx" "-g" "daemon…   0B                  
<missing>           3 months ago        /bin/sh -c #(nop)  STOPSIGNAL SIGTERM           0B                  
<missing>           3 months ago        /bin/sh -c #(nop)  EXPOSE 80                    0B                  
<missing>           3 months ago        /bin/sh -c ln -sf /dev/stdout /var/log/nginx…   22B                 
<missing>           3 months ago        /bin/sh -c set -x     && addgroup --system -…   57.1MB              
<missing>           3 months ago        /bin/sh -c #(nop)  ENV PKG_RELEASE=1~buster     0B                  
<missing>           3 months ago        /bin/sh -c #(nop)  ENV NJS_VERSION=0.3.7        0B                  
<missing>           3 months ago        /bin/sh -c #(nop)  ENV NGINX_VERSION=1.17.6     0B                  
<missing>           3 months ago        /bin/sh -c #(nop)  LABEL maintainer=NGINX Do…   0B                  
<missing>           3 months ago        /bin/sh -c #(nop)  CMD ["bash"]                 0B                  
<missing>           3 months ago        /bin/sh -c #(nop) ADD file:bc8179c87c8dbb3d9…   69.2MB              
```

__Dockerfile base__

https://github.com/nginxinc/docker-nginx/blob/5971de30c487356d5d2a2e1a79e02b2612f9a72f/mainline/buster/Dockerfile#L9-L93

### LayerDB

FileInfo
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/image/overlay2/layerdb/sha256/
0271b8eebde3fa9a6126b1f2335e170f902731ab4942f9f1914e77016540c7bb  9007f5987db353ec398a223bc5a135c5a9601798ba20a1abba537ea2f8ac765f
03901b4a2ea88eeaad62dbe59b072b28b6efa00491962b8741081c5df50c65e0  9071f958c5f953bec16358be105ff37bbb16577863ec6c09d3a6c094de9ca2d6
13e5b4c12c9f74dbab782c9593f9b08720b707473e62d31b3c8fa80490b26394  9cbe4314e608bdf2eafcd5e040d4c0db08083bf04986bb46e8f6e878b58b7179
225df95e717ceb672de0e45aa49f352eace21512240205972aca0fccc9612722  ac566949508d8e2e1bde46da8a4adf52e9656ec8264a791c6c55d01c8a067350
233ce78bbb2326584003e98666dfbfa8f989cdade6f83b8024c6142d3d90bf86  be613fd7386331bcad0cab3b7c2e2240096afa896f3b484a22f786d4c8a7e6a3
246a84598d6332bda90d6b0fda68021d1ee8403999c6fd1968af7525c3545452  c64dfa1d370fcf337a3f234e3d5f4f8b756d33b10f3bf701d0be5b571875c064
43822e2d958b024ddc1a87c6381229afeeb5b46e03204c7e74bbe7f320515010  c74a8da3cb5c88b86c67b17cd2e59d33d5fa2ac9b73d869744ece743404cb427
45ef3201fb19a421287beca9b361551dc62e4dd301965179ef1e568b790dd0dd  d19ca28a3d57fefb960b21f4c016a3117ec71aac037c909154e2fe57842fe6cf
561eaf129957ad575ddb973e7455175c1b74ed4bbaa925ff79080483c621087e  dc8adf8fa0fc82a56c32efac9d0da5f84153888317c88ab55123d9e71777bc62
5fb753ca965362ff62bc63dbc4f25ddc0af00184fc13f06ad550f67222462685  deffa7fa4d8c97d544a59568781dd0c6abdd0370c48ee4c650911072a1d855e6
6111c1cdaab1b6e4a2513f4bd48e07af0a06917ccfd7530b30986724164a7ee9  e17133b79956ad6f69ae7f775badd1c11bad2fc64f0529cab863b9d12fbaa5c4
64338652c36cea4908405858fd3d1d79615d801cd3e5b9834e68b8db22d752da  e94eccf67428940dbe13ec51739aa1919bfcb48b5c38cb783d95a4ef49e953c5
77fcff986d3b13762e4777046b9210a109fda20cb261bd3bbe5d7161d4e73c8e  eece166e3cf274a197be2a12d6c1111c5d1410d1e6c4cadcde37c70c1484ba41
7bff100f35cb359a368537bb07829b055fe8e0b1cb01085a3a628ae9c187c7b8  fbdfe08b001c6861c50073c98ed175d54e2d6440df7b797e52be97df0065098c
7c24a2c8ffdeb003469a74e92f64a2cebf76f64383df81717654406f62ea654e  fc4976bd934b81a21629d4eb3545d59ce82b5d101823f93a423c748708de1e92
831c5620387fb9efec59fc82a42b948546c6be601e3ab34a87108ecf852aa15f  fe9a8b4f1dccd77105b8423a26536ff756f1ee99fdcae6893eb737ae1c527c7a
8b8361f883ee5319753e4cbfbaa07bd3a75f25eef333c72f949894a6e66b6e1e
```

Cont.d __nginx:latest__ image, its inspection contains:
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker image inspect 231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145 | jq .[0].GraphDriver
{
  "Data": {
    "LowerDir": "/var/lib/docker/overlay2/06f2a5aae7c610b7e566d62ed6c75ccf26af6439daaf710a6becb27d9a602873/diff:/var/lib/docker/overlay2/332335ea2f7bb50941bf495f1332f59413b5477708502c130f8e3d60201260f7/diff",
    "MergedDir": "/var/lib/docker/overlay2/d80f3d7b32fd6217f213ecfbc6c4826fbccc8a37fc8b5aff60faff50654670bd/merged",
    "UpperDir": "/var/lib/docker/overlay2/d80f3d7b32fd6217f213ecfbc6c4826fbccc8a37fc8b5aff60faff50654670bd/diff",
    "WorkDir": "/var/lib/docker/overlay2/d80f3d7b32fd6217f213ecfbc6c4826fbccc8a37fc8b5aff60faff50654670bd/work"
  },
  "Name": "overlay2"
}
```

and
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker image inspect 231d40e811cd970168fb0c4770f2161aa30b9ba6fe8e68527504df69643aa145 | jq .[0].RootFS
{
  "Type": "layers",
  "Layers": [
    "sha256:831c5620387fb9efec59fc82a42b948546c6be601e3ab34a87108ecf852aa15f",
    "sha256:5fb987d2e54d85820d95d6c31f3fe4cd95bf71fe6d9d9e4684082cb551b728b0",
    "sha256:4fc1aa8003a3d0d2481f10d17773869cbff12c1008df30e0bab8259086a0311c"
  ]
}
```

RootFS layer in layerdb and distribution (AKA dockerhub)
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo find /var/lib/docker/image -name 831c5620387fb9efec59fc82a42b948546c6be601e3ab34a87108ecf852aa15f
/var/lib/docker/image/overlay2/layerdb/sha256/831c5620387fb9efec59fc82a42b948546c6be601e3ab34a87108ecf852aa15f
/var/lib/docker/image/overlay2/distribution/v2metadata-by-diffid/sha256/831c5620387fb9efec59fc82a42b948546c6be601e3ab34a87108ecf852aa15f
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo find /var/lib/docker/image -name 5fb987d2e54d85820d95d6c31f3fe4cd95bf71fe6d9d9e4684082cb551b728b0
/var/lib/docker/image/overlay2/distribution/v2metadata-by-diffid/sha256/5fb987d2e54d85820d95d6c31f3fe4cd95bf71fe6d9d9e4684082cb551b728b0
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo find /var/lib/docker/image -name 4fc1aa8003a3d0d2481f10d17773869cbff12c1008df30e0bab8259086a0311c
/var/lib/docker/image/overlay2/distribution/v2metadata-by-diffid/sha256/4fc1aa8003a3d0d2481f10d17773869cbff12c1008df30e0bab8259086a0311c
```

### Overlay2 Store

__base image layer__

id
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo cat /var/lib/docker/image/overlay2/layerdb/sha256/831c5620387fb9efec59fc82a42b948546c6be601e3ab34a87108ecf852aa15f/cache-id
332335ea2f7bb50941bf495f1332f59413b5477708502c130f8e3d60201260f7
```

store
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo find /var/lib/docker/overlay2 -name 332335ea2f7bb50941bf495f1332f59413b5477708502c130f8e3d60201260f7
/var/lib/docker/overlay2/332335ea2f7bb50941bf495f1332f59413b5477708502c130f8e3d60201260f7
```

File System
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/overlay2/332335ea2f7bb50941bf495f1332f59413b5477708502c130f8e3d60201260f7/diff/
bin  boot  dev	etc  home  lib	lib64  media  mnt  opt	proc  root  run  sbin  srv  sys  tmp  usr  var
```

__nginx layers__

- bin layer in https://github.com/nginxinc/docker-nginx/blob/5971de30c487356d5d2a2e1a79e02b2612f9a72f/mainline/buster/Dockerfile#L9-L93
- log layer in https://github.com/nginxinc/docker-nginx/blob/5971de30c487356d5d2a2e1a79e02b2612f9a72f/mainline/buster/Dockerfile#L95-L97

Find bin layer store
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo grep -rl '/var/lib/docker/image/overlay2/layerdb' -e '5fb987d2e54d85820d95d6c31f3fe4cd95bf71fe6d9d9e4684082cb551b728b0'
/var/lib/docker/image/overlay2/layerdb/sha256/77fcff986d3b13762e4777046b9210a109fda20cb261bd3bbe5d7161d4e73c8e/diff
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo cat /var/lib/docker/image/overlay2/layerdb/sha256/77fcff986d3b13762e4777046b9210a109fda20cb261bd3bbe5d7161d4e73c8e/cache-id
06f2a5aae7c610b7e566d62ed6c75ccf26af6439daaf710a6becb27d9a602873
```

File System
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls -l /var/lib/docker/overlay2/06f2a5aae7c610b7e566d62ed6c75ccf26af6439daaf710a6becb27d9a602873/diff/etc/nginx
total 40
drwxr-xr-x 2 root root 4096 Nov 23 01:12 conf.d
-rw-r--r-- 1 root root 1007 Nov 19 12:50 fastcgi_params
-rw-r--r-- 1 root root 2837 Nov 19 12:50 koi-utf
-rw-r--r-- 1 root root 2223 Nov 19 12:50 koi-win
-rw-r--r-- 1 root root 5231 Nov 19 12:50 mime.types
lrwxrwxrwx 1 root root   22 Nov 19 12:50 modules -> /usr/lib/nginx/modules
-rw-r--r-- 1 root root  643 Nov 19 12:50 nginx.conf
-rw-r--r-- 1 root root  636 Nov 19 12:50 scgi_params
-rw-r--r-- 1 root root  664 Nov 19 12:50 uwsgi_params
-rw-r--r-- 1 root root 3610 Nov 19 12:50 win-utf
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls -l /var/lib/docker/overlay2/06f2a5aae7c610b7e566d62ed6c75ccf26af6439daaf710a6becb27d9a602873/diff/usr/sbin/
total 2724
-rwxr-xr-x 1 root root 1330984 Nov 19 12:50 nginx
-rwxr-xr-x 1 root root 1457960 Nov 19 12:50 nginx-debug
```

Find log layer store
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo grep -rl '/var/lib/docker/image/overlay2/layerdb' -e '4fc1aa8003a3d0d2481f10d17773869cbff12c1008df30e0bab8259086a0311c'
/var/lib/docker/image/overlay2/layerdb/sha256/dc8adf8fa0fc82a56c32efac9d0da5f84153888317c88ab55123d9e71777bc62/diff
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo cat /var/lib/docker/image/overlay2/layerdb/sha256/dc8adf8fa0fc82a56c32efac9d0da5f84153888317c88ab55123d9e71777bc62/cache-id
d80f3d7b32fd6217f213ecfbc6c4826fbccc8a37fc8b5aff60faff50654670bd
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls -l /var/lib/docker/overlay2/d80f3d7b32fd6217f213ecfbc6c4826fbccc8a37fc8b5aff60faff50654670bd/diff/var/log/nginx
total 0
lrwxrwxrwx 1 root root 11 Nov 23 01:12 access.log -> /dev/stdout
lrwxrwxrwx 1 root root 11 Nov 23 01:12 error.log -> /dev/stderr
```

## Docker Container

run nginx
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker run -d -p 80 --name=nginx-test nginx:latest
62ec5dfb81e3455eb418c88e0b4235d861608556b0a493d9a21ea675baeb3d6f
```

Port DNat
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker port nginx-test
80/tcp -> 0.0.0.0:32768
```

curl
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ curl http://127.0.0.1:32768
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

id
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker container inspect nginx-test --format='{{.Id}}'
62ec5dfb81e3455eb418c88e0b4235d861608556b0a493d9a21ea675baeb3d6f
```

## Container Layers

Store
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker container inspect nginx-test | jq .[0].GraphDriver
{
  "Data": {
    "LowerDir": "/var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c-init/diff:/var/lib/docker/overlay2/d80f3d7b32fd6217f213ecfbc6c4826fbccc8a37fc8b5aff60faff50654670bd/diff:/var/lib/docker/overlay2/06f2a5aae7c610b7e566d62ed6c75ccf26af6439daaf710a6becb27d9a602873/diff:/var/lib/docker/overlay2/332335ea2f7bb50941bf495f1332f59413b5477708502c130f8e3d60201260f7/diff",
    "MergedDir": "/var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c/merged",
    "UpperDir": "/var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c/diff",
    "WorkDir": "/var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c/work"
  },
  "Name": "overlay2"
}
```

runtime additons
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls -l /var/lib/docker/containers | grep 62ec5dfb81e3455eb418c88e0b4235d861608556b0a493d9a21ea675baeb3d6f
drwx------ 4 root root 4096 Feb 22 02:39 62ec5dfb81e3455eb418c88e0b4235d861608556b0a493d9a21ea675baeb3d6f
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls -l /var/lib/docker/containers/62ec5dfb81e3455eb418c88e0b4235d861608556b0a493d9a21ea675baeb3d6f
total 36
-rw-r----- 1 root root  169 Feb 22 02:39 62ec5dfb81e3455eb418c88e0b4235d861608556b0a493d9a21ea675baeb3d6f-json.log
drwx------ 2 root root 4096 Feb 22 02:39 checkpoints
-rw------- 1 root root 2925 Feb 22 02:39 config.v2.json
-rw-r--r-- 1 root root 1476 Feb 22 02:39 hostconfig.json
-rw-r--r-- 1 root root   13 Feb 22 02:39 hostname
-rw-r--r-- 1 root root  174 Feb 22 02:39 hosts
drwx------ 3 root root 4096 Feb 22 02:39 mounts
-rw-r--r-- 1 root root  585 Feb 22 02:39 resolv.conf
-rw-r--r-- 1 root root   71 Feb 22 02:39 resolv.conf.hash
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c-init/diff
dev  etc
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c-init/diff/etc/
hostname  hosts  mtab  resolv.conf
```

Diff
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c/diff/
run  var
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c/diff/run/
nginx.pid
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c/diff/var/
cache
```

Merged
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c/merged/
bin  boot  dev	etc  home  lib	lib64  media  mnt  opt	proc  root  run  sbin  srv  sys  tmp  usr  var
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ls /var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c/merged/usr/sbin/nginx
/var/lib/docker/overlay2/6e808b8e118620a253cabfdaac59bd5a9faf45401216437143c49b4bc2dab77c/merged/usr/sbin/nginx
```

## Mounts and Volumes