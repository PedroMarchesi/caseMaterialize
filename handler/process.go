package handler

import (
	"encoding/json"
	"imports/driver"
	"imports/repository"
	"net/http"
)

func Process(responseWriter http.ResponseWriter, r *http.Request) {
	var driver driver.StructApp
	var p repository.ListProcess

	//Convertendo o JSON passado na requisição para a model
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		driver.ResponseWithErrorJSON(responseWriter, http.StatusInternalServerError, nil, "Objeto Inválido")
		return
	}
	defer r.Body.Close()

	result, message, err := p.WorkingWithNode()

	if err != nil {
		driver.ResponseWithErrorJSON(responseWriter, http.StatusInternalServerError, nil, message)
		return
	}
	driver.ResponseWithJSON(responseWriter, http.StatusOK, result, message)
}
