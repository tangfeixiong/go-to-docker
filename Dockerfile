FROM busybox
LABEL "maintainer"="tangfeixiong <tangfx128@gmail.com>" \
    "project"="https://github.com/tangfeixiong/go-to-docker" \
    "name"="go-to-docker" \
    "version"="0.1" \
    "created-by"='{"name":"go-to-docker","namespace":"default","version":"0.1"}'

COPY bin/gotodocker /

ENV DOCKER_API_VERSION='1.12' \
  DOCKER_CONFIG_JSON='{"auths": {"localhost:5000": {"auth": "","email": ""}}}' \
  REGISTRY_CERTS_JSON='{"localhost:5000": {"ca_base64": "", "crt_base64": "", "key_base64": ""}}' \
  STREAMING_HOST='FAKEHOST'

EXPOSE 10051 10052

ENTRYPOINT ["/gotodocker", "serve"]
CMD ["-v", "2", "--logtostderr=true"]
