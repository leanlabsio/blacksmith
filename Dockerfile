FROM scratch

ENV DOCKER_HOST="tcp://127.0.0.1:2376" \
    DOCKER_CERT_PATH="/certs" \
    DOCKER_TLS_VERIFY=0 \
    REDIS_ADDR="redis:6379"

COPY ./blacksmith /

VOLUME ["/certs"]

EXPOSE 9000

ENTRYPOINT ["/blacksmith"]
