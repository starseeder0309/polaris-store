package services

import (
	"errors"
)

type Service interface {
	Member
	Order
	Product
}

var ErrInvalidPassword = errors.New("INVALID PASSWORD")
