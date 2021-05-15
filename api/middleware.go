package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thanhqt2002/hackathon/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		authorization := strings.ToLower(fields[0])

		if authorization != authorizationTypeBearer {
			err := fmt.Errorf("%v is not an supported authorization type", authorization)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		AccessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(AccessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
