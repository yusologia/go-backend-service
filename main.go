package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/yusologia/go-core/config"
	"github.com/yusologia/go-core/router"
	"net/http"
	baseConf "service/config"
	router2 "service/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	config.SetHost()

	//baseConf.InitDB()
	baseConf.InitCors()
	baseConf.InitValidation()

	newCors := cors.New(baseConf.CorsOptions)

	newRouter := mux.NewRouter()
	router.RegisterRouter(newRouter, router2.Register)

	fmt.Println(fmt.Sprintf("Server started on %s", config.HostFull))

	err = http.ListenAndServe(config.Host+":"+config.Port, newCors.Handler(newRouter))
	if err != nil {
		panic(err)
	}
}
