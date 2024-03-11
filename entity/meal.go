package entity

import (
	"time"
)

type Meal struct {
	ID          uint      `json:"id" gorm:"primary_key;"`
	Name        string    `json:"name" binding:"required"`
	Kalori      uint      `json:"kalori" binding:"required"`
	Protein     uint      `json:"protein" binding:"required"`
	Karbohidrat uint      `json:"karbohidrat" binding:"required"`
	Lemak       uint      `json:"lemak" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
