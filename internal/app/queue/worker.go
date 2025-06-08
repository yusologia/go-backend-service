package queue

import (
	"github.com/yusologia/go-core/v2/queue"
	Telegram "service/internal/app/queue/job"
	"service/internal/pkg/constant"
)

func Register() []logiaqueue.JobConf {
	return []logiaqueue.JobConf{
		{
			Context:     Telegram.TelegramMessageJob{},
			JobFunc:     (*Telegram.TelegramMessageJob).Consume,
			Concurrency: 1,
			QueueName:   constant.QUEUE_HIGH,
			JobName:     constant.JOB_TELEGRAM_MESSAGE,
			Priority:    10,
		},
	}
}
