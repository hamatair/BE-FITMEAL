package entity

import (
	"time"

	"github.com/google/uuid"
)

type Meal struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;"`
	Name        string    `json:"name" binding:"required"`
	Jenis       string    `json:"jenis" binding:"required"`
	Kalori      float32      `json:"kalori" binding:"required"`
	Protein     float32      `json:"protein" binding:"required"`
	Karbohidrat float32      `json:"karbohidrat" binding:"required"`
	Lemak       float32      `json:"lemak" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
