package entity

import (
	"github.com/google/uuid"
)

type TopUp struct {
	ID      uuid.UUID `json:"-"`
	UserID  uuid.UUID `json:"userId"`
	Status  uint      `json:"status"`
	Amount  uint      `json:"amount"`
	SnapUrl string    `json:"snapUrl"`
}



