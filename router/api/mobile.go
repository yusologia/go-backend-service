package api

import (
	"github.com/gorilla/mux"
	"service/app/Handler"
)

type Mobile struct{}

func (api Mobile) Routers(router *mux.Router) {
	var handler Handler.Handler
	router.HandleFunc("/dev-test", handler.Index).Methods("GET")
}
