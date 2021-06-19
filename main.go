package main

import (
	"fmt"
	"imports/models"
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

	valueCors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	})

	structApplication.Router = mux.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.LoggingHandler(os.Stdout, valueCors.Handler(structApplication.Router))))
}
