package routes

import (
	"imports/handler"

	"github.com/gorilla/mux"
)

func Process(subrouter *mux.Router) {

	//Rota responsável pela execução do processo
	subrouter.HandleFunc("/process", handler.Process).Methods("POST")
}
