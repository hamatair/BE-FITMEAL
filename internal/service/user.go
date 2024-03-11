package service

import (
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
)

type UserServiceInterface interface {
	FindAll() ([]entity.User, error)
	FindByID(ID int) (entity.User, error)
	Create(user entity.NewUser) (entity.User, error)
}

type UserService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{repository}
}

func (m *UserService) Create(user entity.NewUser) (entity.User, error) {
	nuser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	newUser, err := m.userRepository.Create(nuser)

	return newUser, err

}

func (m *UserService) FindAll() ([]entity.User, error) {
	user, err := m.userRepository.FindAll()

	return user, err
}

func (m *UserService) FindByID(ID int) (entity.User, error) {
	user, err := m.userRepository.FindByID(ID)

	return user, err
}
