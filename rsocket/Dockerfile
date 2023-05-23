FROM alpine

RUN mkdir /gobin/

COPY image/rsocketping /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/rsocketping"]

EXPOSE 8080