package repository

import (
	"intern-bcc/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TopUpRepositoryI interface {
	FindById(id string) (entity.TopUp, error)
	Insert(t *entity.TopUp) error
	Update(t *entity.TopUp) error
}

type TopUpRepository struct {
	db *gorm.DB
}

func NewTopUp(db *gorm.DB) TopUpRepositoryI {
	return &TopUpRepository{
		db: db,
	}
}

// FindById implements TopUpRepository.
func (r *TopUpRepository) FindById(id string) (entity.TopUp, error) {
	var topUpData entity.TopUp
	Id, _ := uuid.Parse(id)
	err := r.db.Where("id = ?", Id).First(&topUpData).Error
	return	topUpData, err
}

// Insert implements TopUpRepository.
func (r *TopUpRepository) Insert(t *entity.TopUp) error {

	err := r.db.Create(&t).Error

	return err
}

// Update implements TopUpRepository.
func (r *TopUpRepository) Update(t *entity.TopUp) error {

	var topUpData entity.TopUp
	err := r.db.Where("id = ?", t.ID).First(&topUpData).Error
	if err != nil {
		return err
	}

	topUpData.UserID = t.ID
	topUpData.Status = t.Status
	topUpData.Amount = t.Amount
	topUpData.SnapUrl = t.SnapUrl

	err = r.db.Where("id = ?", t.ID).Updates(topUpData).Error

	return err
}


