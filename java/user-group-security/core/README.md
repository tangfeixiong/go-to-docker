

## Development

Maven
```
[vagrant@kubedev-172-17-4-59 user-group-security]$ mvn install --projects core
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building core 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:resources (default-resources) @ core ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/core/src/main/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:compile (default-compile) @ core ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 46 source files to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/core/target/classes
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:testResources (default-testResources) @ core ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/core/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:testCompile (default-testCompile) @ core ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-surefire-plugin:2.12.4:test (default-test) @ core ---
[INFO] No tests to run.
[INFO] 
[INFO] --- maven-jar-plugin:2.4:jar (default-jar) @ core ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/core/target/core-0.0.1-SNAPSHOT.jar
[INFO] 
[INFO] --- maven-install-plugin:2.4:install (default-install) @ core ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/core/target/core-0.0.1-SNAPSHOT.jar to /home/vagrant/.m2/repository/https0x3A0x2F0x2Fgithub0x2Ecom0x2Fstackdocker/fairymagic/core/0.0.1-SNAPSHOT/core-0.0.1-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/core/pom.xml to /home/vagrant/.m2/repository/https0x3A0x2F0x2Fgithub0x2Ecom0x2Fstackdocker/fairymagic/core/0.0.1-SNAPSHOT/core-0.0.1-SNAPSHOT.pom
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 36.942 s
[INFO] Finished at: 2018-06-18T01:46:25Z
[INFO] Final Memory: 28M/69M
[INFO] ------------------------------------------------------------------------
```

