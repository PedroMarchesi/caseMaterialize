package main

import (
	"fmt"
	"imports/models"
	"imports/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//StructApplication representa o aplicativo
type StructApplication models.Application

func main() {
	structApplication := StructApplication{}

	os.Setenv("PORT", "5000")
	port := os.Getenv("PORT")

	fmt.Printf("Starting Server on port %s", port)
	structApplication.Run(port)
}

//Run inicia o servidor
func (structApplication *StructApplication) Run(port string) {
	structApplication.Initialize()

	valueCors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.LoggingHandler(os.Stdout, valueCors.Handler(structApplication.Router))))
}

//Initialize representa a inicialização do app.
func (structApplication *StructApplication) Initialize() {
	structApplication.Router = mux.NewRouter()
	subrouter := structApplication.Router.NewRoute().Subrouter()

	routes.InitializeRoutes(subrouter)
}
