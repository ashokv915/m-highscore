package grpc

import (
	pbhighscore "github.com/ashok/m-apis/m-highscore/v1"
	"google.golang.org/grpc"
	"context"
	"github.com/rs/zerolog/log"
)

// Grpc struct is used for connecting
type Grpc struct {
	address string
	srv *grpc.Server
}

//HighScore is the highest score
var HighScore = 9999999999999.0

// SetHighScore is used to set Highscore
func (g *Grpc) SetHighScore(ctx context.Context, in *pbhighscore.SetHighScoreRequest, opts ...grpc.CallOption) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in m-highscore is called")
	HighScore = in.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

// GetHighScore is used to get Highscore
func (g *Grpc) GetHighScore(ctx context.Context, in *pbhighscore.GetHighScoreRequest, opts ...grpc.CallOption) (*pbhighscore.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in m-highscore is called")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}


