package models

type RequestBody struct {
	Expression string `json:"expression"`
}

type ResponseBody struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
