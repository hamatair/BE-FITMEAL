package service

import "intern-bcc/internal/repository"

type Service struct {
	UserService UserServiceInterface
	MealService MealServiceInterface
}

func NewService(repository *repository.Repository) *Service{
	return &Service{
		UserService: NewUserService(repository.UserRepository),
		MealService: NewMealService(repository.MealRepository),
	}
}
