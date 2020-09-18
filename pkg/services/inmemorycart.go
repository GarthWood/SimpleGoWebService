// The service package contains all the technical implementations about how to access databases and the like.
package services

import (
	"CartService/pkg/app"
	"CartService/pkg/model"
	"fmt"
)

// The cart service contains the necessary functionality to read or write cart data
// from an im-memory data store.
type InMemoryCartService struct {
	cart *model.Cart
}

// Fetches the cart from a data store.
func (recv *InMemoryCartService) FetchCart(cartId string) (*model.Cart, error) {

	if recv.cart != nil && recv.cart.Id == cartId {
		return recv.cart, nil
	}

	return nil, app.NewError(fmt.Sprintf("Could not find cart %s", cartId), app.NotFoundError)
}

// Writes the cart to a data store.
func (recv *InMemoryCartService) WriteCart(cart *model.Cart) error {

	if recv.cart != nil && recv.cart.Id == cart.Id {
		return app.NewError(fmt.Sprintf("Cart already exists with id %s", cart.Id), app.AlreadyExistsError)
	} else {
		recv.cart = cart
		return nil
	}
}
