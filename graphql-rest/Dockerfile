FROM alpine

RUN mkdir /gobin/

COPY queries/ /gobin/queries
COPY config.json /gobin
COPY graphql /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/graphql"]

EXPOSE 8080
