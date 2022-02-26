package auth_dto

import (
	"github.com/pkg/errors"
	"relia_system/shared"
)

var (
	ErrInvalidEmail = shared.NewCustomError(
		errors.New("email invalid"),
		"email invalid",
		"ErrInvalidEmail")

	ErrInvalidEmailOrPassword = shared.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrInvalidEmailOrPassword",
	)
)

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
