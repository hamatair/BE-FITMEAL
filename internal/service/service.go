package service

import (
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
)

type Service interface {
	FindAll() ([]entity.User, error)
	FindByID(ID int) (entity.User, error)
	Create(newuser entity.NewUser) (entity.User, error)
}

type service struct {
	repository repository.Repositry
}

func NewService(repository repository.Repositry) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.User, error) {
	user, err := s.repository.FindAll()
	return user, err
}

func (s *service) FindByID(ID int) (entity.User, error) {
	user, err := s.repository.FindByID(ID)
	return user, err
}

func (s *service) Create(newuser entity.NewUser) (entity.User, error) {
	user := entity.User{
		Name:     newuser.Name,
		Email:    newuser.Email,
		Password: newuser.Password,
		Umur:     newuser.Umur,
		Gender:   newuser.Gender,
	}
	newuserlah, err := s.repository.Create(user)
	return newuserlah, err
}
