package handler

import (
	"intern-bcc/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func EndPoint(handler *Handler) {
	r := gin.Default()

	v1 := r.Group("/v1")

	v1.POST("/user", handler.NewDataUser)
	v1.GET("/user", handler.GetAllDataUser)
	v1.POST("/meal", handler.NewDataMeal)
	v1.GET("/meal", handler.GetAllDataMeal)

	r.Run(":5000")
}
