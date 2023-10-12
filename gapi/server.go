package gapi

import (
	"fmt"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
)

// Server serves gRPC requests for our Saber service.
type Server struct {
	pb.UnimplementedGojoServer
	config utils.Config
	gojo   db.Gojo
	token  token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config utils.Config, gojo db.Gojo) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config: config,
		gojo:   gojo,
		token:  tokenMaker,
	}

	return server, nil
}
