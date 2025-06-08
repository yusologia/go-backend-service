package runner

import (
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/v2/pkg"
	"service/internal/app/database"
	"service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "logia:seeder",
		Long: "Running Seeder",
		Run: func(cmd *cobra.Command, args []string) {
			logiapkg.InitDevMode()

			config.InitTZ()

			DBClose := config.InitDB()
			defer DBClose()

			database.Seeder()
		},
	})
}
