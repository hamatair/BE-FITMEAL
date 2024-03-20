package service

import (
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
	"intern-bcc/model"

	"github.com/google/uuid"
)

type MealServiceInterface interface {
	FindAll() ([]entity.Meal, error)
	FindAllByName(name string) ([]entity.Meal, error)
	CreateNewDataMeal(user model.NewMeal) (entity.Meal, error)
	FindAllByJenis(jenis string) ([]entity.Meal, error)
}

type MealService struct {
	mealRepository repository.MealRepositoryInterface
}

func NewMealService(repository repository.MealRepositoryInterface) MealServiceInterface {
	return &MealService{repository}
}

func (m *MealService) CreateNewDataMeal(meal model.NewMeal) (entity.Meal, error) {
	meal.ID = uuid.New()
	nmeal := entity.Meal{
		ID:          meal.ID,
		Name:        meal.Name,
		Jenis:       meal.Jenis,
		Kalori:      meal.Kalori,
		Protein:     meal.Protein,
		Karbohidrat: meal.Karbohidrat,
		Lemak:       meal.Lemak,
	}

	newMeal, err := m.mealRepository.CreateNewDataMeal(nmeal)

	return newMeal, err

}

func (m *MealService) FindAll() ([]entity.Meal, error) {
	meal, err := m.mealRepository.FindAll()

	return meal, err
}

func (m *MealService) FindAllByName(name string) ([]entity.Meal, error) {
	meal, err := m.mealRepository.FindAllByName(name)

	return meal, err
}

func (m *MealService) FindAllByJenis(jenis string) ([]entity.Meal, error) {
	meal, err := m.mealRepository.FindAllByJenis(jenis)

	return meal, err
}
