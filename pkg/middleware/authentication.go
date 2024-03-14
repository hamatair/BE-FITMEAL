package middleware

import (
	"errors"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "empty token", errors.New(""))
		ctx.Abort()
	}

	token := strings.Split(bearer, " ")[1]
	userId, err := m.jwtauth.ValidateToken(token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed validate token", err)
		ctx.Abort()
	}
	
	user, err := m.service.UserService.GetUser(model.UserParam{
		ID: userId,
	})
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed get user", err)
		ctx.Abort()
	}

	ctx.Set("user", user)

	ctx.Next()
}
