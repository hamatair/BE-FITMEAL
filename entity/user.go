package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;"`
	UserName    string    `json:"userName" gorm:"type:varchar(255);not null;" binding:"required"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null;unique" binding:"required"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null;" binding:"required"`
	Aktivitas   string    `json:"aktivitas" gorm:"type:varchar(255);not null;" binding:"required"`
	Gender      string    `json:"gender" gorm:"type:varchar(255);not null;" binding:"required"`
	Umur        uint      `json:"umur" gorm:"type:varchar(255);not null;" binding:"required"`
	Alamat      string    `json:"alamat" gorm:"type:varchar(255);not null;" binding:"required"`
	BeratBadan  uint      `json:"beratBadan" binding:"required,number"`
	TinggiBadan uint      `json:"tinggiBadan" binding:"required,number"`
	Kalori      float32   `json:"kalori"`
	Protein     float32   `json:"protein"`
	Karbohidrat float32   `json:"karbohidrat"`
	Lemak       float32   `json:"lemak"`
	PhotoLink   string    `json:"photoLink" gorm:"type:varchar(200);"`
	PhotoName   string    `json:"photoName"`
	Balance     uint      `json:"balance"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
