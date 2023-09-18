package service

import (
	"context"
	"errors"
	"testing"
	"time"

	mockdb "github.com/djudju12/ms-products/db/mock"
	db "github.com/djudju12/ms-products/db/sqlc"
	"github.com/djudju12/ms-products/model"
	"github.com/djudju12/ms-products/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetProduct(t *testing.T) {
	product := RandomProduct()
	testCases := []struct {
		name        string
		description string
		productID   int32
		buildStubs  func(repository *mockdb.MockQuerier)
		check       func(t *testing.T, product *model.Product, err error)
	}{
		{
			name:        "Happy case",
			productID:   product.ID,
			description: "call GetProduct with a valid productID",
			buildStubs: func(repository *mockdb.MockQuerier) {
				repository.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(product, nil)
			},
			check: func(t *testing.T, productModel *model.Product, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, productModel)
				require.Equal(t, productModel, model.ProductDbToModel(product))
			},
		},
		{
			name:        "Repository returns an error",
			productID:   product.ID,
			description: "call GetProduct and repository returns an error",
			buildStubs: func(repository *mockdb.MockQuerier) {
				repository.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Product{}, errors.New("some error"))
			},
			check: func(t *testing.T, productModel *model.Product, err error) {
				require.Error(t, err)
				require.Empty(t, productModel)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			test := NewTest(t)
			tC.buildStubs(test.repository)

			p, err := test.service.GetProduct(context.Background(), tC.productID)

			tC.check(t, p, err)
		})
	}
}
func TestCreateProduct(t *testing.T) {
	product := RandomProduct()

	req := model.CreateProductRequest{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}

	testCases := []struct {
		name       string
		request    model.CreateProductRequest
		buildStubs func(repository *mockdb.MockQuerier)
		check      func(t *testing.T, product *model.Product, err error)
	}{
		{
			name:    "Happy case",
			request: req,
			buildStubs: func(repository *mockdb.MockQuerier) {
				expectedArg := req.ToDB()

				repository.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(expectedArg)).
					Times(1).
					Return(product, nil)
			},
			check: func(t *testing.T, productModel *model.Product, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, productModel)
				require.Equal(t, productModel, model.ProductDbToModel(product))
			},
		},
		{
			name:    "Repository returns an error",
			request: model.CreateProductRequest{},
			buildStubs: func(repository *mockdb.MockQuerier) {
				repository.EXPECT().
					CreateProduct(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Product{}, errors.New("some error"))
			},
			check: func(t *testing.T, productModel *model.Product, err error) {
				require.Error(t, err)
				require.Empty(t, productModel)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			test := NewTest(t)
			tC.buildStubs(test.repository)

			p, err := test.service.CreateProduct(context.Background(), tC.request)

			tC.check(t, p, err)
		})
	}
}

func TestListProduct(t *testing.T) {
	n := 5
	products := make([]db.Product, 0)
	for i := 0; i < n; i++ {
		products = append(products, RandomProduct())
	}

	req := model.ListProductsRquest{
		PageID:   1,
		PageSize: int32(n),
	}

	testCases := []struct {
		name       string
		request    model.ListProductsRquest
		buildStubs func(repository *mockdb.MockQuerier)
		check      func(t *testing.T, productsModel []*model.Product, err error)
	}{
		{
			name:    "Happy case",
			request: req,
			buildStubs: func(repository *mockdb.MockQuerier) {
				expectedArg := req.ToDB()

				repository.EXPECT().
					ListProducts(gomock.Any(), gomock.Eq(expectedArg)).
					Times(1).
					Return(products, nil)
			},
			check: func(t *testing.T, productsModel []*model.Product, err error) {
				require.NoError(t, err)
				require.Len(t, productsModel, n)
				for _, pm := range productsModel {
					require.NotEmpty(t, pm)
				}
			},
		},
		{
			name:    "Repository returns an error",
			request: model.ListProductsRquest{},
			buildStubs: func(repository *mockdb.MockQuerier) {
				repository.EXPECT().
					ListProducts(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Product{}, errors.New("some error"))
			},
			check: func(t *testing.T, productsModel []*model.Product, err error) {
				require.Error(t, err)
				require.Empty(t, productsModel)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			test := NewTest(t)
			tC.buildStubs(test.repository)

			p, err := test.service.ListProducts(context.Background(), tC.request)

			tC.check(t, p, err)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	product := RandomProduct()
	request := model.UpdateProductStatusRequest{
		ID:     product.ID,
		Status: "out_of_stock",
	}

	testCases := []struct {
		name       string
		request    model.UpdateProductStatusRequest
		buildStubs func(repository *mockdb.MockQuerier)
		check      func(t *testing.T, productModel *model.Product, err error)
	}{
		{
			name:    "Happy case",
			request: request,
			buildStubs: func(repository *mockdb.MockQuerier) {
				expectedArg := db.UpdateProductStatusParams{
					ID:     request.ID,
					Status: request.Status,
				}

				repository.EXPECT().
					UpdateProductStatus(gomock.Any(), gomock.Eq(expectedArg)).
					Times(1).
					Return(product, nil)
			},
			check: func(t *testing.T, productModel *model.Product, err error) {
				require.NoError(t, err)
				require.Equal(t, productModel, model.ProductDbToModel(product))
			},
		},
		{
			name:    "Repository returns an error",
			request: request,
			buildStubs: func(repository *mockdb.MockQuerier) {
				repository.EXPECT().
					UpdateProductStatus(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Product{}, errors.New("some error"))
			},
			check: func(t *testing.T, productModel *model.Product, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			test := NewTest(t)
			tC.buildStubs(test.repository)

			p, err := test.service.UpdateProductStatus(context.Background(), tC.request)

			tC.check(t, p, err)
		})
	}
}

func TestInactiveProduct(t *testing.T) {
	product := RandomProduct()

	testCases := []struct {
		name       string
		productID  int32
		buildStubs func(repository *mockdb.MockQuerier)
		check      func(t *testing.T, err error)
	}{
		{
			name:      "Happy case",
			productID: product.ID,
			buildStubs: func(repository *mockdb.MockQuerier) {
				expectedArg := db.UpdateProductStatusParams{
					ID:     product.ID,
					Status: "inactive",
				}

				repository.EXPECT().
					UpdateProductStatus(gomock.Any(), gomock.Eq(expectedArg)).
					Times(1).
					Return(db.Product{}, nil)
			},
			check: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name:      "Repository returns an error",
			productID: product.ID,
			buildStubs: func(repository *mockdb.MockQuerier) {
				repository.EXPECT().
					UpdateProductStatus(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Product{}, errors.New("some error"))
			},
			check: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			test := NewTest(t)
			tC.buildStubs(test.repository)

			err := test.service.InactiveProduct(context.Background(), tC.productID)

			tC.check(t, err)
		})
	}
}

func RandomProduct() db.Product {
	return db.Product{
		ID:          utils.RandomProductID(),
		Name:        utils.RandomProductName(),
		Price:       utils.RandomProductPrice(),
		Description: utils.RandomProductDescription(),
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
