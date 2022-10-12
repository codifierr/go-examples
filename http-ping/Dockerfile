FROM --platform=linux/amd64 alpine

RUN mkdir /gobin/

COPY image/httpping /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/httpping"]

EXPOSE 8080