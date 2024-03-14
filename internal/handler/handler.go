package handler

import (
	"fmt"
	"intern-bcc/internal/service"
	"intern-bcc/pkg/middleware"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service    *service.Service
	Router     *gin.Engine
	Middleware middleware.Interface
}

func NewHandler(service *service.Service, middleware middleware.Interface) *Handler {
	return &Handler{
		Service:    service,
		Router:     gin.Default(),
		Middleware: middleware,
	}
}

func (h *Handler) EndPoint() {
	h.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	
		MaxAge: 12 * time.Hour,
	  }))
	h.Router.Use(h.Middleware.TimeoutMiddleware())
	
	v1 := h.Router.Group("/v1")

	v1.GET("/user", h.GetAllDataUser)
	v1.POST("/meal", h.NewDataMeal)
	v1.GET("/meal", h.GetAllDataMeal)
	v1.POST("/user/register", h.UserRegisterAndPersonalization)
	v1.PATCH("/tes/:name", h.UserEditProfile)
	v1.POST("user/login", h.Login)
	v1.POST("user/login-user", h.Middleware.AuthenticateUser, h.getLoginUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	h.Router.Run(fmt.Sprintf(":%s", port))
}
