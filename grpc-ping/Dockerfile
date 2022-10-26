FROM alpine

RUN mkdir /gobin/

COPY image/grpcping /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/grpcping"]

EXPOSE 8080