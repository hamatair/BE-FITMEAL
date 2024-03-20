package model

import "github.com/google/uuid"

type NewMeal struct {
	ID          uuid.UUID `json:"-"`
	Name        string    `json:"name"`
	Jenis       string    `json:"jenis"`
	Kalori      uint      `json:"kalori"`
	Protein     uint      `json:"protein"`
	Karbohidrat uint      `json:"karbohidrat"`
	Lemak       uint      `json:"lemak"`
}

type cariMeal struct {
	ID    uuid.UUID `json:"-"`
	Name  string    `json:"-"`
	Jenis string    `json:"-"`
}
