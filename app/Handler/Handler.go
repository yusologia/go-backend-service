package Handler

import (
	logiares "github.com/yusologia/go-response"
	"net/http"
)

type Handler struct{}

func (hand Handler) Index(w http.ResponseWriter, r *http.Request) {
	res := logiares.Response{}
	res.Success(w)
}
