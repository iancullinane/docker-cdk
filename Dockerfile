FROM alpine:3.10

USER root
WORKDIR /root/cdkbot

RUN apk add --no-cache nodejs npm make gcc libc-dev git go

COPY ./bin/docker-cdk .

ENTRYPOINT ["./docker-cdk"]
