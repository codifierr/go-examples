module github.com/codifierr/go-kafka

go 1.16

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/segmentio/kafka-go v0.4.39 // indirect
	gokafka v0.0.0-00010101000000-000000000000
)

replace gokafka => ./lib/gokafka
