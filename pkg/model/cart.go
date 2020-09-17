// The model package contains all the data containers.
package model

// The cart model hold information related to a user's cart.
type CartModel struct {
	Id       string   `json:"id"`
	Products []string `json:"products"`
}
