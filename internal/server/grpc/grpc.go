package grpc

import (
	pbhighscore "github.com/ashok/m-apis/m-highscore/v1"
	"google.golang.org/grpc"
	"context"
	"github.com/rs/zerolog/log"
	"net"
	"github.com/pkg/errors"
)

// Grpc struct is used for connecting
type Grpc struct {
	address string
	srv     *grpc.Server
}

//HighScore is the highest score
var HighScore = 9999999999999.0

// NewServer return address
func NewServer (address string) *Grpc {
	return &Grpc {
		address: address,
	}
}

// SetHighScore is used to set Highscore
func (g *Grpc) SetHighScore(ctx context.Context, in *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in m-highscore is called")
	HighScore = in.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

// GetHighScore is used to get Highscore
func (g *Grpc) GetHighScore(ctx context.Context, in *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in m-highscore is called")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

// ListenAndServe helps to listen and serve
func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "Failed to open the tcp port")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	pbhighscore.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for m-highscore microservice")
	
	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "Failed to start gRPC server for m-higscore microservice")
	}
	return nil
}


