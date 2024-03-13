package repository

import (
	"fmt"
	"intern-bcc/entity"
	"intern-bcc/model"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindAll() ([]entity.User, error)
	FindByID(ID int) (entity.User, error)
	Create(user entity.User) (entity.User, error)
	UserEditProfile(user model.EditProfile, name string) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
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

func (u *UserRepository) UserEditProfile(user model.EditProfile, name string) (entity.User, error) {
	var data entity.User
	err := u.db.Where("name = ?", name).First(&data).Error
	if err != nil {
		fmt.Println("di First ada", err)
	}

	data.Aktivitas = user.Aktivitas
	data.Alamat = user.Alamat
	data.Gender = user.Gender
	data.BeratBadan = user.BeratBadan
	data.TinggiBadan = user.TinggiBadan
	data.Umur = user.Umur

	err = u.db.Where("name = ?", name).Updates(&data).Error
	if err != nil {
		fmt.Println("pada save ada", err)
	}

	return data, err
}

func (u *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}