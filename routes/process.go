package routes

import (
	"imports/handler"

	"github.com/gorilla/mux"
)

func Process(subrouter *mux.Router) {
	subrouter.HandleFunc("/process", handler.Config).Methods("POST")
}
