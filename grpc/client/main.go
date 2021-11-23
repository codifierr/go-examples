package main

import (
	"context"
	"time"

	pb "proto"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	log.Info().Msg("Starting message processor grpc client")

	con, err := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to stats processor")
	}
	defer con.Close()

	client := pb.NewMessageProcessorClient(con)
	// Contact the server and print out its response.
	for {
		SendMessage(client)
	}
}

func SendMessage(client pb.MessageProcessorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	Id := uuid.New().String()
	resp, err := client.ProcessMessages(ctx, &pb.MessagePayload{Id: Id, Message: "This is test message"})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to send request stats")
	}
	log.Info().Msgf("Response: %v", resp)

	time.Sleep(500 * time.Millisecond)
}
