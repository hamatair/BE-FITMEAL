package entity

import "github.com/google/uuid"

type DailyNutritionUser struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;"`
	Email       string    `json:"email" binding:"required,email"`
	Kalori      float32   `json:"kalori"`
	Protein     float32   `json:"protein"`
	Karbohidrat float32   `json:"karbohidrat"`
	Lemak       float32   `json:"lemak"`
}
