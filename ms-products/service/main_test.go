package service

import (
	"testing"

	mockdb "github.com/djudju12/ms-products/db/mock"
	"go.uber.org/mock/gomock"
)

type TestProductService struct {
	ctrl       *gomock.Controller
	repository *mockdb.MockQuerier
	service    ProductService
}

func NewTest(t *testing.T) *TestProductService {
	ctrl := gomock.NewController(t)
	ctrl.Finish()
	repository := mockdb.NewMockQuerier(ctrl)
	sevice := NewProductService(repository)

	return &TestProductService{
		ctrl:       ctrl,
		repository: repository,
		service:    sevice,
	}
}
