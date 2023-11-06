package animeMovie

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
)

// AnimeMovieServer serves gRPC requests for Anime Movie endpoints.
type AnimeMovieServer struct {
	config     utils.Config
	gojo       db.Gojo
	tokenMaker token.Maker
}

func NewAnimeMovieServer(config utils.Config, gojo db.Gojo, tokenMaker token.Maker) *AnimeMovieServer {
	server := &AnimeMovieServer{
		config:     config,
		gojo:       gojo,
		tokenMaker: tokenMaker,
	}

	return server
}