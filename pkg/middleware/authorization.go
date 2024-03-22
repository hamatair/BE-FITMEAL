package middleware

import (
	"errors"
	"intern-bcc/entity"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) PremiumAccess(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusForbidden, "failed to authorize user", errors.New(""))
		c.Abort()
	}

	if user.(entity.User).Role != 1 {
		response.Error(c, http.StatusForbidden, "failed to let user", errors.New("user don't have access"))
		c.Abort()
	}

	c.Next()
}
