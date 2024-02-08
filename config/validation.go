package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/yusologia/go-core/middleware"
)

func InitValidation() {
	v := middleware.Validator{}
	v.RegisterValidation(func(validate *validator.Validate) {
		// Enter your custom validation rules
	})
}
