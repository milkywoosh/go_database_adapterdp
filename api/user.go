package api

import (
	_ "database/sql"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luke_design_pattern/db"
)

type createUserRequest struct {
	Username, Email, Firstname, Lastname, Password string
}

func (server *Server) CreateUser(ctx *gin.Context) {

	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
	return
}
