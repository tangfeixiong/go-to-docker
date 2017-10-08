FROM scratch
LABEL "maintainer"="tangfeixiong <tangfx128@gmail.com>" \
    "project"="https://github.com/tangfeixiong/go-to-docker" \
    "name"="metering" \
    "version"="alpha" \
    "namespace"="stackdocker" \
    "annotation"='{"stackdocker.io/created-by":"n/a"}' \
    "tag"="linux golang"

COPY bin/metering /

EXPOSE 12315 12316
# EXPOSE 12305 12306

ENTRYPOINT ["/metering", "--v=2", "--logtostderr"]
CMD ["export", "--meter=cadvisor=http://localhost:8080", "--collector=grpc=localhost:12305"]
# CMD ["collect", "--storage=elasticsearch=http://localhost:9200"]
