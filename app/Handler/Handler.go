package Handler

import (
	"github.com/yusologia/go-core/response"
	"net/http"
)

type Handler struct{}

func (hand Handler) Index(w http.ResponseWriter, r *http.Request) {
	res := response.Response{}
	res.Success(w)
}
