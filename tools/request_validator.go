package tools

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var validate = validator.New()

func Validate(data interface{}) []string {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	msgErrs := []string{}

	if errs := validate.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			msg := err.Translate(trans)
			msgErrs = append(msgErrs, msg)
		}
	}

	return msgErrs
}
