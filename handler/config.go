package handler

import (
	"imports/driver"
	"imports/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func Config(responseWriter http.ResponseWriter, r *http.Request) {
	var driver driver.StructApp

	valueNode := mux.Vars(r)["N"] //Valor informado na rota da requisição

	message, node, err := repository.CreatNode(valueNode)

	if err != nil {
		driver.ResponseWithErrorJSON(responseWriter, http.StatusInternalServerError, nil, message)
		return
	}
	driver.ResponseWithJSON(responseWriter, http.StatusOK, float64(node), "'Nós' configurados com sucesso")
}
