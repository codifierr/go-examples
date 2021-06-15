module github.com/codifierr/go-kafka

go 1.16

require (
	github.com/google/uuid v1.2.0
	github.com/segmentio/kafka-go v0.4.16 // indirect
	gokafka v0.0.0-00010101000000-000000000000
)

replace gokafka => ./lib/gokafka
