package runner

import (
	"service/app/Console"
)

func init() {
	Console.RegisterCommand(rootCmd)
}
