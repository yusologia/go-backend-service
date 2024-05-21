package Command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/config"
	"os"
)

type TestCommand struct{}

func (class *TestCommand) Command(cobraCmd *cobra.Command) {
	cobraCmd.AddCommand(&cobra.Command{
		Use:  "dev-test",
		Long: "Development Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitEnv()

			class.Handle()
		},
	})
}

func (class *TestCommand) Handle() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Working Directory:", dir)
}
