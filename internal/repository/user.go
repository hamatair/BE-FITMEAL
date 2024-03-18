package repository

import (
	"fmt"
	"intern-bcc/entity"
	"intern-bcc/model"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindAll() ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
	UserEditProfile(user model.EditProfile, id string) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
	UserChangePassword(param model.ChangePassword, id string) (entity.User, error)
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

func hitungNutrisi(aktivitas string, gender string, umur uint, BB uint, TB uint) (float32, float32, float32, float32) {
	var kalori float32
	if gender == "male" {
		kalori = 66 + (13.7*float32(BB) + (5 * float32(TB)) - (6.8 * float32(umur)))
	} else if gender == "female" {
		kalori = 655 + (9.6*float32(BB) + (1.8 * float32(TB)) - (4.7 * float32(umur)))
	}

	if aktivitas == "sangat jarang olahraga" {
		kalori *= 1.2
	} else if aktivitas == "jarang olahraga" {
		kalori *= 1.375
	} else if aktivitas == "sering olahraga" {
		kalori *= 1.55
	} else if aktivitas == "sangat sering olahraga" {
		kalori *= 1.725
	}

	protein := (0.15 * kalori) / 4
	karbohidrat := (0.6 * kalori) / 4
	lemak := (0.15 * kalori) / 9

	return kalori, protein, karbohidrat, lemak
}

func (u *UserRepository) UserEditProfile(user model.EditProfile, id string) (entity.User, error) {
	var data entity.User
	err := u.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}

	data.UserName = user.UserName
	data.Umur = user.Umur
	data.Alamat = user.Alamat
	data.BeratBadan = user.BeratBadan
	data.TinggiBadan = user.TinggiBadan

	kalori, protein, karbohidrat, lemak := hitungNutrisi(data.Aktivitas, data.Gender, data.Umur, data.BeratBadan, data.TinggiBadan)

	data.Kalori = kalori
	data.Protein = protein
	data.Karbohidrat = karbohidrat
	data.Lemak = lemak

	err = u.db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return data, err
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

func (u *UserRepository) UserChangePassword(param model.ChangePassword, id string) (entity.User, error) {
	var data entity.User
	err := u.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		fmt.Println("di pencarian ada", err)
	}

	data.Password = param.NewPassword

	err = u.db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		fmt.Println("pada save ada", err)
	}

	return data, err

}
