// The filters package contains all middleware.
package filters

import (
	"net/http"
)

// An abstraction to validate a user access token.
type UserValidator interface {
	Validate(accessToken string) bool
}

// The filter that validates the authentication using the Authorization header.
type SecurityFilter struct {
	Validator UserValidator `inject:""`
}

// The middleware function.
func (recv *SecurityFilter) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		token := req.Header.Get("Authorization")

		if recv.Validator.Validate(token) {
			next.ServeHTTP(w, req)
		} else {
			http.Error(w, "Forbidden", http.StatusUnauthorized)
		}
	})
}
