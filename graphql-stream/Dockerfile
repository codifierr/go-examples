FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.20.4 AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
ADD . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o graphql-stream  main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:latest
WORKDIR /app/
COPY --from=builder /app/graphql-stream  /app/graphql-stream
ENTRYPOINT ["/app/graphql-stream"]

EXPOSE 8080/tcp
