package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository UserRepositoryInterface
	MealRepository MealRepositoryInterface
}

func NewRepository(db *gorm.DB) *Repository{
	return &Repository{
		UserRepository: NewUserRepository(db),
		MealRepository: NewMealRepository(db),
	}
}