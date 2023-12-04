FROM alpine:3.18.5

ADD ./kiam-watchdog /kiam-watchdog

ENTRYPOINT ["/kiam-watchdog"]
