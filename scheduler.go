package main

import (
	"github.com/joho/godotenv"
	"github.com/yusologia/go-core/console"
	"service/app/Console"
	"service/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	config.InitDB()

	console.Schedules(Console.Schedules)
}
