### Instruction

Maven Proerties
```
		<image.namespace>172.17.4.50:5000</image.namespace>
		<image.name>isc-flag</image.name>
		<image.tag>latest</image.tag>
```

Build with alternative repository and push
```
$ mvn clean compile package spring-boot:repackage docker:build docker:push -Dimage.namespace=192.168.1.209:5000 -Dimage.name=flag-server -Dimage.tag=0.0.1-SNAPSHOT
```

Build example
```
[vagrant@localhost flag]$ mvn clean compile package spring-boot:repackage docker:build
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Flag 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.6.1:clean (default-clean) @ Flag ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/flag/target
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ Flag ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 1 resource
[INFO] Copying 5 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ Flag ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 31 source files to /Users/fanhongling/Downloads/workspace/src/flag/target/classes
[WARNING] /Users/fanhongling/Downloads/workspace/src/flag/src/main/java/cn/com/isc/SwaggerConfig.java: /Users/fanhongling/Downloads/workspace/src/flag/src/main/java/cn/com/isc/SwaggerConfig.java使用或覆盖了已过时的 API。
[WARNING] /Users/fanhongling/Downloads/workspace/src/flag/src/main/java/cn/com/isc/SwaggerConfig.java: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ Flag ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 1 resource
[INFO] Copying 5 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ Flag ---
[INFO] Nothing to compile - all classes are up to date
[INFO] 
[INFO] --- maven-resources-plugin:2.6:testResources (default-testResources) @ Flag ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/flag/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:testCompile (default-testCompile) @ Flag ---
[INFO] Nothing to compile - all classes are up to date
[INFO] 
[INFO] --- maven-surefire-plugin:2.18.1:test (default-test) @ Flag ---
[INFO] No tests to run.
[INFO] 
[INFO] --- maven-jar-plugin:2.6:jar (default-jar) @ Flag ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/flag/target/Flag-0.0.1-SNAPSHOT.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.5.3.RELEASE:repackage (default) @ Flag ---
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.5.3.RELEASE:repackage (default-cli) @ Flag ---
[INFO] 
[INFO] --- docker-maven-plugin:0.21.0:build (default-cli) @ Flag ---
[INFO] Copying files to /Users/fanhongling/Downloads/workspace/src/flag/target/docker/172.17.4.50/5000/isc-flag/latest/build/maven
[INFO] Building tar: /Users/fanhongling/Downloads/workspace/src/flag/target/docker/172.17.4.50/5000/isc-flag/latest/tmp/docker-build.tar
[INFO] DOCKER> [172.17.4.50:5000/isc-flag:latest]: Created docker-build.tar in 11 seconds 
[INFO] DOCKER> [172.17.4.50:5000/isc-flag:latest]: Built image sha256:f34e7
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 04:44 min
[INFO] Finished at: 2017-08-22T18:01:10+00:00
[INFO] Final Memory: 39M/332M
[INFO] ------------------------------------------------------------------------
```

Image
```
[vagrant@localhost flag]$ docker images 172.17.4.50:5000/isc-flag
REPOSITORY                   TAG                 IMAGE ID            CREATED              SIZE
172.17.4.50:5000/isc-flag   latest              f34e79b410dd        About a minute ago   385.5 MB
```
