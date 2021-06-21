package response

//ResponseError representa o response
type ResponseError struct {
	Error Error       `json:"error"`
	Node  interface{} `json:"node"`
}
