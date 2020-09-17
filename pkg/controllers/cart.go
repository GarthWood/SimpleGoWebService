// The controller package is the primary business logic for cart management.
package controllers

import (
	"CartService/pkg/model"
	"net/http"
)

// The CartReaders provide an abstraction to fetch a cart.
type CartReader interface {
	FetchCart(cartId string) (*model.CartModel, error)
}

// The cart controller handles all business logic related to fetching a cart, writing to it and applying
// promotional rules and coupons.
type CartController struct {
	CartReader CartReader `inject:""`
}

// Gets a cart using the specified cartId.
// It returns either an error or success response.
func (recv *CartController) GetCart(cartId string) model.Response {

	if cart, err := recv.CartReader.FetchCart(cartId); err != nil {
		return model.NewErrorResponse(http.StatusNotFound, "Cannot find cart", err.Error())
	} else {
		return model.NewSuccessResponse(cart)
	}
}
