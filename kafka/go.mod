module github.com/codifierr/go-kafka

go 1.20

require (
	github.com/google/uuid v1.3.0
	gokafka v0.0.0-00010101000000-000000000000
)

require (
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/segmentio/kafka-go v0.4.42 // indirect
)

replace gokafka => ./lib/gokafka
