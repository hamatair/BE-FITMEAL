package handler

import (
	"intern-bcc/entity"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *Handler) GetAllDataMeal(c *gin.Context) {
	var findData []entity.Meal

	findData, err := u.Service.MealService.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "fail to get data", err)
	}

	response.Success(c, http.StatusOK, "success to get data", findData)
}

func (u *Handler) NewDataMeal(c *gin.Context) {
	var newmeal model.NewMeal

	c.ShouldBindJSON(&newmeal)

	newMeal, err := u.Service.MealService.Create(newmeal)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "fail to make a new meal data", err)
	}

	response.Success(c, http.StatusAccepted, "success to make a new meal data", newMeal)
}
