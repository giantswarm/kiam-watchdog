FROM alpine:3.18.4

ADD ./kiam-watchdog /kiam-watchdog

ENTRYPOINT ["/kiam-watchdog"]
