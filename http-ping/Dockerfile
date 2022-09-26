FROM --platform=linux/amd64 alpine

RUN mkdir /gobin/

# RUN apk update \
#     && apk upgrade \
#     && apk add --no-cache libc-dev \
#     && apk add --no-cache ca-certificates \
#     && update-ca-certificates 2>/dev/null || true

COPY image/httpping /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/httpping"]

EXPOSE 8080