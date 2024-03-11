package model

type NewUser struct {
	Name     string `json:"name" gorm:"type:varchar(255);not null;" binding:"required"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(255);not null;" binding:"required"`
}

