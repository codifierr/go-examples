module grpc-ping

go 1.19

require (
	github.com/google/uuid v1.1.2
	github.com/rs/zerolog v1.28.0
	google.golang.org/grpc v1.50.1
	proto v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace proto => ./proto
