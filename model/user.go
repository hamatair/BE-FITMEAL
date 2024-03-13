package model

import "github.com/google/uuid"

type Register struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
}

type Personalization struct {
	Aktivitas   string `json:"aktivitas" gorm:"type:varchar(255);not null;" binding:"required"`
	Gender      string `json:"gender" gorm:"type:varchar(255);not null;" binding:"required"`
	Umur        uint   `json:"umur" gorm:"not null;" binding:"required"`
	Alamat      string `json:"alamat" gorm:"type:varchar(255);not null;" binding:"required"`
	BeratBadan  uint   `json:"beratbadan" gorm:"not null;" binding:"required,number"`
	TinggiBadan uint   `json:"tinggibadan" gorm:"not null;" binding:"required,number"`
}
