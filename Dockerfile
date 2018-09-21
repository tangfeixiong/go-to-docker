FROM busybox
LABEL "maintainer"="tangfeixiong <tangfx128@gmail.com>" \
    "project"="https://github.com/tangfeixiong/go-to-docker" \
    "name"="msgo2docker" \
    "version"="0.2"

# ca-certificates
#apk add --update ca-certificates && rm -rf /var/cache/apk/*
COPY etc/ssl/certs/ /etc/ssl/certs/
	
COPY bin/gotodocker /

# https://docs.docker.com/engine/reference/api/docker_remote_api/
# docker version should be at least 1.11.x. this mean that minimal docker api vesion is 1.23
#ENV DOCKER_API_VERSION='1.23' \
#  DOCKER_CONFIG_JSON='{"auths": {"localhost:5000": {"auth": "","email": ""}}}' \
#  REGISTRY_CERTS_JSON='{"localhost:5000": {"ca_base64": "", "crt_base64": "", "key_base64": ""}}' \
#  STREAMING_HOST='FAKEHOST'

# EXPOSE 10051 10052
EXPOSE 10053

ENTRYPOINT ["/gotodocker", "serve2"]
CMD ["--logtostderr=true", "-v", "2"]
