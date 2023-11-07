package response
type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"after exicution"`
	Error      interface{} `json:"error,omitempty"`
}

func Responses(statusCode int, message string, data interface{}, err interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}
}
