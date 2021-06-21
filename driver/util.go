package driver

import (
	"encoding/json"
	"imports/models"
	"imports/models/response"
	"net/http"
)

type StructApp models.App

//ResponseWithErrorJSON corresponde a resposta com json de Erro
func (appStruct *StructApp) ResponseWithErrorJSON(responseWriter http.ResponseWriter, code int, payload interface{}, message string) {
	var responseError response.ResponseError

	responseError.Error.Message = message
	responseError.Node = payload

	response, _ := json.Marshal(responseError)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(response)
}

//ResponseWithJSON corresponde a resposta com reposta com json
func (appStruct *StructApp) ResponseWithJSON(responseWriter http.ResponseWriter, code int, result float64, message string) {
	var responseSuccess response.ResponseSuccess

	responseSuccess.Message = message
	responseSuccess.Error = nil
	responseSuccess.Result = result

	response, _ := json.Marshal(responseSuccess)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(response)
}
