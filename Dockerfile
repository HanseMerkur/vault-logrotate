FROM alpine:3.16
LABEL author="Lennart Weller <lennart.weller@hansemerkur.de>

ENV CRONTAB="0 * * * *"

RUN apk add --no-cache logrotate \
    && adduser \
        --uid 1000 \
        --gid 1000 \
        --no-create-home \
        --home "/tmp" \
        --disabled-password \
        crond

ADD crond-logrotate /usr/local/bin/crond-logrotate

ENTRYPOINT ["/usr/local/bin/crond-logrotate"]
