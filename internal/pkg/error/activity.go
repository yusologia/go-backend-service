package error

import (
	"github.com/yusologia/go-core/v2/response"
	"net/http"
)

func ErrXtremeActivityNotFound() {
	logiares.Error(http.StatusNotFound, "Object not found", "", false, nil)
}

func ErrXtremeActivitySave(internalMsg string) {
	logiares.Error(http.StatusInternalServerError, "Unable to save activity", internalMsg, false, nil)
}

func ErrXtremeActivityActionType() {
	logiares.Error(http.StatusInternalServerError, "Object action doesn't exists!!", "", false, nil)
}
