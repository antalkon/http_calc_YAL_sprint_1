package models

// RequestBody структура для парсинга входного JSON
type RequestBody struct {
	Expression string `json:"expression"`
}

// ResponseBody структура для успешного ответа
type ResponseBody struct {
	Result float64 `json:"result"`
}

// ErrorResponse структура для ответа с ошибкой
type ErrorResponse struct {
	Error string `json:"error"`
}
