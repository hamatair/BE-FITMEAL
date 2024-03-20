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
		response.Error(c, http.StatusInternalServerError, "failed to get all data user", err)
	}

	response.Success(c, http.StatusOK, "success to get all data user", findData)
}

func (u *Handler) GetUserById(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusInternalServerError, "failed get login user", errors.New(""))
		return
	}

	response.Success(c, http.StatusOK, "success to get user data", user.(entity.User))
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
	dataUser, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusNotFound, "failed to get data user", errors.New(""))
	}

	param := model.EditProfile{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", err)
		c.Next()
	}

	user, err := h.Service.UserService.UserEditProfile(param, dataUser.(entity.User).ID.String())
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
	dataUser, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusInternalServerError, "failed get login user", errors.New(""))
		return
	}

	param := model.ChangePassword{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	user, err := h.Service.UserService.UserChangePassword(param, string(dataUser.(entity.User).ID.String()))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to change password", err)
		return
	}

	response.Success(c, http.StatusAccepted, "success to change password", user)
}

func (h *Handler) CreateCodeVerification(c *gin.Context) {
	newCode := model.ForgotPassword{}

	err := c.ShouldBindJSON(&newCode)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", err)
		c.Next()
	}
	err = h.Service.UserService.CreateCodeVerification(newCode)

	response.Success(c, http.StatusOK, "success to send verification code", err)
}

func (h *Handler) ForgotPasswordUser(c *gin.Context) {
	var checkCode model.ForgotPassword

	err := c.ShouldBindJSON(&checkCode)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", err)
	}

	err = h.Service.UserService.CheckCode(checkCode)
	if err != nil {
		response.Error(c, http.StatusNotFound, "failed to validating code", err)
		return
	}

	response.Success(c, http.StatusOK, "success to validating code", checkCode.Email)
}

func (h *Handler) ChangePasswordBeforeLogin(c *gin.Context) {
	var getData model.ChangePasswordBeforeLogin

	err := c.ShouldBindJSON(&getData)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = h.Service.UserService.ChangePasswordBeforeLogin(getData)
	if err != nil {
		response.Error(c, http.StatusNotFound, "failed to get data", err)
		return
	}

	response.Success(c, http.StatusCreated, "success to change password", err)
}

func (h *Handler) DailyNutrition(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusInternalServerError, "failed get data user", errors.New(""))
		return
	}

	data, err := h.Service.UserService.GetDailyNutrition(user.(entity.User).ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get data daily nutrition", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get data daily nutrition", data)
}

func (h *Handler) TambahNutrisi(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusNotFound, "failed to get data user", errors.New(""))
	}
	var tambah model.TambahNutrisi
	err := c.ShouldBindJSON(&tambah)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = h.Service.UserService.TambahNutrisi(user.(entity.User).ID, tambah)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to tambah nutrisi", err)
	}

	response.Success(c, http.StatusOK, "success to tambah nutrisi", nil)
}
