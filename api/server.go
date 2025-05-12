package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luke_design_pattern/config"
	"github.com/luke_design_pattern/db"
)

type Server struct {
	config config.CredentialDB
	store  db.Store
	router *gin.Engine
}

func NewServer(config *config.CredentialDB, store db.Store) (*Server, error) {

	server := &Server{
		config: *config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// controller rest API not method
	router.POST("/users/createTx", server.CreateUser)

	// later kasih authorization routes

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
