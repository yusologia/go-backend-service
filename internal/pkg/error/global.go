package error

import (
	"github.com/yusologia/go-core/v2/response"
	"net/http"
)

func ErrXtremePrivateAPIAuthentication(internalMsg string) {
	logiares.Error(http.StatusUnauthorized, "Doesn't have access to private api", internalMsg, false, nil)
}
