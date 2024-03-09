package handler

import (
	"intern-bcc/entity"
	"intern-bcc/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userhandler service.Service
}

func NewUserHandler(userHandler service.Service) *UserHandler {
	return &UserHandler{userHandler}
}

func EndPoint(handler *UserHandler) {
	r := gin.Default()

	v1 := r.Group("/v1")

	v1.POST("/user", handler.NewSetDataUser)

	r.Run(":5000")
}

func (u *UserHandler) NewSetDataUser(c *gin.Context) {
	var newUserHandler entity.NewUser

	c.ShouldBindJSON(&newUserHandler)

	newUser, err := u.userhandler.Create(newUserHandler)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "Data user berhasil ditambahkan",
		"data":    newUser,
	})
}

