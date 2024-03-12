package tools

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func CustomPassword(fl validator.FieldLevel) bool {
	var (
		hasUpper   bool
		hasLower   bool
		hasNumeric bool
		hasSpecial bool
	)

	password := fl.Field().String()

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumeric = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumeric || !hasSpecial {
		return true
	}

	return false
}
