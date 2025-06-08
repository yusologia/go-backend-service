package runner

import (
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/v2/pkg"
	"github.com/yusologia/go-core/v2/queue"
	"service/internal/app/queue"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "logia:queue",
		Long: "Running Queue",
		Run: func(cmd *cobra.Command, args []string) {
			logiapkg.InitDevMode()

			config.InitTZ()

			queueNames := cmd.Flags().String("q", constant.QUEUE_HIGH, "Queue name")
			configurations := queue.Register()

			worker := logiaqueue.Queue{Names: *queueNames}
			worker.Work(configurations)
		},
	})
}
