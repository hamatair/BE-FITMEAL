package middleware

import (
	"errors"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthenticateUser(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		response.Error(c, http.StatusUnauthorized, "empty token", errors.New(""))
		c.Abort()
	}

	token := strings.Split(bearer, " ")[1]
	userId, err := m.jwtauth.ValidateToken(token)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "failed validate token", err)
		c.Abort()
	}
	
	user, err := m.service.UserService.GetUser(model.UserParam{
		ID: userId,
	})
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "failed get user", err)
		c.Abort()
	}

	c.Set("user", user)

	c.Next()
}
