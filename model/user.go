package model

import (
	"time"

	"github.com/google/uuid"
)

type Register struct {
	ID          uuid.UUID `json:"-"`
	UserName    string    `json:"userName" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required"`
	Aktivitas   string    `json:"aktivitas" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	Umur        uint      `json:"umur" binding:"required,number"`
	Alamat      string    `json:"alamat" binding:"required"`
	BeratBadan  uint      `json:"beratBadan" binding:"required,number"`
	TinggiBadan uint      `json:"tinggiBadan" binding:"required,number"`
}

type EditProfile struct {
	UserName    string `json:"userName"`
	Umur        uint   `json:"umur"`
	Alamat      string `json:"alamat"`
	BeratBadan  uint   `json:"beratBadan"`
	TinggiBadan uint   `json:"tinggiBadan"`
}

type ChangePassword struct {
	OldPassword     string `json:"oldPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type ForgotPassword struct {
	Email       string `json:"email" binding:"required,email"`
	Kode        int    `json:"kode"`
	ExpiredTime time.Time
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserParam struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ChangePasswordBeforeLogin struct {
	Email           string `json:"email" binding:"required,email"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type TambahNutrisi struct {
	Kalori      float32 `json:"kalori"`
	Protein     float32 `json:"protein"`
	Karbohidrat float32 `json:"karbohidrat"`
	Lemak       float32 `json:"lemak"`
}
