package repository
import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

type Repositry interface {
	FindAll() ([]entity.User, error)
	FindByID(ID int) (entity.User, error)
	Create(user entity.User) (entity.User, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var user []entity.User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *repository) FindByID(ID int) (entity.User, error) {
	var user entity.User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *repository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
