package entity

type User struct {
	ID          uint   `json:"id" gorm:"primary_key;"`
	Name        string `json:"name" gorm:"type:varchar(255);not null;" binding:"required"`
	Email       string `json:"email" gorm:"type:varchar(255);not null;unique" binding:"required"`
	Password    string `json:"password" gorm:"type:varchar(255);not null;" binding:"required"`
	Umur        uint   `json:"umur"`
	Gender      string `json:"gender"`
	TinggiBadan uint   `json:"tinggibadan"`
	BeratBadan  uint   `json:"beratbadan"`
	Aktivitas   uint   `json:"aktivitas"`
}

type NewUser struct {
	Name     string `json:"name" gorm:"type:varchar(255);not null;" binding:"required"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(255);not null;" binding:"required"`
	Umur     uint   `json:"umur" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
}
