package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RandomMiddleware(param string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		number := 100
		log.Printf("check middleware log %d", number)
		log.Printf("check ctx %v", param)
	}
}

const (
	authorizationHeaderKey = "Authorization"
)

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("auth header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err, http.StatusUnauthorized))
			return
		}

		log.Printf("auth header ==> %v", authorizationHeader)

		ctx.Next()
	}
}
