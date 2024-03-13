package handler

import (
	"intern-bcc/entity"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *Handler) GetAllDataUser(c *gin.Context) {
	var findData []entity.User

	findData, err := u.Service.UserService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": findData,
	})
}

func (h *Handler) UserRegister(c *gin.Context) {
	param := model.Register{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	newUser, err := h.Service.UserService.Create(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to register new user", err)
		return
	}

	response.Success(c, http.StatusCreated, "success register new user", newUser)
}

func (h *Handler) UserPersonalization(c *gin.Context) {
	str := c.Param("name")

	param := model.Personalization{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", err)
	}

	user, err := h.Service.UserService.UserPersonalization(param, str)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to personalize data user", err)
	}

	response.Success(c, http.StatusAccepted, "success to personalization", user)

}

func (h *Handler) Login(ctx *gin.Context) {
	param := model.Login{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	token, err := h.Service.UserService.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to login", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success login to system", token)
}

