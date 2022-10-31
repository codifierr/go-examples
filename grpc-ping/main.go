package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	pb "proto"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crtFilePath = "./cert/server.cert"
	keyFilePath = "./cert/server.key"
)

// grpc server
type server struct {
	pb.UnimplementedPingProcessorServer
}

func (s *server) ProcessPing(ctx context.Context, msg *pb.Ping) (*pb.Pong, error) {
	log.Info().Msgf("Received message: %v", msg)
	traceId := uuid.New().String()
	return &pb.Pong{Id: msg.GetId(), TraceId: traceId, Message: msg.Message}, nil
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Initializing grpc echo server!")

	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		cancel()
	}()

	// start grpc server
	log.Info().Msg("Starting grpc server at 8080!")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}

	c, err := credentials.NewServerTLSFromFile(crtFilePath, keyFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Credentials from tls file")
	}

	s := grpc.NewServer(grpc.Creds(c))
	pb.RegisterPingProcessorServer(s, &server{})

	log.Info().Str("Address", lis.Addr().String()).Msg("grpc server started!")
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve")
	}
	<-ctx.Done()
	// stop server
	s.Stop()
}
