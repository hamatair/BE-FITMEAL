package handler

import (
	"errors"
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

func (h *Handler) UserRegisterAndPersonalization(c *gin.Context) {
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

	response.Success(c, http.StatusCreated, "success create new user", newUser)
}

func (h *Handler) UserEditProfile(c *gin.Context) {
	id := c.Param("id")

	param := model.EditProfile{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", err)
		c.Next()
	}

	user, err := h.Service.UserService.UserEditProfile(param, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to personalize data user", err)
		return
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

	response.Success(ctx, http.StatusOK, "success login", token)
}

func (h *Handler) getLoginUser(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusInternalServerError, "failed get login user", errors.New(""))
		return
	}

	response.Success(c, http.StatusOK, "get login user", user.(entity.User))
}

func (h *Handler) changePasswordUser(c *gin.Context) {
	id := c.Param("id")

	param := model.ChangePassword{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	user, err := h.Service.UserService.UserChangePassword(param, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to change password", err)
		return
	}

	response.Success(c, http.StatusAccepted, "success to change password", user)
}
