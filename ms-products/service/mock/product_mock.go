// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/djudju12/ms-products/service (interfaces: ProductService)
//
// Generated by this command:
//
//	mockgen -package mockservice -destination service/mock/product_mock.go github.com/djudju12/ms-products/service ProductService
//
// Package mockservice is a generated GoMock package.
package mockservice

import (
	context "context"
	reflect "reflect"

	model "github.com/djudju12/ms-products/model"
	gomock "go.uber.org/mock/gomock"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductService) CreateProduct(arg0 context.Context, arg1 model.CreateProductRequest) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductServiceMockRecorder) CreateProduct(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductService)(nil).CreateProduct), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockProductService) DeleteProduct(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductServiceMockRecorder) DeleteProduct(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductService)(nil).DeleteProduct), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockProductService) GetProduct(arg0 context.Context, arg1 int32) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockProductServiceMockRecorder) GetProduct(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductService)(nil).GetProduct), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockProductService) ListProducts(arg0 context.Context, arg1 model.ListProductsRquest) ([]*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockProductServiceMockRecorder) ListProducts(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockProductService)(nil).ListProducts), arg0, arg1)
}
