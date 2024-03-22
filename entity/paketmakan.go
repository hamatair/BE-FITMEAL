package entity

import "github.com/google/uuid"

type PaketMakan struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;"`
	UserID      uuid.UUID `json:"userId" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Kalori      float32   `json:"kalori" binding:"required"`
	Protein     float32   `json:"protein" binding:"required"`
	Karbohidrat float32   `json:"karbohidrat" binding:"required"`
	Lemak       float32   `json:"lemak" binding:"required"`
}
