// The routing package contains all HTTP routes served by the service.
package routes

import (
	"CartService/pkg/app"
	"CartService/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// The error response data object return for all HTTP errors.
type httpError struct {
	Reason    string `json:"reason"`
	ErrorCode string `json:"errorCode"`
}

// Handles writing either an error or a success response using JSON as the
// default response body type.
func writeResponse(response model.Response, writer http.ResponseWriter) {

	err := response.GetError()

	if err != nil {
		handleError(err, writer)
	} else {
		handleSuccess(response.GetBody(), writer)
	}
}

// Translates application errors into HTPP errors and returns a parsable
// error response for clients to handle.
func handleError(err error, writer http.ResponseWriter) {

	statusCode := http.StatusInternalServerError
	appError := err.(*app.Error)

	switch appError.Code {
	case app.NotFoundError:
		statusCode = http.StatusNotFound
	case app.AlreadyExistsError:
		statusCode = http.StatusBadRequest
	}

	herror := &httpError{
		Reason:    appError.Reason,
		ErrorCode: appError.Code,
	}

	if result, err := json.Marshal(herror); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(statusCode)
		writer.Write(result)
	}
}

// Completes the HTTP request with a response body and successful
// HTTP status.
func handleSuccess(body interface{}, writer http.ResponseWriter) {

	if body != nil {
		if result, err := json.Marshal(body); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			writer.WriteHeader(http.StatusOK)
			writer.Write(result)
		}
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

// A helper function to construct an HTTP path from a base path and replacement values.
func path(p string, values ...string) string {

	for i, v := range values {
		index := fmt.Sprintf("{%d}", i)
		value := fmt.Sprintf("{%s}", v)
		p = strings.Replace(p, index, value, 1)
	}

	return p
}
