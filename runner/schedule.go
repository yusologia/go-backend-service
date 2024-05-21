package runner

import (
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/config"
	"github.com/yusologia/go-core/console"
	"service/app/Console"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "schedule",
		Long: "Running Schedule",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitEnv()
			//baseConf.InitDB()

			console.Schedules(Console.RegisterSchedule)
		},
	})
}
