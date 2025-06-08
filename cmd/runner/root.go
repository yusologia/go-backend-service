package runner

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/v2"
	"github.com/yusologia/go-core/v2/pkg"
	"net/http"
	"os"
	"service/internal/app/api"
	"service/internal/pkg/config"
)

var rootCmd = &cobra.Command{
	Use:  "root",
	Long: "Running service api",
	Run: func(cmd *cobra.Command, args []string) {
		logiapkg.InitDevMode()
		logiapkg.InitHost()

		config.InitTZ()
		config.InitCors()
		config.InitMail()
		config.InitRPC()
		config.InitValidation()

		DBClose := config.InitDB()
		defer DBClose()

		// TODO: Aktifkan saat up ke operational
		//xtremedb.Migrate(config.PgSQL, database.Migrations())

		rabbitMQClose := config.InitRabbitMQ()
		defer rabbitMQClose()

		newCors := cors.New(config.CorsOptions)

		newRoute := mux.NewRouter()
		logiacore.RegisterRouter(newRoute, api.Register)

		fmt.Println(fmt.Sprintf("Server started on %s", logiapkg.HostFull))

		err := http.ListenAndServe(logiapkg.Host, newCors.Handler(newRoute))
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
	rootCmd.PersistentFlags().BoolVar(&logiapkg.DevMode, "dev", false, "Set for development mode")
}
