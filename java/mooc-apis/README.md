

### MVN Wrapper

https://github.com/takari/maven-wrapper 
```
[vagrant@kubedev-172-17-4-59 mooc-apis]$ mvn -N io.takari:maven:wrapper -Dmaven=3.5.2
[INFO] Scanning for projects...
Downloading from aliyun: http://maven.aliyun.com/nexus/content/groups/public/io/takari/maven/maven-metadata.xml
Downloading from spring-plugins-milestone: https://repo.spring.io/plugins-milestone/io/takari/maven/maven-metadata.xml
Downloading from spring-plugins-snapshot: https://repo.spring.io/plugins-snapshot/io/takari/maven/maven-metadata.xml
Downloaded from spring-plugins-snapshot: https://repo.spring.io/plugins-snapshot/io/takari/maven/maven-metadata.xml (689 B at 174 B/s)
Downloaded from aliyun: http://maven.aliyun.com/nexus/content/groups/public/io/takari/maven/maven-metadata.xml (689 B at 168 B/s)
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building Go to Docker :: MOOC API Project 0.0.1-alpha
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven:0.6.0:wrapper (default-cli) @ mooc-apis ---
[INFO] 
[INFO] The Maven Wrapper version 0.4.0 has been successfully setup for your project.
[INFO] Using Apache Maven 3.5.2
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 9.724 s
[INFO] Finished at: 2018-03-21T23:56:10Z
[INFO] Final Memory: 11M/31M
[INFO] ------------------------------------------------------------------------
```

```
[vagrant@kubedev-172-17-4-59 mooc-apis]$ ls mvn*
mvnw  mvnw.cmd
```