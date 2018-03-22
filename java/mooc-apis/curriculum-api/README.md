# Development

## Build and Run

Maven (via aliyun mirror)
```
[vagrant@kubedev-172-17-4-59 curriculum-api]$ mvn package spring-boot:run
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building Go to Docker :: MOOC API :: Curriculum API 0.0.1-alpha
[INFO] ------------------------------------------------------------------------
Downloading from aliyun: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/boot/spring-boot-starter-web/2.0.0.RELEASE/spring-boot-starter-web-2.0.0.RELEASE.pom
Downloaded from aliyun: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/boot/spring-boot-starter-web/2.0.0.RELEASE/spring-boot-starter-web-2.0.0.RELEASE.pom (3.2 kB at 982 B/s)
<<< ...snippets... >>>
Downloaded from aliyun: http://maven.aliyun.com/nexus/content/groups/public/org/apache/tomcat/embed/tomcat-embed-core/8.5.28/tomcat-embed-core-8.5.28.jar (3.1 MB at 227 kB/s)
[INFO] 
[INFO] --- spring-boot-maven-plugin:2.0.0.RELEASE:build-info (generate build info) @ curriculum-api ---
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:resources (default-resources) @ curriculum-api ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 9 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:compile (default-compile) @ curriculum-api ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:testResources (default-testResources) @ curriculum-api ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/mooc-apis/curriculum-api/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:testCompile (default-testCompile) @ curriculum-api ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-surefire-plugin:2.12.4:test (default-test) @ curriculum-api ---
[INFO] No tests to run.
[INFO] 
[INFO] --- maven-jar-plugin:2.4:jar (default-jar) @ curriculum-api ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/mooc-apis/curriculum-api/target/curriculum-api-0.0.1-alpha.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:2.0.0.RELEASE:repackage (default) @ curriculum-api ---
[INFO] Layout: JAR
[INFO] 
[INFO] --- dockerfile-maven-plugin:1.3.7:build (default) @ curriculum-api ---
[INFO] Skipping execution because 'dockerfile.skip' is set
[INFO] 
[INFO] >>> spring-boot-maven-plugin:2.0.0.RELEASE:run (default-cli) > test-compile @ curriculum-api >>>
[INFO] 
[INFO] --- spring-boot-maven-plugin:2.0.0.RELEASE:build-info (generate build info) @ curriculum-api ---
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:resources (default-resources) @ curriculum-api ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 9 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:compile (default-compile) @ curriculum-api ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:testResources (default-testResources) @ curriculum-api ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/mooc-apis/curriculum-api/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:testCompile (default-testCompile) @ curriculum-api ---
[INFO] No sources to compile
[INFO] 
[INFO] <<< spring-boot-maven-plugin:2.0.0.RELEASE:run (default-cli) < test-compile @ curriculum-api <<<
[INFO] 
[INFO] 
[INFO] --- spring-boot-maven-plugin:2.0.0.RELEASE:run (default-cli) @ curriculum-api ---

  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::        (v2.0.0.RELEASE)

2018-03-22 00:54:18.393  INFO 28332 --- [           main] io.stackdocker.curriculum.App            : Starting App on kubedev-172-17-4-59 with PID 28332 (/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/mooc-apis/curriculum-api/target/classes started by vagrant in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/mooc-apis/curriculum-api)
2018-03-22 00:54:18.411  INFO 28332 --- [           main] io.stackdocker.curriculum.App            : No active profile set, falling back to default profiles: default
<<< ...snippets... >>>
2018-03-22 00:54:35.645  INFO 28332 --- [           main] o.s.b.w.embedded.tomcat.TomcatWebServer  : Tomcat started on port(s): 8080 (http) with context path ''
2018-03-22 00:54:35.684  INFO 28332 --- [           main] io.stackdocker.curriculum.App            : Started App in 19.196 seconds (JVM running for 63.134)

```

Then go __http://<your machine IP>:8080/__

Terminate CI (via ctrl-c)
```
^C2018-03-22 00:57:51.835  INFO 28332 --- [       Thread-3] ConfigServletWebServerApplicationContext : Closing org.springframework.boot.web.servlet.context.AnnotationConfigServletWebServerApplicationContext@9dd30a2: startup date [Thu Mar 22 00:54:18 UTC 2018]; root of context hierarchy
2018-03-22 00:57:51.840  INFO 28332 --- [       Thread-3] o.s.j.e.a.AnnotationMBeanExporter        : Unregistering JMX-exposed beans on shutdown
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
```

## CI/CD (Continuously Integration & Continuously Delivery)

Docker
```
[vagrant@kubedev-172-17-4-59 curriculum-api]$ mvn package spring-boot:repackage docker:build

```

### Reference

https://spring.io/guides/gs/serving-web-content/

https://github.com/spring-guides/gs-serving-web-content

https://github.com/spring-projects/spring-boot/tree/master/spring-boot-samples/spring-boot-sample-web-ui

https://github.com/spring-projects/spring-boot/tree/master/spring-boot-samples/spring-boot-sample-web-static