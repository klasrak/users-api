package rerrors

import (
	"errors"
	"fmt"
	"net/http"
)

// package rerrors shares errors accross
// all application layers.
// the name 'rerrors' is to not conflic with golang native errors pkg

// Type holds a type string and integer code for the error
type Type string

// Set of valid errorTypes
const (
	BadRequest Type = "BADREQUEST" // Validation errors
	Internal   Type = "INTERNAL"   // Server (500) and fallback errors
	NotFound   Type = "NOTFOUND"   // For not finding resource
	Forbidden  Type = "FORBIDDEN"  // The client has no access rights to the content so the server is refusing to respond
	Conflict   Type = "CONFLICT"   // Already exists - 409
)

// Error holds a custom error for the application
// which is helpful in returning a consistent
// error type/message from API endpoints
type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

// Error satisfies standard error interface
// we can return errors from this package as
// a regular old go _error_
func (e *Error) Error() string {
	return e.Message
}

// Status is a mapping errors to status codes
// Of course, this is somewhat redundant since
// our errors already map http status codes
func (e *Error) Status() int {
	switch e.Type {
	case BadRequest:
		return http.StatusBadRequest
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case Forbidden:
		return http.StatusForbidden
	case Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// Status checks the runtime type
// of the error and returns an http
// status code if the error is model.Error
func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

/*
* Error "Factories"
 */

// NewBadRequest to create 400 errors
func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

// NewInternal for 500 errors
func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: "Internal server error.",
	}
}

// NewNotFound to create an error for 404
func NewNotFound(name string, value string) *Error {
	var message string

	if name != "" && value == "" {
		message = fmt.Sprintf("resource: %v not found", name)
	} else {
		message = fmt.Sprintf("resource: %v with value: %v not found", name, value)
	}

	return &Error{
		Type:    NotFound,
		Message: message,
	}
}

// NewForbidden to create 403 errors
func NewForbidden(reason string) *Error {
	return &Error{
		Type:    Forbidden,
		Message: fmt.Sprintf("Forbidden. Reason: %v", reason),
	}
}

// NewConflict to create 409 erros
func NewConflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v not created: %v", name, value),
	}
}
