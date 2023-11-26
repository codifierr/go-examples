FROM alpine

RUN mkdir /gobin/

COPY image/wsping /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/wsping"]

EXPOSE 8080
