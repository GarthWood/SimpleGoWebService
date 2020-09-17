package controllers

import (
	"CartService/pkg/model"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCart(t *testing.T) {

	Convey("When getting a valid cart ID", t, func() {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cartReaderMock := NewMockCartReader(ctrl)

		// object under test
		cartController := &CartController{
			Reader: cartReaderMock,
		}

		const CartId = "cart1234"

		cartModel := &model.Cart{
			Id: CartId,
		}

		cartReaderMock.
			EXPECT().
			FetchCart(gomock.Eq(CartId)).
			Return(cartModel, nil).
			Times(1)

		// execute the action
		response := cartController.GetCart(CartId)

		Convey("It should return a valid response", func() {
			So(response.GetStatus(), ShouldEqual, http.StatusOK)
			So(response.GetBody().(*model.Cart).Id, ShouldEqual, CartId)
		})
	})

	Convey("When getting an invalid cart ID", t, func() {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cartReaderMock := NewMockCartReader(ctrl)

		cartController := &CartController{
			Reader: cartReaderMock,
		}

		const CartId = "cart12347"

		err := errors.New("some error code")

		cartReaderMock.
			EXPECT().
			FetchCart(gomock.Eq(CartId)).
			Return(nil, err).
			Times(1)

		response := cartController.GetCart(CartId)

		Convey("It should return an error response", func() {
			So(response.GetStatus(), ShouldEqual, http.StatusNotFound)
			So(response.GetBody().(*model.ErrorBody).Reason, ShouldEqual, "Cannot find cart")
			So(response.GetBody().(*model.ErrorBody).ErrorCode, ShouldEqual, "some error code")
		})
	})
}
