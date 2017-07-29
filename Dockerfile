FROM busybox
LABEL "maintainer"="tangfeixiong <tangfx128@gmail.com>" \
    "project"="https://github.com/tangfeixiong/go-to-docker" \
    "name"="go-to-docker" \
    "version"="0.1" \
    "created-by"='{"name":"go-to-docker","namespace":"default","version":"0.1"}'

COPY bin/gotodocker /

ENV DOCKER_API_VERSION='1.12' \
  DOCKER_CONFIG_JSON='{"auths": {"127.0.0.1:5000": {"auth": "","email": ""}}}' \
  STREAMING_HOST='FAKEHOST'

EXPOSE 10051 10052

ENTRYPOINT ["/gotodocker", "serve"]
CMD ["-v", "2", "--logtostderr=true"]
