package api

import (
	"github.com/gin-contrib/cors"
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

	// controller rest API not method
	router.POST("/user/create", server.CreateUser)

	// double layer middleware [/rand/auth]

	// randRoute := router.Group("/rand")
	// randRoute.Use(RandomMiddleware("rand"))

	// randRoute1 := router.Group("/rand1")
	// randRoute1.Use(RandomMiddleware("rand1"))

	// // test double layer middleware grouping
	// authRoute := randRoute.Group("/auth")
	// authRoute.Use(authMiddleware())

	// authRoute1 := randRoute1.Group("/auth")
	// authRoute1.Use(authMiddleware())

	router.POST("/purchase/create", server.CreatePurchase)
	router.GET("/purchase/datatable", server.DatatablePurchase)
	router.DELETE("/purchase/delete/:prc_number", server.DeletePurchase)

	// authRoute1.DELETE("/purchase/delete/:prc_number", server.DeletePurchase)

	server.router = router
}

func (server *Server) SetupCORS() {
	server.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8000"}, // your Swagger UI origin must be allowed
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))
}

func (server *Server) Start(address string) error {
	server.SetupCORS()
	return server.router.Run(address)
}

type ErrorResponse struct {
	Message    string `json:"message" example:"invalid request body"`
	StatusCode int    `json:"status_code" example:"400"`
}

func errorResponse(err error, statusCode int) gin.H {
	return gin.H{
		"message":     err.Error(),
		"status_code": statusCode,
	}
}

type SuccessResponse struct {
	Message    string `json:"message"`
	Data       any    `json:"data"`
	StatusCode int    `json:"status_code" example:"200"`
}

func successResponse(msg string, data any, statusCode int) gin.H {
	return gin.H{
		"message":     msg,
		"data":        data,
		"status_code": statusCode,
	}
}
