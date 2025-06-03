package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/luke_design_pattern/cmd/app/docs"
	"github.com/luke_design_pattern/config"
	"github.com/luke_design_pattern/internal"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag"
)

type Server struct {
	config config.CredentialDB
	store  *internal.SQLStore
	router *gin.Engine
}

func NewServer(config *config.CredentialDB, store *internal.SQLStore) (*Server, error) {

	server := &Server{
		config: *config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	// later kasih authorization routes

	// controller rest API not method
	router.POST("/users/createTx", server.CreateUser)

	// controller PURCHASE
	router.GET("/purchase/datatable", server.DatatablePurchase)
	router.DELETE("/purchase/delete/:prc_number", server.DeletePurchase)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}

func successResponse(msg string, data any) gin.H {
	return gin.H{
		"message": msg,
		"data":    data,
	}
}
