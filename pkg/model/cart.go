// The model package contains all the data containers.
package model

// The cart model hold information related to a user's cart.
type Cart struct {
	Id       string    `json:"id"`
	Products []Product `json:"products"`
}
