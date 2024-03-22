package model

import "github.com/google/uuid"

type TopUp struct {
	Amount uint `json:"amount"`
}

type TopUpReq struct {
	Amount uint      `json:"amount"`
	UserId uuid.UUID `json:"-"`
}

type TopUpRes struct {
	SnapUrl string `json:"snapUrl"`
}
