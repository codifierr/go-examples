module github.com/codifierr/go-examples/grpc/v0.0.1

go 1.17

require (
	github.com/google/uuid v1.3.0
	github.com/rs/zerolog v1.29.0
	google.golang.org/grpc v1.53.0
	proto v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/protobuf v1.29.1 // indirect
)

replace proto => ./proto
