package validator

import (
	"go-react-todo/server/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	ValidateUser(user *models.User) error
}

type UserValidator struct{}

func NewUserValidator() IUserValidator {
	return &UserValidator{}
}

func (v *UserValidator) ValidateUser(user *models.User) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.RuneLength(8, 100).Error("Password must be between 8 and 100 characters")),
	)
}
