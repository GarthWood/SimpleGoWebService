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
		handleResult(response.GetResult(), writer)
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

	// translate the application error into an dedicated HTTP error object

	herror := &httpError{
		Reason:    appError.Reason,
		ErrorCode: appError.Code,
	}

	writeJsonResponse(herror, statusCode, writer)
}

// Completes the HTTP request with a response result and successful
// HTTP status.
func handleResult(result interface{}, writer http.ResponseWriter) {
	writeJsonResponse(result, http.StatusOK, writer)
}

// A helper function to write a JSON HTTP response.
func writeJsonResponse(object interface{}, statusCode int, writer http.ResponseWriter) {

	if object != nil {
		if result, err := json.Marshal(object); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			writer.WriteHeader(statusCode)
			writer.Write(result)
		}
	} else {
		writer.WriteHeader(statusCode)
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
