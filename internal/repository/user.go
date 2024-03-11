package repository

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindAll() ([]entity.User, error)
	FindByID(ID int) (entity.User, error)
	Create(user entity.User) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user entity.User) (entity.User, error) {
	err := u.db.Create(&user).Error

	return user, err
}

func (u *UserRepository) FindAll() ([]entity.User, error) {
	var user []entity.User
	err := u.db.Find(&user).Error

	return user, err
}

func (u *UserRepository) FindByID(ID int) (entity.User, error) {
	var user entity.User
	err := u.db.Where("id = ?", ID).First(&user).Error

	return user, err
}
