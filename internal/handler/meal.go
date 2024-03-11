package handler

import (
	"intern-bcc/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *Handler) GetAllDataMeal(c *gin.Context) {
	var findData []entity.Meal

	findData, err := u.Service.MealService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": findData,
	})
}

func (u *Handler) NewDataMeal(c *gin.Context) {
	var newmeal entity.NewMeal

	c.ShouldBindJSON(&newmeal)

	newMeal, err := u.Service.MealService.Create(newmeal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "Data meal berhasil ditambahkan",
		"data":    newMeal,
	})
}
