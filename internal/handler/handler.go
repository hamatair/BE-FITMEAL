package handler

import (
	"intern-bcc/internal/service"
	"intern-bcc/pkg/middleware"

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
	h.Router.Use(h.Middleware.TimeoutMiddleware())
	v1 := h.Router.Group("/v1")

	// v1.POST("/user", h.NewDataUser)
	v1.GET("/user", h.GetAllDataUser)
	v1.POST("/meal", h.NewDataMeal)
	v1.GET("/meal", h.GetAllDataMeal)
	v1.POST("/user", h.UserRegister)
	v1.PATCH("/tes/:name", h.UserPersonalization)

	h.Router.Run(":5000")
}
