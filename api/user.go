package api

import (
	"context"
	_ "database/sql"
	"errors"
	_ "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luke_design_pattern/internal"
)

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}
type CreateUserResponse struct {
	Message    string                      `json:"message"`
	Data       internal.CreateUserTxResult `json:"data"`
	StatusCode int                         `json:"status_code"`
}

// GetUser godoc
// @Summary      Create New User of App
// @Description  Creating New User which never existed before
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body CreateUserRequest true 	"request for creating new user"
// @Success      200 {object} CreateUserResponse 		"reponse if success creating new user"
// @Failure 	 400 {object} ErrorResponse "Bad Request"
// @Router       /user/create [post]
func (server *Server) CreateUser(c *gin.Context) {

	ctx := c.Request.Context()
	ctx, cancelFunc := context.WithTimeout(ctx, 100*time.Millisecond)

	defer cancelFunc()

	// chTimeOut := make(chan string)

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	userParams := internal.CreateUserParams{
		Username:  req.Username,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Password:  req.Password,
	}

	args := internal.CreateUserTxParams{
		CreateUserParams: userParams,
		AfterCreate: func(user internal.Users) error {
			return nil
		},
	}

	// users, err := server.store.CreateUserTx(ctx, args)
	users, err := server.store.CreateUserTx(ctx, args)
	if err != nil {
		// catch err from db driver
		// catch err from defined deadline context
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, errorResponse(err, http.StatusGatewayTimeout))
			return
		}
		c.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	respSuccess := successResponse(
		"success create new user",
		internal.CreateUserTxResult{
			// embedded struct
			Users: internal.Users{
				Username:  users.Username,
				Email:     users.Email,
				Firstname: users.Firstname,
				Lastname:  users.Lastname,
			},
		},
		http.StatusAccepted,
	)
	c.JSON(http.StatusOK, respSuccess)

}
