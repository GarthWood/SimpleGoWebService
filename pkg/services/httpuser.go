// The service package contains all the technical implementations about how to access databases and the like.
package services

// The user service provides functionality for all user related activities.
type HttpUserService struct {
}

// Checks if the access token provided is valid and has the necessary scopes and roles.
func (recv *HttpUserService) Validate(accessToken string) bool {

	// execute HTTP request here
	// todo

	return accessToken == "1234"
}
