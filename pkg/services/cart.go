// The service package contains all the technical implementations about how to access databases and the like.
package services

import (
	"CartService/pkg/model"
	"errors"
)

// The cart service contains the necessary functionality to read or write cart data
// as well as aply promotional rules and coupons.
type CartService struct {
}

// Fetches the cart from a data store.
func (recv *CartService) FetchCart(cartId string) (*model.CartModel, error) {

	if cartId == "cart1234" {
		cart := &model.CartModel{
			Id:       "abcd",
			Products: []string{"p1", "p2", "p2"},
		}
		return cart, nil
	}

	return nil, errors.New("not_found")
}
