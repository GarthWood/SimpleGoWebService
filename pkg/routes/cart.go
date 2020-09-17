// The routing package contains all HTTP routes served by the service.
package routes

import (
	"CartService/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	cartIdPathParam = "cartId"
)

// The cart routes expose an API for cart management.
type CartRoutes struct {
	Controller controllers.CartController `inject:"inline"`
}

// Creates a new cart router.
func (recv *CartRoutes) Create(router *mux.Router) {
	router.HandleFunc(path("/v1/cart/{0}", cartIdPathParam), recv.getCart).Methods("GET")
}

// The handler for the GET cart route.
func (recv *CartRoutes) getCart(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	response := recv.Controller.GetCart(vars[cartIdPathParam])
	writeResponse(response, w)
}
