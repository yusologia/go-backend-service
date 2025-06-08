package runner

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/v2"
	"github.com/yusologia/go-core/v2/pkg"
	"net/http"
	"service/internal/app/privateapi"
	"service/internal/pkg/config"
	"time"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "logia:private-api",
		Long: "Running Private API",
		Run: func(cmd *cobra.Command, args []string) {
			logiapkg.InitDevMode()

			config.InitPrivateAPI()
			config.InitTZ()
			config.InitCors()
			config.InitMail()
			config.InitValidation()
			config.InitCache(time.Hour, time.Hour)

			DBClose := config.InitDB()
			defer DBClose()

			rabbitMQClose := config.InitRabbitMQ()
			defer rabbitMQClose()

			newRoute := mux.NewRouter()
			logiacore.RegisterRouter(newRoute, privateapi.Register)

			fmt.Println(fmt.Sprintf("Server started on %s", config.PrivateHost))

			err := http.ListenAndServe(config.PrivateHost, newRoute)
			if err != nil {
				panic(err)
			}
		},
	})
}
