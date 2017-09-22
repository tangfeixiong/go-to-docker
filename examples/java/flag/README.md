# Instruction

## Table of contents

Docker

Development

Reference

## Docker

Relative properties in `pom.xml`
```
	<properties>
    
		<image.namespace>docker.io</image.namespace>
		<image.name>tangfeixiong/refresh-cm</image.name>
		<image.tag>0.1</image.tag>
	</properties>
```

Build with alternative repository and push
```
$ mvn clean compile package spring-boot:repackage docker:build docker:push -Dimage.namespace=172.17.4.50:5000 -Dimage.name=awesome-rest-server -Dimage.tag=latest
```

Example build
```
[vagrant@bogon flag]$ mvn clean compile package spring-boot:repackage docker:build
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Refresh Controller Manager 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.6.1:clean (default-clean) @ refresh-cm ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/java/flag/target
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ refresh-cm ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 2 resources
[INFO] Copying 6 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ refresh-cm ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 36 source files to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/java/flag/target/classes
[WARNING] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/java/flag/src/main/java/io/stackdocker/iscc/flagserver/SwaggerConfig.java: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/java/flag/src/main/java/io/stackdocker/iscc/flagserver/SwaggerConfig.java使用或覆盖了已过时的 API。
[WARNING] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/java/flag/src/main/java/io/stackdocker/iscc/flagserver/SwaggerConfig.java: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ refresh-cm ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 2 resources
[INFO] Copying 6 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ refresh-cm ---
[INFO] Nothing to compile - all classes are up to date
[INFO] 
[INFO] --- maven-resources-plugin:2.6:testResources (default-testResources) @ refresh-cm ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/java/flag/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:testCompile (default-testCompile) @ refresh-cm ---
[INFO] Nothing to compile - all classes are up to date
[INFO] 
[INFO] --- maven-surefire-plugin:2.18.1:test (default-test) @ refresh-cm ---
[INFO] No tests to run.
[INFO] 
[INFO] --- maven-jar-plugin:2.6:jar (default-jar) @ refresh-cm ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/java/flag/target/refresh-cm-0.0.1-SNAPSHOT.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.5.3.RELEASE:repackage (default) @ refresh-cm ---
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.5.3.RELEASE:repackage (default-cli) @ refresh-cm ---
[INFO] 
[INFO] --- docker-maven-plugin:0.21.0:build (default-cli) @ refresh-cm ---
[INFO] Copying files to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/java/flag/target/docker/docker.io/tangfeixiong/refresh-cm/0.1/build/maven
[INFO] Building tar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/java/flag/target/docker/docker.io/tangfeixiong/refresh-cm/0.1/tmp/docker-build.tar
[INFO] DOCKER> [docker.io/tangfeixiong/refresh-cm:0.1]: Created docker-build.tar in 8 seconds 
[INFO] DOCKER> [docker.io/tangfeixiong/refresh-cm:0.1]: Built image sha256:6c2eb
[INFO] DOCKER> [docker.io/tangfeixiong/refresh-cm:0.1]: Removed old image sha256:e6fcb
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 03:44 min
[INFO] Finished at: 2017-09-19T09:44:11+00:00
[INFO] Final Memory: 41M/427M
[INFO] ------------------------------------------------------------------------
```

Image
```
[vagrant@bogon flag]$ docker images tangfeixiong/refresh-cm
REPOSITORY                          TAG                 IMAGE ID            CREATED             SIZE
docker.io/tangfeixiong/refresh-cm   0.1                 6c2ebbb5863f        6 minutes ago       389.6 MB
```

Run
```
[vagrant@localhost flag]$ docker run --rm -p 8082:8082 -v /tmp/mnt-home:/tmp/mnt-home -e SPRING_DATASOURCE_URL=jdbc:mysql://172.17.0.8:3306/testdb -e SPRING_DATASOURCE_USERNAME=testuser -e SPRING_DATASOURCE_PASSWORD=testpassword docker.io/tangfeixiong/refresh-cm:0.1
OpenJDK 64-Bit Server VM warning: ignoring option PermSize=128m; support was removed in 8.0
OpenJDK 64-Bit Server VM warning: ignoring option MaxPermSize=256m; support was removed in 8.0

  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::        (v1.5.3.RELEASE)

2017-09-19 10:04:18.330  INFO 1 --- [           main] cn.com.isc.App                           : Starting App v0.0.1-SNAPSHOT on 0b01f47d52fe with PID 1 (/refresh-cm.jar started by root in /)
2017-09-19 10:04:18.355  INFO 1 --- [           main] cn.com.isc.App                           : No active profile set, falling back to default profiles: default
2017-09-19 10:04:18.600  INFO 1 --- [           main] ationConfigEmbeddedWebApplicationContext : Refreshing org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@5387f9e0: startup date [Tue Sep 19 10:04:18 UTC 2017]; root of context hierarchy
2017-09-19 10:04:22.717  INFO 1 --- [           main] trationDelegate$BeanPostProcessorChecker : Bean 'org.springframework.transaction.annotation.ProxyTransactionManagementConfiguration' of type [org.springframework.transaction.annotation.ProxyTransactionManagementConfiguration$$EnhancerBySpringCGLIB$$3469a8f4] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)
2017-09-19 10:04:23.373  INFO 1 --- [           main] s.b.c.e.t.TomcatEmbeddedServletContainer : Tomcat initialized with port(s): 8082 (http)
2017-09-19 10:04:23.417  INFO 1 --- [           main] o.apache.catalina.core.StandardService   : Starting service Tomcat
2017-09-19 10:04:23.422  INFO 1 --- [           main] org.apache.catalina.core.StandardEngine  : Starting Servlet Engine: Apache Tomcat/8.5.14
### sinppets ###
2017-09-19 10:04:31.647  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/v1/refresh-deletion],methods=[POST]}" onto public io.stackdocker.iscc.flagserver.api.RefreshReqResp io.stackdocker.iscc.flagserver.controller.RefreshController.deleteRefresh(io.stackdocker.iscc.flagserver.api.RefreshReqResp)
2017-09-19 10:04:31.648  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/v1/refresh-updation],methods=[POST]}" onto public io.stackdocker.iscc.flagserver.api.RefreshReqResp io.stackdocker.iscc.flagserver.controller.RefreshController.updateRefresh(io.stackdocker.iscc.flagserver.api.RefreshReqResp)
2017-09-19 10:04:31.651  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/v1/refresh-creation],methods=[POST]}" onto public io.stackdocker.iscc.flagserver.api.RefreshReqResp io.stackdocker.iscc.flagserver.controller.RefreshController.createRefresh(io.stackdocker.iscc.flagserver.api.RefreshReqResp)
2017-09-19 10:04:31.653  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/v1/Validate],methods=[POST]}" onto public java.util.Map<java.lang.String, java.lang.String> io.stackdocker.iscc.flagserver.controller.RefreshController.get(cn.com.isc.entity.Flag)
2017-09-19 10:04:31.655  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/v1/getCount],methods=[POST]}" onto public java.util.Map<java.lang.String, java.lang.Object> io.stackdocker.iscc.flagserver.controller.RefreshController.get()
2017-09-19 10:04:31.657  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/v1/open],methods=[GET]}" onto public java.util.Map<java.lang.String, java.lang.Object> io.stackdocker.iscc.flagserver.controller.RefreshController.test()
2017-09-19 10:04:31.663  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/swagger-resources/configuration/ui]}" onto public org.springframework.http.ResponseEntity<springfox.documentation.swagger.web.UiConfiguration> springfox.documentation.swagger.web.ApiResourceController.uiConfiguration()
2017-09-19 10:04:31.666  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/swagger-resources/configuration/security]}" onto public org.springframework.http.ResponseEntity<springfox.documentation.swagger.web.SecurityConfiguration> springfox.documentation.swagger.web.ApiResourceController.securityConfiguration()
2017-09-19 10:04:31.668  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/swagger-resources]}" onto public org.springframework.http.ResponseEntity<java.util.List<springfox.documentation.swagger.web.SwaggerResource>> springfox.documentation.swagger.web.ApiResourceController.swaggerResources()
2017-09-19 10:04:31.673  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error],produces=[text/html]}" onto public org.springframework.web.servlet.ModelAndView org.springframework.boot.autoconfigure.web.BasicErrorController.errorHtml(javax.servlet.http.HttpServletRequest,javax.servlet.http.HttpServletResponse)
2017-09-19 10:04:31.676  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error]}" onto public org.springframework.http.ResponseEntity<java.util.Map<java.lang.String, java.lang.Object>> org.springframework.boot.autoconfigure.web.BasicErrorController.error(javax.servlet.http.HttpServletRequest)
2017-09-19 10:04:32.085  INFO 1 --- [           main] pertySourcedRequestMappingHandlerMapping : Mapped URL path [/v2/api-docs] onto method [public org.springframework.http.ResponseEntity<springfox.documentation.spring.web.json.Json> springfox.documentation.swagger2.web.Swagger2Controller.getDocumentation(java.lang.String,javax.servlet.http.HttpServletRequest)]
2017-09-19 10:04:32.390  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerAdapter : Looking for @ControllerAdvice: org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@5387f9e0: startup date [Tue Sep 19 10:04:18 UTC 2017]; root of context hierarchy
2017-09-19 10:04:32.573  INFO 1 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/webjars/**] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
### snippets ###
2017-09-19 10:04:37.439  INFO 1 --- [           main] cn.com.isc.App                           : Started App in 19.985 seconds (JVM running for 21.501)
###
Go to create refreshing: id:1,image_id:1,project_id:1,name:test
Schedule Key : test Value : container_id:1,config_id:1,team_id:1,name:flag,sub_path:demo1/
Runnable Task with test on thread ThreadPoolTaskScheduler1
get flag with: env=1 teamNo=1 round=1
###
Runnable Task with test on thread ThreadPoolTaskScheduler3
get flag with: env=1 teamNo=1 round=2
###
Runnable Task with test on thread ThreadPoolTaskScheduler5
Game over, bout : 10
```

And output
```
[vagrant@localhost flag]$ cat /tmp/mnt-home/test1/demo1/1
11BF33B2D7F8A37B59F6A128A86693FE[vagrant@localhost flag]$ cat /tmp/mnt-home/test1/demo1/1 
C2B4CA9C56973A3EBD65BC91ADB517E3 
```

### Via Redis

First
```
[vagrant@bogon flag]$ docker run -d --name=redis sameersbn/redis --loglevel warning
[vagrant@bogon flag]$ docker inspect -f {{.NetworkSettings.IPAddress}} redis
172.17.0.9
```

Then
```
[vagrant@localhost flag]$ docker run --rm -p 8082:8082 -v /tmp/mnt-home:/tmp/mnt-home -e SPRING_DATASOURCE_URL=jdbc:mysql://172.17.0.8:3306/testdb -e SPRING_DATASOURCE_USERNAME=testuser -e SPRING_DATASOURCE_PASSWORD=testpassword -e SPRING_REDIS_HOST=172.17.0.9 docker.io/tangfeixiong/refresh-cm:0.1
### snip ###
/v1/refresh-creation
Go to create refreshing: id:1,image_id:1,project_id:1,name:test
Search flags repository by project: 1
Hibernate: select flag0_.id as id1_2_, flag0_.round as round2_2_, flag0_.team_no as team_no3_2_, flag0_.token as token4_2_, flag0_.env as env5_2_, flag0_.md5string as md6_2_ from flag flag0_ where flag0_.env=?
Schedule Key : team1 Value : container_id:1,config_id:1,team_id:1,name:flag,sub_path:demo1/
Runnable Task with test on thread ThreadPoolTaskScheduler1
get flag with: env=1 teamNo=1 round=1
Search flags repository by project: 1
Hibernate: select flag0_.id as id1_2_, flag0_.round as round2_2_, flag0_.team_no as team_no3_2_, flag0_.token as token4_2_, flag0_.env as env5_2_, flag0_.md5string as md6_2_ from flag flag0_ where flag0_.env=?
RefreshFlag(id=50271, projectId=1, teamId=1, token=1, round=1, md5String=11BF33B2D7F8A37B59F6A128A86693FE)
Runnable Task with test on thread ThreadPoolTaskScheduler3
get flag with: env=1 teamNo=1 round=2
Search flags repository by project: 1
Hibernate: select flag0_.id as id1_2_, flag0_.round as round2_2_, flag0_.team_no as team_no3_2_, flag0_.token as token4_2_, flag0_.env as env5_2_, flag0_.md5string as md6_2_ from flag flag0_ where flag0_.env=?
RefreshFlag(id=50272, projectId=1, teamId=1, token=1, round=2, md5String=C2B4CA9C56973A3EBD65BC91ADB517E3)
### snip ###
Game over, bout : 10
```

After requested, may verify Redis cache, for example __KEYS__
```
[vagrant@localhost flag]$ echo -e "KEYS *\r\nQUIT\r\n" | curl telnet://172.17.0.9:6379
*1
$94
initialflags:??srjava.lang.Integer⠤???8Ivaluexrjava.lang.Number???
                                                                  ???xp
+OK
```

### Deploy

Refer to [`docker-compose.yml`](./docker-compose.yml)

After _docker-compose up_,  first to create database into `MySQL`

Refert to [`mysql-dump.sql`](./src/main/resources/mysql-dump.sql)

## Development

Run
```
[vagrant@bogon flag]$ mvn compile package spring-boot:run -Dspring.profiles.active=dev
```

Or
```
[vagrant@bogon flag]$ mvn compile package spring-boot:run -Dspring.profiles.active=dev -Dspring.redis.host=172.17.0.9
```

### Test

First 
```
[vagrant@localhost flag]$ mkdir -p /tmp/mnt-home
```

After do stuff of "config team" and "generate flag"

To curl
```
[vagrant@localhost flag]$ ./runtests_curl.sh create
*   Trying 172.17.4.50...
* Connected to 172.17.4.50 (172.17.4.50) port 8082 (#0)
> POST /v1/refresh-creation HTTP/1.1
> Host: 172.17.4.50:8082
> User-Agent: curl/7.43.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 641
> 
* upload completely sent off: 641 out of 641 bytes
< HTTP/1.1 200 
HTTP/1.1 200 
< Set-Cookie: JSESSIONID=E36AF703BA327A0031533396C75FCEBD; Path=/; HttpOnly
Set-Cookie: JSESSIONID=E36AF703BA327A0031533396C75FCEBD; Path=/; HttpOnly
< Content-Type: application/json;charset=UTF-8
Content-Type: application/json;charset=UTF-8
< Transfer-Encoding: chunked
Transfer-Encoding: chunked
< Date: Fri, 22 Sep 2017 10:26:18 GMT
Date: Fri, 22 Sep 2017 10:26:18 GMT

< 
* Connection #0 to host 172.17.4.50 left intact
{"refreshingDatetime":1506075993000,"id":1,"image_id":1,"battlefield_id":1,"name":"test","periodic":15,"refreshing_rfc3339":"2017-09-22T10:26:33","rounds":10,"count":0,"data_store":"test1/","state_code":0,"state_message":"","refreshing_info":{"team1":{"projectId":null,"container_id":1,"refresh_config_id":0,"team_id":1,"name":"1","sub_path":"demo1/","state_code":0,"state_message":"","flag":null}},"config":{"id":1,"common":"20170826","tremcount":null,"count":10,"environmentCount":null}}
```
 
Last to show
```
[vagrant@localhost flag]$ ls /tmp/mnt-home/test1/demo1/
1
[vagrant@localhost flag]$ cat /tmp/mnt-home/test1/demo1/1
C2B4CA9C56973A3EBD65BC91ADB517E3
```

## Reference

https://dmp.fabric8.io/#maven-goals