package main

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Initializing stats processor!")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//test goredis connection
	rClient := redis.NewClient(&redis.Options{
		Addr:     "localhost" + ":" + "31631",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Info().Str("Redis Ping", rClient.Ping(ctx).String()).Msg("Redis service!")

}
