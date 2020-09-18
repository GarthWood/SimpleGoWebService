// The model package contains all the data containers.
package model

// An abstaction representing a generic response
type Response interface {
	GetError() error
	GetBody() interface{}
}

// Create a new response containing a success HTTP status code and an response body
// to be returned to the client.
func NewSuccessResponse(body interface{}) Response {

	resp := &response{
		body: body,
	}

	return resp
}

// Create a new response containing an empty response body
// to be returned to the client.
func NewEmptySuccessResponse() Response {
	return &response{}
}

// Create a new response containing an error HTTP status code and an error response body
// to be returned to the client.
func NewErrorResponse(err error) Response {

	resp := &response{
		err: err,
	}

	return resp
}

// The base response data object.
type response struct {
	err  error
	body interface{}
}

func (recv *response) GetError() error {
	return recv.err
}

func (recv *response) GetBody() interface{} {
	return recv.body
}
