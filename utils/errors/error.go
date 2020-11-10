package errors

import "net/http"

type RestErr struct {
	Message string `josn:"message"`
	Status    int    `josn:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	 return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	 return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not_found.",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "not_found.",
	}
}
