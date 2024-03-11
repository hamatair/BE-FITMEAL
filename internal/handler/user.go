package handler

import (
	"intern-bcc/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *Handler) GetAllDataUser(c *gin.Context) {
	var findData []entity.User

	findData, err := u.Service.UserService.FindAll()
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": findData,
	})
}

func (u *Handler) NewDataUser(c *gin.Context) {
	var newuser entity.NewUser

	c.ShouldBindJSON(&newuser)

	newUser, err := u.Service.UserService.Create(newuser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "Data user berhasil ditambahkan",
		"data":    newUser,
	})
}

