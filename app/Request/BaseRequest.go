package Request

import (
	"encoding/json"
	logiaerr "github.com/yusologia/go-response/error"
	"net/http"
)

type RequestInterface interface {
	Validate(r *http.Request)
	Parse(r *http.Request)
}

type BaseRequest struct{}

func (BaseRequest) Parse(r *http.Request, rule interface{}) interface{} {
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		logiaerr.ErrLogiaBadRequest(err.Error())
	}

	return rule
}
