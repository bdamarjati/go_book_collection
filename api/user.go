package api

import (
	"database/sql"
	"net/http"

	"github.com/bdamarjati/go_book_collection/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := sqlc.CreateUserParams{
		Username: req.Username,
		Role:     sql.NullString{String: req.Role, Valid: true},
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}
