// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/djudju12/ms-products/db/sqlc (interfaces: Querier)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination db/mock/product_mock.go github.com/djudju12/ms-products/db/sqlc Querier
//
// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/djudju12/ms-products/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockQuerier) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockQuerierMockRecorder) CreateProduct(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockQuerier)(nil).CreateProduct), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockQuerier) GetProduct(arg0 context.Context, arg1 int32) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockQuerierMockRecorder) GetProduct(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockQuerier)(nil).GetProduct), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockQuerier) ListProducts(arg0 context.Context, arg1 db.ListProductsParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockQuerierMockRecorder) ListProducts(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockQuerier)(nil).ListProducts), arg0, arg1)
}

// UpdateProductStatus mocks base method.
func (m *MockQuerier) UpdateProductStatus(arg0 context.Context, arg1 db.UpdateProductStatusParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductStatus", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductStatus indicates an expected call of UpdateProductStatus.
func (mr *MockQuerierMockRecorder) UpdateProductStatus(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductStatus", reflect.TypeOf((*MockQuerier)(nil).UpdateProductStatus), arg0, arg1)
}
