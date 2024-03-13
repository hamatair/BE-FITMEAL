package service

import (
	"intern-bcc/internal/repository"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/jwt"
)

type Service struct {
	UserService UserServiceInterface
	MealService MealServiceInterface
}

type InitParam struct {
	Repository *repository.Repository
	JwtAuth    jwt.Interface
	Bcrypt     bcrypt.Interface
}

func NewService(param InitParam) *Service{
	return &Service{
		UserService: NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth),
		MealService: NewMealService(param.Repository.MealRepository),
	}
}
