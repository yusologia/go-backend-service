package config

import (
	"fmt"
	"github.com/yusologia/go-core/helpers/logialog"
	"os"
)

func ErrorHandler() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "panic: %v\n", r)
			logialog.Error(r)
		}
	}()
}
