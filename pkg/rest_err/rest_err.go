package rest_err

import "net/http"

type RestErr struct {
	Code    int
	Err     string
	Message string
}

func (e RestErr) Error() string {
	return e.Message
}

func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Code:    http.StatusBadRequest,
		Err:     "bad_request",
		Message: msg,
	}
}

func NewInternalServeError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusInternalServerError,
		Message: message,
		Err:     "internal_server_error",
	}
}
