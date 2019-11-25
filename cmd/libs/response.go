package libs

import (
	"net/http"
)

// ResponseWrapper provides a general template for a JSON data object
type ResponseWrapper struct {
	Data interface{} `json:"data"`
}

// ErrorWrapper provides a general template for the respones
type ErrorWrapper struct {
	Errors []StandardError `json:"errors"`
}

// StandardError is the normal error to be returned
type StandardError struct {
	Code        int    `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// NewErrorWrapper returns an ErrorWrapper with the appropriate parameters
func NewErrorWrapper(code int, description string) ErrorWrapper {
	return ErrorWrapper{Errors: []StandardError{{Code: code, Title: http.StatusText(code), Description: description}}}
}

// NewStandardError returns a StandardError with the appropriate parameters
func NewStandardError(code int, description string) StandardError {
	return StandardError{Code: code, Title: http.StatusText(code), Description: description}
}
