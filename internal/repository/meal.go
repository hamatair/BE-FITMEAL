package repository

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

type MealRepositoryInterface interface {
	FindAll() ([]entity.Meal, error)
	FindByID(ID int) (entity.Meal, error)
	Create(user entity.Meal) (entity.Meal, error)
}

type MealRepository struct {
	db *gorm.DB
}

func NewMealRepository(db *gorm.DB) MealRepositoryInterface {
	return &MealRepository{
		db: db,
	}
}

func (m *MealRepository) Create(meal entity.Meal) (entity.Meal, error) {
	err := m.db.Create(&meal).Error

	return meal, err
}

func (m *MealRepository) FindAll() ([]entity.Meal, error) {
	var meal []entity.Meal
	err := m.db.Find(&meal).Error

	return meal, err
}

func (m *MealRepository) FindByID(ID int) (entity.Meal, error) {
	var meal entity.Meal
	err := m.db.Where("id = ?", ID).First(&meal).Error

	return meal, err
}
