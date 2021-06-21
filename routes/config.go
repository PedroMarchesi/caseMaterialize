package routes

import (
	"imports/handler"

	"github.com/gorilla/mux"
)

func Config(subrouter *mux.Router) {

	//Rota responsável pela configuração dos "nós"
	subrouter.HandleFunc("/config/{N}", handler.Config).Methods("POST")
}
