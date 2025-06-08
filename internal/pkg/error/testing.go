package error

import (
	"github.com/yusologia/go-core/v2/response"
	"net/http"
)

func ErrXtremeTestingGet(internalMsg string) {
	logiares.Error(http.StatusNotFound, "Testing not found", internalMsg, false, nil)
}

func ErrXtremeTestingSave(internalMsg string) {
	logiares.Error(http.StatusInternalServerError, "Unable to save testing", internalMsg, false, nil)
}

func ErrXtremeTestingDelete(internalMsg string) {
	logiares.Error(http.StatusInternalServerError, "Unable to delete testing", internalMsg, false, nil)
}

func ErrXtremeTestingSubSave(internalMsg string) {
	logiares.Error(http.StatusInternalServerError, "Unable to save testing sub", internalMsg, false, nil)
}

func ErrXtremeTestingSubDelete(internalMsg string) {
	logiares.Error(http.StatusInternalServerError, "Unable to delete testing sub", internalMsg, false, nil)
}
