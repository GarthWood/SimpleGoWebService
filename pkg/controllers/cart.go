// The controller package is the primary business logic for cart management.
package controllers

import (
	"CartService/pkg/model"
)

// The CartReader provide an abstraction to fetch a cart.
type CartReader interface {
	FetchCart(cartId string) (*model.Cart, error)
}

// The CartWriter provide an abstraction to write to a cart.
type CartWriter interface {
	WriteCart(cart *model.Cart) error
}

// The cart controller handles all business logic related to fetching a cart, writing to it and applying
// promotional rules and coupons.
type CartController struct {
	Reader CartReader `inject:""`
	Writer CartWriter `inject:""`
}

// Gets a cart using the specified cartId.
// It returns either an error or success response.
func (recv *CartController) GetCart(cartId string) model.Response {

	if cart, err := recv.Reader.FetchCart(cartId); err != nil {
		return model.NewErrorResponse(err)
	} else {
		return model.NewSuccessResponse(cart)
	}
}

func (recv *CartController) CreateCart() model.Response {

	// mock data
	cart := &model.Cart{
		Id:       "1234",
		Products: []model.Product{model.Product{Id: "p1"}, model.Product{Id: "p4"}, model.Product{Id: "p3"}},
	}

	if err := recv.Writer.WriteCart(cart); err != nil {
		return model.NewErrorResponse(err)
	} else {
		return model.NewEmptySuccessResponse()
	}
}
