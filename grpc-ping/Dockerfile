FROM alpine

RUN mkdir /gobin/

COPY cert/server.cert /gobin/cert/server.cert

COPY cert/server.key /gobin/cert/server.key

COPY image/grpcping /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/grpcping"]

EXPOSE 8080