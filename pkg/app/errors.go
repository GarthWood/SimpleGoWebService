// A general application package containing system-wide components such as errors and config.
package app

// The supported error codes
const (
	UnexpectedError    = "unexpected_error"
	NotFoundError      = "not_found_error"
	AlreadyExistsError = "already_exists_error"
)

// Creates an appplication error with a reason and an error code.
func NewError(reason string, code string) error {
	return &Error{Reason: reason, Code: code}
}

// A general error container.
type Error struct {
	Reason string
	Code   string
}

// Override the Go error function.
func (recv *Error) Error() string {
	return recv.Reason
}
