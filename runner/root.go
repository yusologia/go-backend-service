package runner

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/config"
	"github.com/yusologia/go-core/router"
	"net/http"
	"os"
	baseConf "service/config"
	router2 "service/router"
)

var rootCmd = &cobra.Command{
	Use:  "root",
	Long: "Running service api",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitEnv()
		config.SetHost()

		//baseConf.InitDB()
		baseConf.InitCors()
		baseConf.InitValidation()

		newCors := cors.New(baseConf.CorsOptions)

		newRouter := mux.NewRouter()
		router.RegisterRouter(newRouter, router2.Register)

		fmt.Println(fmt.Sprintf("Server started on %s", config.HostFull))

		err := http.ListenAndServe(config.Host+":"+config.Port, newCors.Handler(newRouter))
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&config.EnvMode, "env", false, "Set for with .env file")
}
