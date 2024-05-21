package runner

import (
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/config"
	"github.com/yusologia/go-core/queue"
	"service/app/Service/Constant/Queue"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "queue",
		Long: "Running Queue",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitEnv()

			queueNames := cmd.Flags().String("q", Queue.QUEUE_HIGH, "Queue name")

			configurations := [...]queue.JobConf{
				//{
				//	Context:     Resizing.Resizing{},
				//	JobFunc:     (*Resizing.Resizing).Image,
				//	Concurrency: 10,
				//	QueueName:   Queue.QUEUE_HIGH,
				//	JobName:     Queue.JOB_REZISE_IMAGE,
				//	Priority:    1,
				//},
			}

			worker := queue.Queue{Names: *queueNames}
			worker.Work(configurations[:])
		},
	})
}
