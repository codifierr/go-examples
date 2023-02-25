module github.com/codifierr/go-examples/grpc/v0.0.1

go 1.17

require (
	github.com/google/uuid v1.1.2
	github.com/rs/zerolog v1.26.0
	google.golang.org/grpc v1.42.0
	proto v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace proto => ./proto
