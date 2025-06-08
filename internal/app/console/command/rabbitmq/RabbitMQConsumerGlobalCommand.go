package rabbitmq

import (
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/v2/pkg"
	"github.com/yusologia/go-core/v2/rabbitmq"
	"service/internal/pkg/config"
)

type RabbitMQConsumerGlobalCommand struct{}

func (class *RabbitMQConsumerGlobalCommand) Command(cobraCmd *cobra.Command) {
	addCommand := cobra.Command{
		Use:  "rabbitmq:consumer-global",
		Long: "RabbitMQ Consumer Global",
		Run: func(cmd *cobra.Command, args []string) {
			logiapkg.InitDevMode()

			DBClose := config.InitDB()
			defer DBClose()

			rabbitmqConn := config.InitRabbitMQ()
			defer rabbitmqConn()

			class.Handle()
		},
	}

	cobraCmd.AddCommand(&addCommand)
}

func (class *RabbitMQConsumerGlobalCommand) Handle() {
	logiarabbitmq.Consume(logiarabbitmq.RABBITMQ_CONNECTION_GLOBAL, []logiarabbitmq.RabbitMQConsumeOpt{
		//{
		//	Exchange: "service.domain.feature.action.exchange",
		//	Consumer: &rabbitmq.TestingConsumer{},
		//},
		//{
		//	Queue:    "service.domain.feature.action.queue",
		//	Consumer: &rabbitmq.TestingConsumer{},
		//},
	})
}
