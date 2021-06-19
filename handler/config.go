package handler

import (
	"imports/driver"
	"imports/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func Config(responseWriter http.ResponseWriter, r *http.Request) {
	var driver driver.StructApp

	valueNode := mux.Vars(r)["N"]

	message, err := repository.CreatKnot(valueNode)

	if err != nil {
		driver.ResponseWithErrorJSON(responseWriter, http.StatusInternalServerError, nil, message)
	}
	driver.ResponseWithJSON(responseWriter, http.StatusOK, nil, "Nós configurados com sucesso")
}
