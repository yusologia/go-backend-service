package config

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/yusologia/go-core/v2/database"
	"github.com/yusologia/go-core/v2/rabbitmq"
	"log"
	"os"
	"time"
)

func InitRabbitMQ() func() {
	var rabbitMQClose func()

	logiarabbitmq.RabbitMQSQL, rabbitMQClose = logiadb.Connect(logiadb.DBConf{
		Driver:    logiadb.MYSQL_DRIVER,
		Host:      os.Getenv("DB_RABBITMQ_HOST"),
		Port:      os.Getenv("DB_RABBITMQ_PORT"),
		Username:  os.Getenv("DB_RABBITMQ_USERNAME"),
		Password:  os.Getenv("DB_RABBITMQ_PASSWORD"),
		Database:  os.Getenv("DB_RABBITMQ_DATABASE"),
		ParseTime: true,
	})

	logiarabbitmq.RabbitMQConf.Connection = make(map[string]logiarabbitmq.RabbitMQConnectionConf, 2)

	globalHost := os.Getenv("RABBITMQ_GLOBAL_HOST")
	if globalHost != "" {
		logiarabbitmq.RabbitMQConf.Connection[logiarabbitmq.RABBITMQ_CONNECTION_GLOBAL] = logiarabbitmq.RabbitMQConnectionConf{
			Host:     globalHost,
			Port:     os.Getenv("RABBITMQ_GLOBAL_PORT"),
			Username: os.Getenv("RABBITMQ_GLOBAL_USER"),
			Password: os.Getenv("RABBITMQ_GLOBAL_PASSWORD"),
		}
	}

	localHost := os.Getenv("RABBITMQ_LOCAL_HOST")
	if localHost != "" {
		logiarabbitmq.RabbitMQConf.Connection[logiarabbitmq.RABBITMQ_CONNECTION_LOCAL] = logiarabbitmq.RabbitMQConnectionConf{
			Host:     localHost,
			Port:     os.Getenv("RABBITMQ_LOCAL_PORT"),
			Username: os.Getenv("RABBITMQ_LOCAL_USER"),
			Password: os.Getenv("RABBITMQ_LOCAL_PASSWORD"),
		}
	}

	logiarabbitmq.RabbitMQConf.Timeout = 5 * time.Second

	return rabbitMQClose
}

func InitRabbitMQConnection() func() {
	logiarabbitmq.RabbitMQConnectionDial = make(map[string]*amqp091.Connection, 2)
	for connectionName, connectionConf := range logiarabbitmq.RabbitMQConf.Connection {
		conn, err := amqp091.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
			connectionConf.Username, connectionConf.Password, connectionConf.Host, connectionConf.Port))
		if err != nil {
			log.Panicf("Failed to connect to RabbitMQ: %s", err)
		}

		logiarabbitmq.RabbitMQConnectionDial[connectionName] = conn
	}

	rabbitMQConnClose := func() {
		for _, connection := range logiarabbitmq.RabbitMQConnectionDial {
			connection.Close()
		}
	}

	return rabbitMQConnClose
}
