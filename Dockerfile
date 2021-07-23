FROM alpine:3.14.0

ADD ./kiam-watchdog /kiam-watchdog

ENTRYPOINT ["/kiam-watchdog"]
