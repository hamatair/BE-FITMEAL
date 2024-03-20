package repository

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

type MealRepositoryInterface interface {
	FindAll() ([]entity.Meal, error)
	FindAllByName(name string) ([]entity.Meal, error)
	CreateNewDataMeal(user entity.Meal) (entity.Meal, error)
	FindAllByJenis(jenis string) ([]entity.Meal, error)
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

func (m *MealRepository) FindAllByName(name string) ([]entity.Meal, error) {
	var meal []entity.Meal
	err := m.db.Where("name = ?", name).Find(&meal).Error

	return meal, err
}

func (m *MealRepository) FindAllByJenis(jenis string) ([]entity.Meal, error) {
	var meal []entity.Meal
	err := m.db.Where("jenis = ?", jenis).Find(&meal).Error

	return meal, err
}