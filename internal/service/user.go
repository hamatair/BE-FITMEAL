package service

import (
	"fmt"
	"intern-bcc/entity"
	"intern-bcc/internal/repository"
	"intern-bcc/model"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/jwt"

	"github.com/google/uuid"
)

type UserServiceInterface interface {
	FindAll() ([]entity.User, error)
	FindByID(ID int) (entity.User, error)
	Create(user model.Register) (entity.User, error)
	UserPersonalization(user model.Personalization, name string) (entity.User, error)
	Login(param model.Login) (model.LoginResponse, error)
	GetUser(param model.UserParam) (entity.User, error)
}

type UserService struct {
	userRepository repository.UserRepositoryInterface
	bcrypt bcrypt.Interface
	jwtAuth jwt.Interface
}

func NewUserService(repository repository.UserRepositoryInterface, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) UserServiceInterface {
	return &UserService{
		userRepository: repository,
		bcrypt: bcrypt,
		jwtAuth: jwtAuth,
	}
}

func (u *UserService) Create(param model.Register) (entity.User, error) {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil{
		return  entity.User{}, err
	}

	param.ID = uuid.New()
	param.Password = hashPassword
	nuser := entity.User{
		ID:       param.ID,
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	}

	newUser, err := u.userRepository.Create(nuser)

	return newUser, err

}

func (u*UserService) FindAll() ([]entity.User, error) {
	user, err := u.userRepository.FindAll()

	return user, err
}

func (u *UserService) FindByID(ID int) (entity.User, error) {
	user, err := u.userRepository.FindByID(ID)

	return user, err
}

func (u *UserService) UserPersonalization(user model.Personalization, name string) (entity.User, error){
	UserPersonalization , err := u.userRepository.UserPersonalization(user, name)
	if err != nil{
		fmt.Println("service", err)
	}
	
	return UserPersonalization, err
}

func (u *UserService) Login(param model.Login) (model.LoginResponse, error) {
	result := model.LoginResponse{}

	user, err := u.userRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	err = u.bcrypt.CompareAndHashPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := u.jwtAuth.CreateJWTToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return u.userRepository.GetUser(param)
}