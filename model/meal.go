package model

import "github.com/google/uuid"

type NewMeal struct {
	ID          uuid.UUID `json:"-"`
	Name        string    `json:"name" binding:"required"`
	Kalori      uint      `json:"kalori" binding:"required"`
	Protein     uint      `json:"protein" binding:"required"`
	Karbohidrat uint      `json:"karbohidrat" binding:"required"`
	Lemak       uint      `json:"lemak" binding:"required"`
}
