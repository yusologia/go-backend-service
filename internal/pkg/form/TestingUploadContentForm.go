package form

import (
	"github.com/yusologia/go-core/v2/middleware"
	"net/http"
	"service/internal/pkg/core"
)

type TestingUploadContentForm struct {
	Content string `json:"content" validate:"required"`
}

func (rule *TestingUploadContentForm) Validate() {
	va := logiamdw.Validator{}
	va.Make(rule)
}

func (rule *TestingUploadContentForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
