package model

import "github.com/google/uuid"

type NewMeal struct {
	ID          uuid.UUID `json:"-"`
	Name        string    `json:"name"`
	Jenis       string    `json:"jenis"`
	Kalori      float32   `json:"kalori"`
	Protein     float32   `json:"protein"`
	Karbohidrat float32   `json:"karbohidrat"`
	Lemak       float32   `json:"lemak"`
}

type cariMeal struct {
	ID    uuid.UUID `json:"-"`
	Name  string    `json:"-"`
	Jenis string    `json:"-"`
}
