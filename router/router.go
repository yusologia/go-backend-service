package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"service/router/api"
)

func Register(router *mux.Router) {
	version := os.Getenv("VERSION")
	service := os.Getenv("SERVICE")

	router = router.PathPrefix("/api").Subrouter()

	web := api.Web{}
	web.Routers(router.PathPrefix(fmt.Sprintf("/web/%s/%s", version, service)).Subrouter())

	mobile := api.Mobile{}
	mobile.Routers(router.PathPrefix(fmt.Sprintf("/mobile/%s/%s", version, service)).Subrouter())
}
