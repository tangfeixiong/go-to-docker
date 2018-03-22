ROM openjdk:8-jre
LABEL maintainer="tangfeixiong <tangfx128@gmail.com>" \
      project="https://github.com/tangfeixiong/go-to-docker" \
      name="curriculum-api" \
      annotation='{"example.com/mooc-apis":"curriculum-api"}' \
      tag="centos java1.8 openjdk springboot"

ARG jarTgt
ARG javaOpt

COPY ${jarTgt:-curriculum-api.jar} /curriculum-api.jar

ENV JAVA_OPTIONS="${javaOpt:-'-Xms128m -Xmx512m -XX:PermSize=128m -XX:MaxPermSize=256m'}" \
    GO_TO_DOCKER="curriculum-api"

ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/curriculum-api.jar"]
