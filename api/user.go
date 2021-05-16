package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/thanhqt2002/hackathon/db/sqlc"
	"github.com/thanhqt2002/hackathon/db/util"
	"github.com/thanhqt2002/hackathon/token"
)

type CreateUserRequest struct {
	Username string `json:"Username" binding:"required,alphanum"`
	Fullname string `json:"Fullname" binding:"required"`
	Password string `json:"Password" binding:"required,min=6"`
	Usertype int32  `json:"Usertype" binding:"required,min=1,max=3"`
}

type UserResponse struct {
	Username string `json:"Username" binding:"required,alphanum"`
	Fullname string `json:"Fullname" binding:"required"`
	Usertype int32  `json:"Usertype" binding:"required,min=1,max=3"`
}

func NewUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Fullname: user.Fullname,
		Usertype: user.Usertype,
	}
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	authorized_user, err := server.store.GetUser(ctx, authPayload.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if authorized_user.Usertype <= req.Usertype {
		err := errors.New("this user is not allowed to add new user of role higher or equal them him/her-self")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	PasswordEncoded, err := util.HassPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:        req.Username,
		Fullname:        req.Fullname,
		PasswordEncoded: PasswordEncoded,
		Usertype:        req.Usertype,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, NewUserResponse(user))
}

type LoginUserRequest struct {
	Username string `json:"Username" binding:"required,alphanum"`
	Password string `json:"Password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token" binding:"access_token"`
	User        UserResponse
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.PasswordEncoded)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	AccessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.TokenAccessDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, LoginUserResponse{
		User:        NewUserResponse(user),
		AccessToken: AccessToken,
	})
}
