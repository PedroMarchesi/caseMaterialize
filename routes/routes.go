package routes

import "github.com/gorilla/mux"

func InitializeRoutes(subrouter *mux.Router) {
	Config(subrouter)
	Process(subrouter)
}
