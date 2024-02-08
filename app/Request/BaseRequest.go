package Request

import (
	"encoding/json"
	error2 "github.com/yusologia/go-core/response/error"
	"net/http"
)

type RequestInterface interface {
	Validate(r *http.Request)
	Parse(r *http.Request)
}

type BaseRequest struct{}

func (BaseRequest) Parse(r *http.Request, rule interface{}) interface{} {
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		error2.ErrLogiaBadRequest(err.Error())
	}

	return rule
}
