package errors

import "net/http"

// RestError wraps app errors
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// InternalServerError wraps internal server error
func InternalServerError(e string) *RestError {
	return &RestError{
		Message: e,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

// BadRequestError wraps bad request error
func BadRequestError(e string) *RestError {
	return &RestError{
		Message: e,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NotFoundError wraps not found error
func NotFoundError(e string) *RestError {
	return &RestError{
		Message: e,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
