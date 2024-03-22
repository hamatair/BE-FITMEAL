package service

import (
	"intern-bcc/internal/repository"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/jwt"
	"intern-bcc/pkg/midtranss"
	"intern-bcc/pkg/supabase"
)

type Service struct {
	UserService  UserServiceInterface
	MealService  MealServiceInterface
	TopUpService TopUpServiceI
}

type InitParam struct {
	Repository *repository.Repository
	JwtAuth    jwt.Interface
	Bcrypt     bcrypt.Interface
	Supabase   supabase.Interface
	TopUp      repository.TopUpRepositoryI
	Midtrans   midtranss.MidtransServiceI
}

func NewService(param InitParam) *Service {
	return &Service{
		UserService: NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth, param.Supabase),
		MealService: NewMealService(param.Repository.MealRepository),
		TopUpService: NewTopUp(param.TopUp, param.Midtrans, param.Repository),
	}
}
