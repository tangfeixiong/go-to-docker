# FROM openshift/base-centos7
FROM docker.io/centos:centos7
LABEL maintainer="tangfeixiong <tangfx128@gmail.com>" \
    project="cepo/ctf" \
    name="flag-rest-server" \
    repository="default" \
    tag="centos java springboot"

ARG jarTgt
ARG javaOpt

COPY /maven/${jarTgt:-Flag.jar} /

ENV JAVA_OPTIONS="${javaOpt:-'-Xms128m -Xmx512m -XX:PermSize=128m -XX:MaxPermSize=256m'}" \
    SPRING_DATASOURCE_URL="jdbc:mysql://localhost:3306/testdb" \
    SPRING_DATASOURCE_USERNAME="testuser" \
    SPRING_DATASOURCE_PASSWORD="testpassword" \
    SERVER_PORT="8082"

RUN set -x \
    && install_Pkgs=" \
        tar \
        unzip \
        bc \
        which \
        lsof \
        java-1.8.0-openjdk-headless \
    " \
    && yum install -y $install_Pkgs \
    && yum clean all -y \
    && echo

# This default user is created in the openshift/base-centos7 image
# USER 1001

CMD java -Djava.security.egd=file:/dev/./urandom $JAVA_OPTIONS -jar /${jarTgt:-Flag.jar}