package response

//ResponseSuccess representa o response de sucesso
type ResponseSuccess struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Error   error       `json:"error"`
}
