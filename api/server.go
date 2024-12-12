package api

import (
	"github.com/bdamarjati/go_book_collection/db/sqlc"
	"github.com/bdamarjati/go_book_collection/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config util.Config
	store  sqlc.Store
}

func NewServer(config util.Config, store sqlc.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
