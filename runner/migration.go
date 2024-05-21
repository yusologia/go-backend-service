package runner

import (
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/config"
	"github.com/yusologia/go-core/database/migration"
	baseConf "service/config"
	migrationDB "service/database/migration"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "migration",
		Long: "Running Migration",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitEnv()
			baseConf.InitDB()

			migration.Migrate(migrationDB.Tables(), migrationDB.Columns())
		},
	})
}
