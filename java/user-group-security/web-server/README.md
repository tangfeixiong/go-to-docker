

```
fanhonglingdeMacBook-Pro:user-group-security fanhongling$ cd web-server/
```

```
fanhonglingdeMacBook-Pro:user-group-security fanhongling$ mvn package spring-boot:run -Dmaven.test.skip=true -Dspring.profiles.active=dev
```

```
tgt=$(ls web-server/target/web-server*.war)
docker build --force-rm --no-cache \
    -t docker.io/tangfeixiong/ugs \
    --build-arg jarTgt=web-server/target/${basename $tgt} \
    -f Dockerfile.tomcat ./web-server

```