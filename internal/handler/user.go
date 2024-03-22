package handler

import (
	"errors"
	"intern-bcc/entity"
	"intern-bcc/model"
	"intern-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		return
	}

	param := model.EditProfile{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", err)
		return
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
		return
	}
	err = h.Service.UserService.CreateCodeVerification(newCode)

	response.Success(c, http.StatusOK, "success to send verification code", err)
}

func (h *Handler) ForgotPasswordUser(c *gin.Context) {
	var checkCode model.ForgotPassword

	err := c.ShouldBindJSON(&checkCode)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", err)
		return
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
		return
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
		return
	}

	response.Success(c, http.StatusOK, "success to tambah nutrisi", nil)
}

func (h *Handler) UploadPhoto(c *gin.Context) {
	photo, err := c.FormFile("photo")
	if err != nil {
		response.Success(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = h.Service.UserService.UploadPhoto(c, model.UserUploadPhoto{Photo: photo})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to upload photo", err)
		return
	}

	response.Success(c, http.StatusOK, "success upload photo", nil)
}

func (h *Handler) PaketMakan(c *gin.Context) {
	var paket model.PaketMakan
	err := c.ShouldBindJSON(&paket)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusNotFound, "failed to get data user", errors.New(""))
		return
	}

	err = h.Service.UserService.CreatePaket(paket, user.(entity.User).ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to create paket", err)
		return
	}

	response.Success(c, http.StatusOK, "success to create paket", nil)
}

func (h *Handler) GetAllDataPaketByUserId(c *gin.Context) {
	var findData []entity.PaketMakan

	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusNotFound, "failed to get data user", errors.New(""))
		return
	}

	findData, err := h.Service.UserService.FindAllPaketByUserId(user.(entity.User).ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "fail to get data", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get data", findData)
}