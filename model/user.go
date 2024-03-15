package model

import "github.com/google/uuid"

type Register struct {
	ID          uuid.UUID `json:"-"`
	UserName    string    `json:"userName" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required,min=8"`
	Aktivitas   string    `json:"aktivitas" gorm:"type:varchar(255);not null;" binding:"required"`
	Gender      string    `json:"gender" gorm:"type:varchar(255);not null;" binding:"required"`
	Umur        uint      `json:"umur" gorm:"type:varchar(255);not null;" binding:"required"`
	Alamat      string    `json:"alamat" gorm:"type:varchar(255);not null;" binding:"required"`
	BeratBadan  uint      `json:"beratBadan" binding:"required,number"`
	TinggiBadan uint      `json:"tinggiBadan" binding:"required,number"`
}

type EditProfile struct {
	UserName    string `json:"userName" binding:"required"`
	Umur        uint   `json:"umur" gorm:"not null;" binding:"required"`
	Alamat      string `json:"alamat" gorm:"type:varchar(255);not null;" binding:"required"`
	BeratBadan  uint   `json:"beratBadan" gorm:"not null;" binding:"required,number"`
	TinggiBadan uint   `json:"tinggiBadan" gorm:"not null;" binding:"required,number"`
}

type ChangePassword struct {
	OldPassword     string `json:"oldPassword" binding:"required,min=8"`
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=8"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
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
