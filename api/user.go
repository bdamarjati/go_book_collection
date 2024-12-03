package api

import (
	"database/sql"
	"net/http"
	"time"

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
		Username: sql.NullString{String: req.Username, Valid: true},
		Role:     sql.NullString{String: req.Role, Valid: true},
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getUserResponse struct {
	ID        int32     `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := getUserResponse{
		ID:        user.ID,
		Username:  user.Username.String,
		Role:      user.Role.String,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, res)
}
