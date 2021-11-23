package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	pb "proto"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// grpc server
type server struct {
	pb.UnimplementedMessageProcessorServer
}

func (s *server) ProcessMessages(ctx context.Context, msg *pb.MessagePayload) (*pb.Ack, error) {
	log.Info().Msgf("Received message: %v", msg)
	traceId := uuid.New().String()
	return &pb.Ack{Id: msg.GetId(), TraceId: traceId, Message: "Message processed successfully"}, nil
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Initializing grpc example!")

	_, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	cleanupDone := make(chan bool)
	go func() {
		for range signalChan {
			cancel()
			cleanupDone <- true
		}
	}()

	// start grpc server
	log.Info().Msg("Starting grpc server!")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8083))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}

	s := grpc.NewServer()
	pb.RegisterMessageProcessorServer(s, &server{})

	log.Info().Str("Address", lis.Addr().String()).Msg("grpc server started!")
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve")
	}

	time.Sleep(5 * time.Second)
	<-cleanupDone
	// stop server
	s.Stop()
}
