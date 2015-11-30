FROM scratch

ENV DOCKER_HOST="tcp://127.0.0.1:2376" \
    DOCKER_CERT_PATH="/certs"

COPY ./blacksmith /

VOLUME ["/certs"]

EXPOSE 9000

ENTRYPOINT ["/blacksmith"]