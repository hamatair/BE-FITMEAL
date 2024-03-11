package entity

type User struct {
	ID          uint   `json:"id" gorm:"primary_key;"`
	Name        string `json:"name" gorm:"type:varchar(255);not null;" binding:"required"`
	Email       string `json:"email" gorm:"type:varchar(255);not null;unique" binding:"required"`
	Password    string `json:"password" gorm:"type:varchar(255);not null;" binding:"required"`
	Umur        uint   `json:"umur" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	TinggiBadan uint   `json:"tinggibadan" binding:"required"`
	BeratBadan  uint   `json:"beratbadan" binding:"required"`
	Aktivitas   uint   `json:"aktivitas" binding:"required"`
}

type NewUser struct {
	Name     string `json:"name" gorm:"type:varchar(255);not null;" binding:"required"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(255);not null;" binding:"required"`
}
