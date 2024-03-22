package mysql

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(entity.User{})
	db.AutoMigrate(entity.Meal{})
	db.AutoMigrate(entity.PasswordValidation{})
	db.AutoMigrate(entity.DailyNutritionUser{})
	db.AutoMigrate(entity.TopUp{})
}
