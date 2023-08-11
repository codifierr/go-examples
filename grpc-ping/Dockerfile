FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
ADD . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o grpc-ping  main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine
WORKDIR /app/

COPY --from=builder /app/cert/server.cert /app/cert/server.cert
COPY --from=builder /app/cert/server.key /app/cert/server.key

COPY --from=builder /app/grpc-ping  /app/grpc-ping 
ENTRYPOINT ["/app/grpc-ping"]

EXPOSE 8080/tcp
