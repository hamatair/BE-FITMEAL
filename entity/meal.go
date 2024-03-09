package entity

type Meal struct {
	ID          uint   `json:"id" gorm:"primary_key;"`
	Name        string `json:"name" binding:"required"`
	Kalori      uint   `json:"kalori"`
	Protein     uint   `json:"protein"`
	Karbohidrat uint   `json:"karbohidrat"`
	Lemak       uint   `json:"lemak"`
}

type NewMeal struct {
	Name        string `json:"name" binding:"required"`
	Kalori      uint   `json:"kalori" binding:"required"`
	Protein     uint   `json:"protein" binding:"required"`
	Karbohidrat uint   `json:"karbohidrat" binding:"required"`
	Lemak       uint   `json:"lemak" binding:"required"`
}
