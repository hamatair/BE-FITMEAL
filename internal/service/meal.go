package service

import (
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
	"intern-bcc/model"
)

type MealServiceInterface interface {
	FindAll() ([]entity.Meal, error)
	FindByName(name string) (entity.Meal, error)
	Create(user model.NewMeal) (entity.Meal, error)
}

type MealService struct {
	mealRepository repository.MealRepositoryInterface
}

func NewMealService(repository repository.MealRepositoryInterface) MealServiceInterface {
	return &MealService{repository}
}

func (m *MealService) Create(meal model.NewMeal) (entity.Meal, error) {
	nmeal := entity.Meal{
		Name:        meal.Name,
		Kalori:      meal.Kalori,
		Protein:     meal.Protein,
		Karbohidrat: meal.Karbohidrat,
		Lemak:       meal.Lemak,
	}

	newMeal, err := m.mealRepository.Create(nmeal)

	return newMeal, err

}

func (m *MealService) FindAll() ([]entity.Meal, error) {
	meal, err := m.mealRepository.FindAll()

	return meal, err
}

func (m *MealService) FindByName(name string) (entity.Meal, error) {
	meal, err := m.mealRepository.FindByName(name)

	return meal, err
}
