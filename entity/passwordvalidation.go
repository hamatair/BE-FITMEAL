package entity

import (
	"time"
)

type PasswordValidation struct {
	Email       string
	Kode        int
	ExpiredTime time.Time
}
