// The model package contains all the data containers.
package model

import (
	"net/http"
)

// An abstaction representing a generic response
type Response interface {
	GetStatus() int
	GetBody() interface{}
}

// Create a new response containing a success HTTP status code and an response body
// to be returned to the client.
func NewSuccessResponse(body interface{}) Response {

	resp := &response{
		statusCode: http.StatusOK,
		body:       body,
	}

	return resp
}

// Create a new response containing an error HTTP status code and an error response body
// to be returned to the client.
func NewErrorResponse(statusCode int, reason string, errorCode string) Response {

	body := &ErrorBody{
		Reason:    reason,
		ErrorCode: errorCode,
	}

	resp := &response{
		statusCode: statusCode,
		body:       body,
	}

	return resp
}

// The base response data object.
type response struct {
	statusCode int
	body       interface{}
}

func (recv *response) GetStatus() int {
	return recv.statusCode
}

func (recv *response) GetBody() interface{} {
	return recv.body
}

// The error response data object return for all HTTP errors.
type ErrorBody struct {
	Reason    string `json:"reason"`
	ErrorCode string `json:"errorCode"`
}
