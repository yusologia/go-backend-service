package runner

import (
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/v2/console"
	"github.com/yusologia/go-core/v2/pkg"
	"service/internal/app/console"
	"service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "logia:schedule",
		Long: "Running Schedule",
		Run: func(cmd *cobra.Command, args []string) {
			logiapkg.InitDevMode()

			config.InitTZ()

			DBClose := config.InitDB()
			defer DBClose()

			logiaconsole.Schedules(console.RegisterSchedule)
		},
	})
}
