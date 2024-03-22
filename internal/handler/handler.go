package handler

import (
	"fmt"
	"intern-bcc/internal/service"
	"intern-bcc/pkg/middleware"
	"log"

	"os"
	"time"

	"github.com/gin-contrib/cors"
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
	h.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	h.Router.Use(h.Middleware.TimeoutMiddleware())

	go h.CheckTime()

	v1 := h.Router.Group("/v1")

	v1.GET("/user/get-user-profile", h.Middleware.AuthenticateUser, h.GetUserById)
	v1.POST("/user/register", h.UserRegisterAndPersonalization)
	v1.PATCH("/user/edit-profile", h.Middleware.AuthenticateUser, h.UserEditProfile)
	v1.PATCH("/user/edit-profile/change-password", h.Middleware.AuthenticateUser, h.changePasswordUser)
	v1.POST("/user/login", h.Login)
	v1.POST("/user/login-user", h.Middleware.AuthenticateUser, h.getLoginUser)
	v1.POST("/user/forgot-password/get", h.CreateCodeVerification)
	v1.POST("/user/forgot-password", h.ForgotPasswordUser)
	v1.POST("/user/forgot-password/change-password", h.ChangePasswordBeforeLogin)

	v1.POST("/user/upload-photo", h.Middleware.AuthenticateUser, h.UploadPhoto)

	v1.POST("/user/top-up", h.Middleware.AuthenticateUser, h.TopUp)
	v1.POST("/user/top-up/status", h.VerifyPayment)

	v1.GET("/user/daily-nutrition", h.Middleware.AuthenticateUser, h.DailyNutrition)
	v1.POST("/user/tambah-nutrisi", h.Middleware.AuthenticateUser, h.TambahNutrisi)
	v1.POST("/user/tambah-paket", h.Middleware.AuthenticateUser, h.Middleware.PremiumAccess, h.PaketMakan)
	v1.GET("/user/paket/get", h.Middleware.AuthenticateUser, h.Middleware.PremiumAccess, h.GetAllDataPaketByUserId)

	v1.POST("/meal", h.NewDataMeal)
	v1.GET("/meal/get", h.GetAllDataMeal)
	v1.GET("/meal/jenis/get/:jenis", h.GetAllDataMealByJenis)
	v1.GET("/meal/get/:name", h.GetAllDataMealByName)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	h.Router.Run(fmt.Sprintf(":%s", port))
}

func (h *Handler) CheckTime() {
	for {
		now := time.Now()

		if now.Hour() == 0 && now.Minute() == 0 && now.Second() == 0 {
			h.ResetDataDailyNutrition()
			time.Sleep(5 * time.Second)
		}

		time.Sleep(1 * time.Second)
	}
}

func (h *Handler) ResetDataDailyNutrition() {
	err := h.Service.UserService.ResetDataDailyNutrition()
	if err != nil {
		log.Fatal()
	}
}
