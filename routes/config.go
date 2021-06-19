package routes

import (
	"imports/handler"

	"github.com/gorilla/mux"
)

func Config(subrouter *mux.Router) {
	subrouter.HandleFunc("/config/{N}", handler.Config).Methods("POST")
}
