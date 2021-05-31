package model

type ErrorDetail struct {
	ErrorType string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}



type Response struct {
	Data    string `json:"data"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
