# Development

## Java

### CLI

Because of JSP Project
```
fanhonglingdeMacBook-Pro:user-group-security fanhongling$ cd web-server/
```

Maven Command
```
fanhonglingdeMacBook-Pro:web-server fanhongling$ mvn package spring-boot:run -Dmaven.test.skip=true -Dspring.profiles.active=dev
```

## Docker

### Maven Build

Spring Boot
```
[vagrant@kubedev-172-17-4-59 web-server]$ mvn clean package spring-boot:repackage
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building Web Server 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.5:clean (default-clean) @ web-server ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/target
[INFO] 
[INFO] --- maven-resources-plugin:2.7:resources (default-resources) @ web-server ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 52 resources
[INFO] Copying 4 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:compile (default-compile) @ web-server ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 42 source files to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/target/classes
[INFO] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/src/main/java/com/blackbird/web/App.java: Some input files use or override a deprecated API.
[INFO] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/src/main/java/com/blackbird/web/App.java: Recompile with -Xlint:deprecation for details.
[INFO] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/src/main/java/com/blackbird/web/service/GameService.java: Some input files use unchecked or unsafe operations.
[INFO] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/src/main/java/com/blackbird/web/service/GameService.java: Recompile with -Xlint:unchecked for details.
[INFO] 
[INFO] --- maven-resources-plugin:2.7:testResources (default-testResources) @ web-server ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:testCompile (default-testCompile) @ web-server ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-surefire-plugin:2.22.0:test (default-test) @ web-server ---
[INFO] Tests are skipped.
[INFO] 
[INFO] --- maven-war-plugin:3.2.0:war (default-war) @ web-server ---
[INFO] Packaging webapp
[INFO] Assembling webapp [web-server] in [/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/target/web-server-0.0.1-SNAPSHOT]
[INFO] Processing war project
[INFO] Copying webapp resources [/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/src/main/webapp]
[INFO] Webapp assembled in [10239 msecs]
[INFO] Building war: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/web-server/target/web-server-0.0.1-SNAPSHOT.war
[INFO] 
[INFO] --- spring-boot-maven-plugin:2.0.2.RELEASE:repackage (all-in-one) @ web-server ---
[INFO] 
[INFO] --- spring-boot-maven-plugin:2.0.2.RELEASE:repackage (default-cli) @ web-server ---
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 01:27 min
[INFO] Finished at: 2018-06-27T02:38:48Z
[INFO] Final Memory: 33M/142M
[INFO] ------------------------------------------------------------------------
```

### Build Image

Destination WAR
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ tgt=$(basename $(ls web-server/target/web-server*.war))
```

```
[vagrant@kubedev-172-17-4-59 user-group-security]$ echo $tgt
web-server-0.0.1-SNAPSHOT.war
```

```
[vagrant@kubedev-172-17-4-59 user-group-security]$ docker build --force-rm --no-cache \
>     -t docker.io/tangfeixiong/ugs-web \
>     --build-arg jarTgt=target/${tgt} \
>     -f ./web-server/Dockerfile.centos7 ./web-server
Sending build context to Docker daemon 34.71 MB
Step 1/8 : FROM docker.io/centos:centos7
 ---> ff426288ea90
Step 2/8 : LABEL maintainer "tangfeixiong <tangfx128@gmail.com>" project "https://github.com/tangfeixiong/go-to-kubernetes" name "api" namespace "stackdocker0x2Eio" annotation '{"stackdocker.io/created-by":"n/a"}' tag "centos java springboot tomcat jsp shiro restTemplate"
 ---> Running in fc8090a57efe
 ---> 878c238595d8
Removing intermediate container fc8090a57efe
Step 3/8 : ARG jarTgt
 ---> Running in a9f66832adbd
 ---> 399102eda241
Removing intermediate container a9f66832adbd
Step 4/8 : ARG javaOpt
 ---> Running in 7539e07b91b4
 ---> 00cacac4a533
Removing intermediate container 7539e07b91b4
Step 5/8 : COPY ${jarTgt:-/target/web-server.war} /web-server.war
 ---> 78c6bf8e7fd0
Removing intermediate container 13802e622a2a
Step 6/8 : ENV JAVA_OPTIONS "${javaOpt:-'-Xms128m -Xmx512m -XX:PermSize=128m -XX:MaxPermSize=256m'}" APISERVER_ADDRESS "http://127.0.0.1:8090" SERVER_PORT "8080"
 ---> Running in cec0c05d59f8
 ---> e278b7f6353a
Removing intermediate container cec0c05d59f8
Step 7/8 : RUN set -x     && install_Pkgs="         tar         unzip         bc         which         lsof         java-1.8.0-openjdk-headless     "     && yum install -y $install_Pkgs     && yum clean all -y     && echo
 ---> Running in 8a2004575d2d
+ install_Pkgs='         tar         unzip         bc         which         lsof         java-1.8.0-openjdk-headless     '
+ yum install -y tar unzip bc which lsof java-1.8.0-openjdk-headless
Loaded plugins: fastestmirror, ovl
Determining fastest mirrors
 * base: mirrors.163.com
 * extras: mirrors.tuna.tsinghua.edu.cn
 * updates: mirrors.163.com
Resolving Dependencies
--> Running transaction check
---> Package bc.x86_64 0:1.06.95-13.el7 will be installed
---> Package java-1.8.0-openjdk-headless.x86_64 1:1.8.0.171-8.b10.el7_5 will be installed
--> Processing Dependency: tzdata-java >= 2015d for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: nss-softokn(x86-64) >= 3.36.0 for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: nss(x86-64) >= 3.36.0 for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: copy-jdk-configs >= 2.2 for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: lksctp-tools(x86-64) for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: libjpeg.so.62(LIBJPEG_6.2)(64bit) for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: jpackage-utils for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: libjpeg.so.62()(64bit) for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: libfreetype.so.6()(64bit) for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
---> Package lsof.x86_64 0:4.87-5.el7 will be installed
---> Package tar.x86_64 2:1.26-32.el7 will be updated
---> Package tar.x86_64 2:1.26-34.el7 will be an update
---> Package unzip.x86_64 0:6.0-19.el7 will be installed
---> Package which.x86_64 0:2.20-7.el7 will be installed
--> Running transaction check
---> Package copy-jdk-configs.noarch 0:3.3-10.el7_5 will be installed
---> Package freetype.x86_64 0:2.4.11-15.el7 will be installed
---> Package javapackages-tools.noarch 0:3.4.1-11.el7 will be installed
--> Processing Dependency: python-javapackages = 3.4.1-11.el7 for package: javapackages-tools-3.4.1-11.el7.noarch
--> Processing Dependency: libxslt for package: javapackages-tools-3.4.1-11.el7.noarch
---> Package libjpeg-turbo.x86_64 0:1.2.90-5.el7 will be installed
---> Package lksctp-tools.x86_64 0:1.0.17-2.el7 will be installed
---> Package nss.x86_64 0:3.28.4-15.el7_4 will be updated
--> Processing Dependency: nss = 3.28.4-15.el7_4 for package: nss-sysinit-3.28.4-15.el7_4.x86_64
--> Processing Dependency: nss(x86-64) = 3.28.4-15.el7_4 for package: nss-tools-3.28.4-15.el7_4.x86_64
---> Package nss.x86_64 0:3.36.0-5.el7_5 will be an update
--> Processing Dependency: nss-util >= 3.36.0-1 for package: nss-3.36.0-5.el7_5.x86_64
--> Processing Dependency: nspr >= 4.19.0 for package: nss-3.36.0-5.el7_5.x86_64
--> Processing Dependency: libnssutil3.so(NSSUTIL_3.31)(64bit) for package: nss-3.36.0-5.el7_5.x86_64
---> Package nss-softokn.x86_64 0:3.28.3-8.el7_4 will be updated
---> Package nss-softokn.x86_64 0:3.36.0-5.el7_5 will be an update
--> Processing Dependency: nss-softokn-freebl(x86-64) >= 3.36.0-5.el7_5 for package: nss-softokn-3.36.0-5.el7_5.x86_64
---> Package tzdata-java.noarch 0:2018e-3.el7 will be installed
--> Running transaction check
---> Package libxslt.x86_64 0:1.1.28-5.el7 will be installed
---> Package nspr.x86_64 0:4.13.1-1.0.el7_3 will be updated
---> Package nspr.x86_64 0:4.19.0-1.el7_5 will be an update
---> Package nss-softokn-freebl.x86_64 0:3.28.3-8.el7_4 will be updated
---> Package nss-softokn-freebl.x86_64 0:3.36.0-5.el7_5 will be an update
---> Package nss-sysinit.x86_64 0:3.28.4-15.el7_4 will be updated
---> Package nss-sysinit.x86_64 0:3.36.0-5.el7_5 will be an update
---> Package nss-tools.x86_64 0:3.28.4-15.el7_4 will be updated
---> Package nss-tools.x86_64 0:3.36.0-5.el7_5 will be an update
---> Package nss-util.x86_64 0:3.28.4-3.el7 will be updated
---> Package nss-util.x86_64 0:3.36.0-1.el7_5 will be an update
---> Package python-javapackages.noarch 0:3.4.1-11.el7 will be installed
--> Processing Dependency: python-lxml for package: python-javapackages-3.4.1-11.el7.noarch
--> Running transaction check
---> Package python-lxml.x86_64 0:3.2.1-4.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package                      Arch    Version                    Repository
                                                                           Size
================================================================================
Installing:
 bc                           x86_64  1.06.95-13.el7             base     115 k
 java-1.8.0-openjdk-headless  x86_64  1:1.8.0.171-8.b10.el7_5    updates   32 M
 lsof                         x86_64  4.87-5.el7                 base     331 k
 unzip                        x86_64  6.0-19.el7                 base     170 k
 which                        x86_64  2.20-7.el7                 base      41 k
Updating:
 tar                          x86_64  2:1.26-34.el7              base     845 k
Installing for dependencies:
 copy-jdk-configs             noarch  3.3-10.el7_5               updates   21 k
 freetype                     x86_64  2.4.11-15.el7              base     392 k
 javapackages-tools           noarch  3.4.1-11.el7               base      73 k
 libjpeg-turbo                x86_64  1.2.90-5.el7               base     134 k
 libxslt                      x86_64  1.1.28-5.el7               base     242 k
 lksctp-tools                 x86_64  1.0.17-2.el7               base      88 k
 python-javapackages          noarch  3.4.1-11.el7               base      31 k
 python-lxml                  x86_64  3.2.1-4.el7                base     758 k
 tzdata-java                  noarch  2018e-3.el7                updates  185 k
Updating for dependencies:
 nspr                         x86_64  4.19.0-1.el7_5             updates  127 k
 nss                          x86_64  3.36.0-5.el7_5             updates  835 k
 nss-softokn                  x86_64  3.36.0-5.el7_5             updates  315 k
 nss-softokn-freebl           x86_64  3.36.0-5.el7_5             updates  222 k
 nss-sysinit                  x86_64  3.36.0-5.el7_5             updates   62 k
 nss-tools                    x86_64  3.36.0-5.el7_5             updates  514 k
 nss-util                     x86_64  3.36.0-1.el7_5             updates   78 k

Transaction Summary
================================================================================
Install  5 Packages (+9 Dependent packages)
Upgrade  1 Package  (+7 Dependent packages)

Total download size: 37 M
Downloading packages:
Delta RPMs disabled because /usr/bin/applydeltarpm not installed.
warning: /var/cache/yum/x86_64/7/base/packages/freetype-2.4.11-15.el7.x86_64.rpm: Header V3 RSA/SHA256 Signature, key ID f4a80eb5: NOKEY
Public key for freetype-2.4.11-15.el7.x86_64.rpm is not installed
Public key for copy-jdk-configs-3.3-10.el7_5.noarch.rpm is not installed
--------------------------------------------------------------------------------
Total                                              7.8 MB/s |  37 MB  00:04     
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Importing GPG key 0xF4A80EB5:
 Userid     : "CentOS-7 Key (CentOS 7 Official Signing Key) <security@centos.org>"
 Fingerprint: 6341 ab27 53d7 8a78 a7c2 7bb1 24c6 a8a7 f4a8 0eb5
 Package    : centos-release-7-4.1708.el7.centos.x86_64 (@CentOS)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Updating   : nspr-4.19.0-1.el7_5.x86_64                                  1/30 
  Updating   : nss-util-3.36.0-1.el7_5.x86_64                              2/30 
  Installing : libxslt-1.1.28-5.el7.x86_64                                 3/30 
  Installing : python-lxml-3.2.1-4.el7.x86_64                              4/30 
  Installing : python-javapackages-3.4.1-11.el7.noarch                     5/30 
  Installing : javapackages-tools-3.4.1-11.el7.noarch                      6/30 
  Updating   : nss-softokn-freebl-3.36.0-5.el7_5.x86_64                    7/30 
  Updating   : nss-softokn-3.36.0-5.el7_5.x86_64                           8/30 
  Updating   : nss-3.36.0-5.el7_5.x86_64                                   9/30 
  Updating   : nss-sysinit-3.36.0-5.el7_5.x86_64                          10/30 
  Installing : tzdata-java-2018e-3.el7.noarch                             11/30 
  Installing : lksctp-tools-1.0.17-2.el7.x86_64                           12/30 
  Installing : copy-jdk-configs-3.3-10.el7_5.noarch                       13/30 
  Installing : freetype-2.4.11-15.el7.x86_64                              14/30 
  Installing : libjpeg-turbo-1.2.90-5.el7.x86_64                          15/30 
  Installing : 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_   16/30 
  Updating   : nss-tools-3.36.0-5.el7_5.x86_64                            17/30 
  Installing : unzip-6.0-19.el7.x86_64                                    18/30 
  Updating   : 2:tar-1.26-34.el7.x86_64                                   19/30 
  Installing : lsof-4.87-5.el7.x86_64                                     20/30 
  Installing : which-2.20-7.el7.x86_64                                    21/30 
install-info: No such file or directory for /usr/share/info/which.info.gz
  Installing : bc-1.06.95-13.el7.x86_64                                   22/30 
  Cleanup    : nss-tools-3.28.4-15.el7_4.x86_64                           23/30 
  Cleanup    : nss-sysinit-3.28.4-15.el7_4.x86_64                         24/30 
  Cleanup    : nss-3.28.4-15.el7_4.x86_64                                 25/30 
  Cleanup    : nss-softokn-3.28.3-8.el7_4.x86_64                          26/30 
  Cleanup    : nss-util-3.28.4-3.el7.x86_64                               27/30 
  Cleanup    : nss-softokn-freebl-3.28.3-8.el7_4.x86_64                   28/30 
  Cleanup    : nspr-4.13.1-1.0.el7_3.x86_64                               29/30 
  Cleanup    : 2:tar-1.26-32.el7.x86_64                                   30/30 
  Verifying  : 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_    1/30 
  Verifying  : python-javapackages-3.4.1-11.el7.noarch                     2/30 
  Verifying  : libjpeg-turbo-1.2.90-5.el7.x86_64                           3/30 
  Verifying  : nss-tools-3.36.0-5.el7_5.x86_64                             4/30 
  Verifying  : python-lxml-3.2.1-4.el7.x86_64                              5/30 
  Verifying  : freetype-2.4.11-15.el7.x86_64                               6/30 
  Verifying  : bc-1.06.95-13.el7.x86_64                                    7/30 
  Verifying  : nss-util-3.36.0-1.el7_5.x86_64                              8/30 
  Verifying  : copy-jdk-configs-3.3-10.el7_5.noarch                        9/30 
  Verifying  : which-2.20-7.el7.x86_64                                    10/30 
  Verifying  : lsof-4.87-5.el7.x86_64                                     11/30 
  Verifying  : nss-3.36.0-5.el7_5.x86_64                                  12/30 
  Verifying  : lksctp-tools-1.0.17-2.el7.x86_64                           13/30 
  Verifying  : libxslt-1.1.28-5.el7.x86_64                                14/30 
  Verifying  : javapackages-tools-3.4.1-11.el7.noarch                     15/30 
  Verifying  : nss-softokn-freebl-3.36.0-5.el7_5.x86_64                   16/30 
  Verifying  : 2:tar-1.26-34.el7.x86_64                                   17/30 
  Verifying  : nspr-4.19.0-1.el7_5.x86_64                                 18/30 
  Verifying  : tzdata-java-2018e-3.el7.noarch                             19/30 
  Verifying  : unzip-6.0-19.el7.x86_64                                    20/30 
  Verifying  : nss-softokn-3.36.0-5.el7_5.x86_64                          21/30 
  Verifying  : nss-sysinit-3.36.0-5.el7_5.x86_64                          22/30 
  Verifying  : nspr-4.13.1-1.0.el7_3.x86_64                               23/30 
  Verifying  : nss-util-3.28.4-3.el7.x86_64                               24/30 
  Verifying  : nss-softokn-3.28.3-8.el7_4.x86_64                          25/30 
  Verifying  : 2:tar-1.26-32.el7.x86_64                                   26/30 
  Verifying  : nss-sysinit-3.28.4-15.el7_4.x86_64                         27/30 
  Verifying  : nss-3.28.4-15.el7_4.x86_64                                 28/30 
  Verifying  : nss-tools-3.28.4-15.el7_4.x86_64                           29/30 
  Verifying  : nss-softokn-freebl-3.28.3-8.el7_4.x86_64                   30/30 

Installed:
  bc.x86_64 0:1.06.95-13.el7                                                    
  java-1.8.0-openjdk-headless.x86_64 1:1.8.0.171-8.b10.el7_5                    
  lsof.x86_64 0:4.87-5.el7                                                      
  unzip.x86_64 0:6.0-19.el7                                                     
  which.x86_64 0:2.20-7.el7                                                     

Dependency Installed:
  copy-jdk-configs.noarch 0:3.3-10.el7_5    freetype.x86_64 0:2.4.11-15.el7    
  javapackages-tools.noarch 0:3.4.1-11.el7  libjpeg-turbo.x86_64 0:1.2.90-5.el7
  libxslt.x86_64 0:1.1.28-5.el7             lksctp-tools.x86_64 0:1.0.17-2.el7 
  python-javapackages.noarch 0:3.4.1-11.el7 python-lxml.x86_64 0:3.2.1-4.el7   
  tzdata-java.noarch 0:2018e-3.el7         

Updated:
  tar.x86_64 2:1.26-34.el7                                                      

Dependency Updated:
  nspr.x86_64 0:4.19.0-1.el7_5                                                  
  nss.x86_64 0:3.36.0-5.el7_5                                                   
  nss-softokn.x86_64 0:3.36.0-5.el7_5                                           
  nss-softokn-freebl.x86_64 0:3.36.0-5.el7_5                                    
  nss-sysinit.x86_64 0:3.36.0-5.el7_5                                           
  nss-tools.x86_64 0:3.36.0-5.el7_5                                             
  nss-util.x86_64 0:3.36.0-1.el7_5                                              

Complete!
+ yum clean all -y
Loaded plugins: fastestmirror, ovl
Cleaning repos: base extras updates
Cleaning up everything
Maybe you want: rm -rf /var/cache/yum, to also free up space taken by orphaned data from disabled or removed repos
Cleaning up list of fastest mirrors
+ echo

 ---> 9bba72a332b5
Removing intermediate container 8a2004575d2d
Step 8/8 : CMD java -Djava.security.egd=file:/dev/./urandom $JAVA_OPTIONS -jar /web-server.war
 ---> Running in 9b94ed1d6f49
 ---> 151acde5a147
Removing intermediate container 9b94ed1d6f49
Successfully built 151acde5a147
```

```
[vagrant@kubedev-172-17-4-59 user-group-security]$ docker images docker.io/tangfeixiong/ugs-web
REPOSITORY                       TAG                 IMAGE ID            CREATED              SIZE
docker.io/tangfeixiong/ugs-web   latest              151acde5a147        About a minute ago   405 MB
```

### Deploy Container

Networking
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ docker network ls -f name=user_organization_security
NETWORK ID          NAME                         DRIVER              SCOPE
b11aa47e055a        user_organization_security   bridge              local
```

Run with API-Server container provided
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ docker run -d --network=user_organization_security \
> --name ugs-web -p18080:8080 \
> -e APISERVER_ADDRESS=http://ugs-api:8090 \
> --restart=no \  
> docker.io/tangfeixiong/ugs-web
03c7248027a68608f19b1504723b330d9fa46634f1f837683b59e4b4fe63567d
```

Log
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ docker logs ugs-web
OpenJDK 64-Bit Server VM warning: ignoring option PermSize=128m; support was removed in 8.0
OpenJDK 64-Bit Server VM warning: ignoring option MaxPermSize=256m; support was removed in 8.0
log4j:WARN No appenders could be found for logger (org.springframework.web.context.support.StandardServletEnvironment).
log4j:WARN Please initialize the log4j system properly.
log4j:WARN See http://logging.apache.org/log4j/1.2/faq.html#noconfig for more info.
SLF4J: Class path contains multiple SLF4J bindings.
SLF4J: Found binding in [jar:file:/web-server.war!/WEB-INF/lib/logback-classic-1.2.3.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: Found binding in [jar:file:/web-server.war!/WEB-INF/lib/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: See http://www.slf4j.org/codes.html#multiple_bindings for an explanation.
SLF4J: Actual binding is of type [ch.qos.logback.classic.util.ContextSelectorStaticBinder]
  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::        (v2.0.2.RELEASE)
2018-06-27 02:42:36.262  INFO 1 --- [           main] o.apache.catalina.core.StandardService   : Starting service [Tomcat]
2018-06-27 02:42:36.273  INFO 1 --- [           main] org.apache.catalina.core.StandardEngine  : Starting Servlet Engine: Apache Tomcat/8.5.31
2018-06-27 02:42:36.322  INFO 1 --- [ost-startStop-1] o.a.catalina.core.AprLifecycleListener   : The APR based Apache Tomcat Native library which allows optimal performance in production environments was not found on the java.library.path: [/usr/java/packages/lib/amd64:/usr/lib64:/lib64:/lib:/usr/lib]
2018-06-27 02:42:38.248  INFO 1 --- [ost-startStop-1] org.apache.jasper.servlet.TldScanner     : At least one JAR was scanned for TLDs yet contained no TLDs. Enable debug logging for this logger for a complete list of JARs that were scanned but no TLDs were found in them. Skipping unneeded JARs during scanning can improve startup time and JSP compilation time.
2018-06-27 02:42:38.879  INFO 1 --- [ost-startStop-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring embedded WebApplicationContext
```

Open web browser, For example _http://172.17.4.59:18080/_
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ ip addr show eth1
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:31:d6:53 brd ff:ff:ff:ff:ff:ff
    inet 172.17.4.59/24 brd 172.17.4.255 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe31:d653/64 scope link 
       valid_lft forever preferred_lft forever
```