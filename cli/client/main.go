package main

import (
	"flag"
	pbhighscore "github.com/ashok/m-apis/m-highscore/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"context"
	"time"
)

func main() {
	var addressPtr = flag.String("address","localhost:50051","address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial m-highscore gRPC service")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("address", *addressPtr).Msg("Failed to Close a connnection")
		}
	}()

	timeCtxt, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c := pbhighscore.NewGameClient(conn)

	if c == nil {
		log.Info().Msg("Client nil")
	}


	r, err := c.GetHighScore(timeCtxt, &pbhighscore.GetHighScoreRequest{}) 

	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}

	if r != nil {
		log.Info().Interface("highscore", r.GetHighScore()).Msg("highscore from m-highscore microservice")
	} else {
		log.Error().Msg("Couldn't get highscore!")
	}

}