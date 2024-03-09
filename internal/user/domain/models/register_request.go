package models

type RegisterRequest struct {
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
}
