package runner

import (
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/config"
	baseConf "service/config"
	"service/database/seeder"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "seeder",
		Long: "Running Seeder",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitEnv()
			baseConf.InitDB()

			seeder.Run()
		},
	})
}
