package core

import (
	"encoding/json"
	"github.com/yusologia/go-core/v2/response"
	"net/http"
)

type FormInterface interface {
	Validate()
}

type APIFormInterface interface {
	APIParse(r *http.Request)
}

type BaseForm struct{}

func (BaseForm) APIParse(r *http.Request, rule interface{}) interface{} {
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		logiares.ErrLogiaBadRequest(err.Error())
	}

	return rule
}
