// Code generated by MockGen. DO NOT EDIT.
// Source: cart.go

// Package mock_controllers is a generated GoMock package.
package controllers

import (
	model "CartService/pkg/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCartReader is a mock of CartReader interface
type MockCartReader struct {
	ctrl     *gomock.Controller
	recorder *MockCartReaderMockRecorder
}

// MockCartReaderMockRecorder is the mock recorder for MockCartReader
type MockCartReaderMockRecorder struct {
	mock *MockCartReader
}

// NewMockCartReader creates a new mock instance
func NewMockCartReader(ctrl *gomock.Controller) *MockCartReader {
	mock := &MockCartReader{ctrl: ctrl}
	mock.recorder = &MockCartReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCartReader) EXPECT() *MockCartReaderMockRecorder {
	return m.recorder
}

// FetchCart mocks base method
func (m *MockCartReader) FetchCart(cartId string) (*model.CartModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCart", cartId)
	ret0, _ := ret[0].(*model.CartModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCart indicates an expected call of FetchCart
func (mr *MockCartReaderMockRecorder) FetchCart(cartId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCart", reflect.TypeOf((*MockCartReader)(nil).FetchCart), cartId)
}
