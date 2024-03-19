package service

import (
	"fmt"
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
	"intern-bcc/model"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/jwt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/google/uuid"
)

type UserServiceInterface interface {
	FindAll() ([]entity.User, error)
	Create(nuser model.Register) (entity.User, error)
	UserEditProfile(nuser model.EditProfile, id string) (entity.User, error)
	Login(param model.Login) (model.LoginResponse, error)
	GetUser(param model.UserParam) (entity.User, error)
	UserChangePassword(param model.ChangePassword, id string) (entity.User, error)
	CreateCodeVerification(param model.ForgotPassword) error
	CheckCode(param model.ForgotPassword) error
	ChangePasswordBeforeLogin(param model.ChangePasswordBeforeLogin) error
}

type UserService struct {
	userRepository repository.UserRepositoryInterface
	bcrypt         bcrypt.Interface
	jwtAuth        jwt.Interface
}

func NewUserService(repository repository.UserRepositoryInterface, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) UserServiceInterface {
	return &UserService{
		userRepository: repository,
		bcrypt:         bcrypt,
		jwtAuth:        jwtAuth,
	}
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

func (u *UserService) Create(param model.Register) (entity.User, error) {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return entity.User{}, err
	}

	kalori, protein, karbohidrat, lemak := hitungNutrisi(param.Aktivitas, param.Gender, param.Umur, param.BeratBadan, param.TinggiBadan)
	nuser := entity.User{
		ID:          uuid.New(),
		UserName:    param.UserName,
		Email:       param.Email,
		Password:    hashPassword,
		Aktivitas:   param.Aktivitas,
		Gender:      param.Gender,
		Umur:        param.Umur,
		Alamat:      param.Alamat,
		BeratBadan:  param.BeratBadan,
		TinggiBadan: param.TinggiBadan,
		Kalori:      kalori,
		Protein:     protein,
		Karbohidrat: karbohidrat,
		Lemak:       lemak,
	}

	newUser, err := u.userRepository.Create(nuser)

	return newUser, err

}

func (u *UserService) FindAll() ([]entity.User, error) {
	nuser, err := u.userRepository.FindAll()

	return nuser, err
}

func (u *UserService) UserEditProfile(nuser model.EditProfile, id string) (entity.User, error) {
	UserPersonalization, err := u.userRepository.UserEditProfile(nuser, id)
	if err != nil {
		return UserPersonalization, err
	}

	return UserPersonalization, err
}

func (u *UserService) Login(param model.Login) (model.LoginResponse, error) {
	result := model.LoginResponse{}

	nuser, err := u.userRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	err = u.bcrypt.CompareAndHashPassword(nuser.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := u.jwtAuth.CreateJWTToken(nuser.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return u.userRepository.GetUser(param)
}

func (u *UserService) UserChangePassword(param model.ChangePassword, id string) (entity.User, error) {
	uuid, _ := uuid.Parse(id)
	cekUser, err := u.userRepository.GetUser(model.UserParam{
		ID: uuid,
	})
	if err != nil {
		return cekUser, err
	}

	err = u.bcrypt.CompareAndHashPassword(cekUser.Password, param.OldPassword)
	if err != nil {
		return cekUser, err
	}

	fmt.Println(param.NewPassword)
	fmt.Println(param.ConfirmPassword)
	newpassword, _ := u.bcrypt.GenerateFromPassword(param.NewPassword)
	err = u.bcrypt.CompareAndHashPassword(newpassword, param.ConfirmPassword)
	if err != nil {
		return cekUser, err
	}

	param.NewPassword = newpassword
	nuser, err := u.userRepository.UserChangePassword(param, id)
	if err != nil {
		return nuser, err
	}

	return nuser, err
}

func (u *UserService) CreateCodeVerification(param model.ForgotPassword) error {
	randomNumber := rand.Intn(8999) + 1000

	param.Kode = randomNumber
	auth := smtp.PlainAuth(
		"",
		"gantengsekalibagas123@gmail.com",
		"mpta kssz otvb zhsd",
		"smtp.gmail.com",
	)

	msg := fmt.Sprintf("Subject: this is message %d", param.Kode)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"gantengsekalibagas123@gmail.com",
		[]string{"akabeaer212@gmail.com"},
		[]byte(msg),
	)

	if err != nil {
		return err
	}

	expiredTime := time.Now().Add(time.Hour)
	param.ExpiredTime = expiredTime

	err = u.userRepository.CreateCodeVerification(param)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) CheckCode(param model.ForgotPassword) error {
	dataCode, err := u.userRepository.GetDataCode(param)
	if err != nil {
		return err
	}

	param.ExpiredTime = time.Now()
	if param.ExpiredTime.After(dataCode.ExpiredTime) {
		return err
	}

	if param.Kode != dataCode.Kode {
		return err
	}

	return err
}

func (u *UserService) ChangePasswordBeforeLogin(param model.ChangePasswordBeforeLogin) error {
	dataUser, err := u.userRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return err
	}
	newPassword, err := u.bcrypt.GenerateFromPassword(param.NewPassword)
	if err != nil {
		return err
	}

	err = u.bcrypt.CompareAndHashPassword(newPassword, param.ConfirmPassword)
	if err != nil {
		return err
	}

	dataUser.Password = newPassword

	err = u.userRepository.ChangePasswordBeforeLogin(dataUser)
	if err != nil {
		return nil
	}

	return err
}
