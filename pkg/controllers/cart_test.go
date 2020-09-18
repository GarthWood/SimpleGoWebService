package controllers

import (
	"CartService/pkg/app"
	"CartService/pkg/model"
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
			So(response.GetError(), ShouldBeNil)
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

		err := app.NewError("some reason", "some error code")

		cartReaderMock.
			EXPECT().
			FetchCart(gomock.Eq(CartId)).
			Return(nil, err).
			Times(1)

		response := cartController.GetCart(CartId)

		Convey("It should return an error response", func() {
			So(response.GetError().(*app.Error).Code, ShouldEqual, "some error code")
			So(response.GetError().(*app.Error).Reason, ShouldEqual, "some reason")
			So(response.GetBody(), ShouldBeNil)
		})
	})
}
