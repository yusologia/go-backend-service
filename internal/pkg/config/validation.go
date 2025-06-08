package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/yusologia/go-core/v2/middleware"
)

func InitValidation() {
	v := logiamdw.Validator{}
	v.RegisterValidation(func(validate *validator.Validate) {
		// Enter your custom validation rules
	})
}
