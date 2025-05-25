package api

import (
	"context"
	_ "database/sql"
	"errors"
	_ "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luke_design_pattern/db"
)

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

func (server *Server) CreateUser(c *gin.Context) {

	ctx := c.Request.Context()
	ctx, cancelFunc := context.WithTimeout(ctx, 100*time.Millisecond)

	defer cancelFunc()

	// chTimeOut := make(chan string)

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userParams := db.CreateUserParams{
		Username:  req.Username,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Password:  req.Password,
	}

	args := db.CreateUserTxParams{
		CreateUserParams: userParams,
		AfterCreate: func(user db.Users) error {
			return nil
		},
	}

	users, err := server.store.CreateUserTx(ctx, args)
	if err != nil {
		// catch err from db driver
		// catch err from defined deadline context
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, users)

}
