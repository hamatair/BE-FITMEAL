package repository

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

type MealRepositoryInterface interface {
	FindAll() ([]entity.Meal, error)
	FindByName(name string) (entity.Meal, error)
	CreateNewDataMeal(user entity.Meal) (entity.Meal, error)
}

type MealRepository struct {
	db *gorm.DB
}

func NewMealRepository(db *gorm.DB) MealRepositoryInterface {
	return &MealRepository{
		db: db,
	}
}

func (m *MealRepository) CreateNewDataMeal(meal entity.Meal) (entity.Meal, error) {
	err := m.db.Create(&meal).Error

	return meal, err
}

func (m *MealRepository) FindAll() ([]entity.Meal, error) {
	var meal []entity.Meal
	err := m.db.Find(&meal).Error

	return meal, err
}

func (m *MealRepository) FindByName(name string) (entity.Meal, error) {
	var meal entity.Meal
	err := m.db.Where("name = ?", name).First(&meal).Error

	return meal, err
}
