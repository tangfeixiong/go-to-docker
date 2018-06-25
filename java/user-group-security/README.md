# Help

A Spring Boot web app with Spring Boot restful server

## Development

Build parent only
```
fanhonglingdeMacBook-Pro:user-group-security fanhongling$ mvn install --non-recursive
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building Organization Management 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-install-plugin:2.4:install (default-install) @ user-group-security ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/java/user-group-security/pom.xml to /Users/fanhongling/.m2/repository/https0x3A0x2F0x2Fgithub0x2Ecom0x2Fstackdocker/fairymagic/user-group-security/0.0.1-SNAPSHOT/user-group-security-0.0.1-SNAPSHOT.pom
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 0.579 s
[INFO] Finished at: 2018-06-20T19:16:51-07:00
[INFO] Final Memory: 9M/309M
[INFO] ------------------------------------------------------------------------
```

### About

sql maven plugin
```
fanhonglingdeMacBook-Pro:github.com fanhongling$ git clone https://github.com/mojohaus/sql-maven-plugin mojohaus/sql-maven-plugin
Cloning into 'mojohaus/sql-maven-plugin'...
remote: Counting objects: 2958, done.
remote: Total 2958 (delta 0), reused 0 (delta 0), pack-reused 2958
Receiving objects: 100% (2958/2958), 1.00 MiB | 206.00 KiB/s, done.
Resolving deltas: 100% (911/911), done.
Checking connectivity... done.
```

```
fanhonglingdeMacBook-Pro:github.com fanhongling$ cd mojohaus/sql-maven-plugin/
```

```
fanhonglingdeMacBook-Pro:sql-maven-plugin fanhongling$ mvn compile package install -Dmaven.test.skip=true
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building SQL Maven Plugin 3.0.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------
Downloading from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-component-metadata/1.7/plexus-component-metadata-1.7.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-component-metadata/1.7/plexus-component-metadata-1.7.pom (3.9 kB at 2.0 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-component-metadata/1.7/plexus-component-metadata-1.7.jar
Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-component-metadata/1.7/plexus-component-metadata-1.7.jar (117 kB at 139 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/axion/axion/1.0-M2-dev/axion-1.0-M2-dev.pom
Downloaded from central: https://repo.maven.apache.org/maven2/axion/axion/1.0-M2-dev/axion-1.0-M2-dev.pom (1.1 kB at 1.1 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/commons-collections/commons-collections/3.0/commons-collections-3.0.pom
Downloaded from central: https://repo.maven.apache.org/maven2/commons-collections/commons-collections/3.0/commons-collections-3.0.pom (6.1 kB at 20 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/commons-primitives/commons-primitives/1.0/commons-primitives-1.0.pom
Downloaded from central: https://repo.maven.apache.org/maven2/commons-primitives/commons-primitives/1.0/commons-primitives-1.0.pom (168 B at 547 B/s)
Downloading from central: https://repo.maven.apache.org/maven2/javacc/javacc/3.2/javacc-3.2.pom
Downloaded from central: https://repo.maven.apache.org/maven2/javacc/javacc/3.2/javacc-3.2.pom (281 B at 930 B/s)
Downloading from central: https://repo.maven.apache.org/maven2/net/java/dev/javacc/javacc/3.2/javacc-3.2.pom
Downloaded from central: https://repo.maven.apache.org/maven2/net/java/dev/javacc/javacc/3.2/javacc-3.2.pom (162 B at 531 B/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/apache/maven/shared/maven-script-interpreter/1.1/maven-script-interpreter-1.1.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/apache/maven/shared/maven-script-interpreter/1.1/maven-script-interpreter-1.1.pom (3.7 kB at 12 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/codehaus/groovy/groovy/2.0.1/groovy-2.0.1.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/groovy/groovy/2.0.1/groovy-2.0.1.pom (16 kB at 49 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-tree/4.0/asm-tree-4.0.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-tree/4.0/asm-tree-4.0.pom (2.1 kB at 6.9 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-parent/4.0/asm-parent-4.0.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-parent/4.0/asm-parent-4.0.pom (0 B at 0 B/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm/4.0/asm-4.0.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm/4.0/asm-4.0.pom (0 B at 0 B/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-commons/4.0/asm-commons-4.0.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-commons/4.0/asm-commons-4.0.pom (2.1 kB at 6.8 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-util/4.0/asm-util-4.0.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-util/4.0/asm-util-4.0.pom (2.1 kB at 6.7 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-analysis/4.0/asm-analysis-4.0.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-analysis/4.0/asm-analysis-4.0.pom (2.1 kB at 6.9 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-utils/3.0.7/plexus-utils-3.0.7.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-utils/3.0.7/plexus-utils-3.0.7.pom (2.5 kB at 8.2 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/axion/axion/1.0-M2-dev/axion-1.0-M2-dev.jar
Downloading from central: https://repo.maven.apache.org/maven2/commons-collections/commons-collections/3.0/commons-collections-3.0.jar
Downloading from central: https://repo.maven.apache.org/maven2/commons-primitives/commons-primitives/1.0/commons-primitives-1.0.jar
Downloading from central: https://repo.maven.apache.org/maven2/commons-logging/commons-logging/1.0/commons-logging-1.0.jar
Downloading from central: https://repo.maven.apache.org/maven2/net/java/dev/javacc/javacc/3.2/javacc-3.2.jar
Downloaded from central: https://repo.maven.apache.org/maven2/commons-logging/commons-logging/1.0/commons-logging-1.0.jar (22 kB at 11 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/apache/maven/shared/maven-script-interpreter/1.1/maven-script-interpreter-1.1.jar
Downloaded from central: https://repo.maven.apache.org/maven2/net/java/dev/javacc/javacc/3.2/javacc-3.2.jar (379 kB at 177 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/codehaus/groovy/groovy/2.0.1/groovy-2.0.1.jar
Downloaded from central: https://repo.maven.apache.org/maven2/org/apache/maven/shared/maven-script-interpreter/1.1/maven-script-interpreter-1.1.jar (21 kB at 8.0 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-tree/4.0/asm-tree-4.0.jar
Downloaded from central: https://repo.maven.apache.org/maven2/axion/axion/1.0-M2-dev/axion-1.0-M2-dev.jar (412 kB at 145 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-commons/4.0/asm-commons-4.0.jar
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-commons/4.0/asm-commons-4.0.jar (38 kB at 12 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm/4.0/asm-4.0.jar
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-tree/4.0/asm-tree-4.0.jar (22 kB at 6.9 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-util/4.0/asm-util-4.0.jar
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm/4.0/asm-4.0.jar (0 B at 0 B/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-analysis/4.0/asm-analysis-4.0.jar
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-analysis/4.0/asm-analysis-4.0.jar (20 kB at 5.4 kB/s)
Downloaded from central: https://repo.maven.apache.org/maven2/org/ow2/asm/asm-util/4.0/asm-util-4.0.jar (37 kB at 9.7 kB/s)
Downloaded from central: https://repo.maven.apache.org/maven2/commons-primitives/commons-primitives/1.0/commons-primitives-1.0.jar (260 kB at 65 kB/s)
Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/groovy/groovy/2.0.1/groovy-2.0.1.jar (3.3 MB at 582 kB/s)
Downloaded from central: https://repo.maven.apache.org/maven2/commons-collections/commons-collections/3.0/commons-collections-3.0.jar (519 kB at 84 kB/s)
[INFO] 
[INFO] --- maven-enforcer-plugin:1.4:enforce (mojo-enforcer-rules) @ sql-maven-plugin ---
[INFO] 
[INFO] --- maven-resources-plugin:3.0.1:resources (default-resources) @ sql-maven-plugin ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 2 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.5.1:compile (default-compile) @ sql-maven-plugin ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 4 source files to /Users/fanhongling/Downloads/workspace/src/github.com/mojohaus/sql-maven-plugin/target/classes
[INFO] 
[INFO] --- maven-enforcer-plugin:1.4:enforce (mojo-enforcer-rules) @ sql-maven-plugin ---
[INFO] 
[INFO] --- maven-resources-plugin:3.0.1:resources (default-resources) @ sql-maven-plugin ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 2 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.5.1:compile (default-compile) @ sql-maven-plugin ---
[INFO] Nothing to compile - all classes are up to date
[INFO] 
[INFO] --- maven-plugin-plugin:3.4:descriptor (default-descriptor) @ sql-maven-plugin ---
[INFO] Using 'UTF-8' encoding to read mojo metadata.
[INFO] Mojo extractor with id: java-javadoc found 0 mojo descriptors.
[INFO] Mojo extractor with id: java-annotations found 1 mojo descriptors.
[INFO] 
[INFO] --- maven-plugin-plugin:3.4:helpmojo (help-mojo) @ sql-maven-plugin ---
[INFO] Using 'UTF-8' encoding to read mojo metadata.
[INFO] Mojo extractor with id: java-javadoc found 0 mojo descriptors.
[INFO] Mojo extractor with id: java-annotations found 1 mojo descriptors.
[INFO] 
[INFO] --- plexus-component-metadata:1.7:generate-metadata (default) @ sql-maven-plugin ---
Downloading from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-container-default/1.7/plexus-container-default-1.7.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-container-default/1.7/plexus-container-default-1.7.pom (2.7 kB at 8.5 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/com/thoughtworks/qdox/qdox/2.0-M2/qdox-2.0-M2.pom
Downloaded from central: https://repo.maven.apache.org/maven2/com/thoughtworks/qdox/qdox/2.0-M2/qdox-2.0-M2.pom (18 kB at 36 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-cli/1.6/plexus-cli-1.6.pom
Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-cli/1.6/plexus-cli-1.6.pom (1.4 kB at 3.1 kB/s)
Downloading from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-container-default/1.7/plexus-container-default-1.7.jar
Downloading from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-cli/1.6/plexus-cli-1.6.jar
Downloading from central: https://repo.maven.apache.org/maven2/com/thoughtworks/qdox/qdox/2.0-M2/qdox-2.0-M2.jar
Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-cli/1.6/plexus-cli-1.6.jar (7.4 kB at 8.9 kB/s)
Downloaded from central: https://repo.maven.apache.org/maven2/com/thoughtworks/qdox/qdox/2.0-M2/qdox-2.0-M2.jar (292 kB at 136 kB/s)
Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-container-default/1.7/plexus-container-default-1.7.jar (231 kB at 107 kB/s)
[INFO] 
[INFO] --- animal-sniffer-maven-plugin:1.15:check (default) @ sql-maven-plugin ---
[INFO] Checking unresolved references to org.codehaus.mojo.signature:java16:1.0
[INFO] 
[INFO] --- maven-resources-plugin:3.0.1:testResources (default-testResources) @ sql-maven-plugin ---
[INFO] Not copying test resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.5.1:testCompile (default-testCompile) @ sql-maven-plugin ---
[INFO] Not compiling test sources
[INFO] 
[INFO] --- maven-surefire-plugin:2.19.1:test (default-test) @ sql-maven-plugin ---
[INFO] Tests are skipped.
[INFO] 
[INFO] --- maven-jar-plugin:3.0.2:jar (default-jar) @ sql-maven-plugin ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/mojohaus/sql-maven-plugin/target/sql-maven-plugin-3.0.0-SNAPSHOT.jar
[INFO] 
[INFO] --- maven-plugin-plugin:3.4:addPluginArtifactMetadata (default-addPluginArtifactMetadata) @ sql-maven-plugin ---
[INFO] 
[INFO] --- maven-enforcer-plugin:1.4:enforce (mojo-enforcer-rules) @ sql-maven-plugin ---
[INFO] 
[INFO] --- maven-resources-plugin:3.0.1:resources (default-resources) @ sql-maven-plugin ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 2 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.5.1:compile (default-compile) @ sql-maven-plugin ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 5 source files to /Users/fanhongling/Downloads/workspace/src/github.com/mojohaus/sql-maven-plugin/target/classes
[INFO] 
[INFO] --- maven-plugin-plugin:3.4:descriptor (default-descriptor) @ sql-maven-plugin ---
[INFO] Using 'UTF-8' encoding to read mojo metadata.
[INFO] Mojo extractor with id: java-javadoc found 0 mojo descriptors.
[INFO] Mojo extractor with id: java-annotations found 2 mojo descriptors.
[INFO] 
[INFO] --- maven-plugin-plugin:3.4:helpmojo (help-mojo) @ sql-maven-plugin ---
[INFO] Using 'UTF-8' encoding to read mojo metadata.
[INFO] Mojo extractor with id: java-javadoc found 0 mojo descriptors.
[INFO] Mojo extractor with id: java-annotations found 2 mojo descriptors.
[INFO] 
[INFO] --- plexus-component-metadata:1.7:generate-metadata (default) @ sql-maven-plugin ---
[INFO] 
[INFO] --- animal-sniffer-maven-plugin:1.15:check (default) @ sql-maven-plugin ---
[INFO] Checking unresolved references to org.codehaus.mojo.signature:java16:1.0
[INFO] 
[INFO] --- maven-resources-plugin:3.0.1:testResources (default-testResources) @ sql-maven-plugin ---
[INFO] Not copying test resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.5.1:testCompile (default-testCompile) @ sql-maven-plugin ---
[INFO] Not compiling test sources
[INFO] 
[INFO] --- maven-surefire-plugin:2.19.1:test (default-test) @ sql-maven-plugin ---
[INFO] Tests are skipped.
[INFO] 
[INFO] --- maven-jar-plugin:3.0.2:jar (default-jar) @ sql-maven-plugin ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/mojohaus/sql-maven-plugin/target/sql-maven-plugin-3.0.0-SNAPSHOT.jar
[INFO] 
[INFO] --- maven-plugin-plugin:3.4:addPluginArtifactMetadata (default-addPluginArtifactMetadata) @ sql-maven-plugin ---
[INFO] 
[INFO] --- maven-install-plugin:2.5.2:install (default-install) @ sql-maven-plugin ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/mojohaus/sql-maven-plugin/target/sql-maven-plugin-3.0.0-SNAPSHOT.jar to /Users/fanhongling/.m2/com.magicbird/org/codehaus/mojo/sql-maven-plugin/3.0.0-SNAPSHOT/sql-maven-plugin-3.0.0-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/mojohaus/sql-maven-plugin/pom.xml to /Users/fanhongling/.m2/com.magicbird/org/codehaus/mojo/sql-maven-plugin/3.0.0-SNAPSHOT/sql-maven-plugin-3.0.0-SNAPSHOT.pom
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 27.111 s
[INFO] Finished at: 2018-06-12T16:04:42-07:00
[INFO] Final Memory: 30M/317M
[INFO] ------------------------------------------------------------------------
```