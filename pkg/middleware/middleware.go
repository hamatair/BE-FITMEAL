package middleware

import (
	"errors"
	"intern-bcc/internal/service"
	"intern-bcc/pkg/jwt"
	"intern-bcc/pkg/response"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

type Interface interface {
	TimeoutMiddleware() gin.HandlerFunc
	AuthenticateUser(ctx *gin.Context)
}

type Middleware struct {
	jwtauth jwt.Interface
	service *service.Service
}

func Init(jwtauth jwt.Interface, service *service.Service) Interface {
	return &Middleware{
		jwtauth: jwtauth,
		service: service,
	}
}

func (m *Middleware) TimeoutMiddleware() gin.HandlerFunc {
	timeOut, _ := strconv.Atoi(os.Getenv("TIME_OUT_LIMIT"))

	return timeout.New(
		timeout.WithTimeout(time.Duration(timeOut)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}

func testResponse(c *gin.Context) {
	response.Error(c, http.StatusRequestTimeout, "Time Out Limit", errors.New(""))
}
