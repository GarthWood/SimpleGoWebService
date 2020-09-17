// The routing package contains all HTTP routes served by the service.
package routes

import (
	"CartService/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Handles writing either an error or a success response using JSON as the
// default response body type.
func writeResponse(response model.Response, writer http.ResponseWriter) {

	body := response.GetBody()

	if body != nil {
		if result, err := json.Marshal(body); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			writer.WriteHeader(response.GetStatus())
			writer.Write(result)
		}
	} else {
		writer.WriteHeader(response.GetStatus())
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
