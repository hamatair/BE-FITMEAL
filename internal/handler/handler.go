package handler

import (
	"fmt"
	"intern-bcc/internal/service"
	"intern-bcc/pkg/middleware"
	"os"

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
	}))
	h.Router.Use(h.Middleware.TimeoutMiddleware())

	v1 := h.Router.Group("/v1")

	v1.GET("/user/get", h.GetAllDataUser)
	v1.GET("/user/get-user-profile",h.Middleware.AuthenticateUser, h.GetUserById)
	v1.POST("/meal", h.NewDataMeal)
	v1.GET("/meal/get", h.GetAllDataMeal)
	v1.POST("/user/register", h.UserRegisterAndPersonalization)
	v1.PATCH("/user/edit-profile",h.Middleware.AuthenticateUser, h.UserEditProfile)
	v1.PATCH("/user/edit-profile/change-password",h.Middleware.AuthenticateUser, h.changePasswordUser)
	v1.POST("/user/login", h.Login)
	v1.POST("/user/login-user", h.Middleware.AuthenticateUser, h.getLoginUser)
	v1.POST("/user/forgot-password/get", h.CreateCodeVerification)
	v1.POST("/user/forgot-password", h.ForgotPasswordUser)
	v1.POST("/user/forgot-password/change-password", h.ChangePasswordBeforeLogin)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	h.Router.Run(fmt.Sprintf(":%s", port))
}
