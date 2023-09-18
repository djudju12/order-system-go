package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/djudju12/ms-products/model"
	mockservice "github.com/djudju12/ms-products/service/mock"
	"github.com/djudju12/ms-products/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetProduct(t *testing.T) {
	product := RandomProduct()

	testCases := []struct {
		name          string
		productID     int32
		buildStubs    func(service *mockservice.MockProductService)
		checkResponse func(t *testing.T, recored *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productID: product.ID,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireMatchProduct(t, recorder.Body, product)
			},
		},
		{
			name:      "Not Found",
			productID: product.ID,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(&model.Product{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "Bad Request",
			productID: 0,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "Internal Server Error",
			productID: product.ID,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(&model.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// given
			test := NewTest(t, fmt.Sprintf("/products/%d", tC.productID))
			tC.buildStubs(test.productService)

			request, err := http.NewRequest(http.MethodGet, test.url, nil)
			require.NoError(t, err)

			// when
			test.server.router.ServeHTTP(test.recorder, request)

			// then
			tC.checkResponse(t, test.recorder)
		})
	}
}

func TestCreateProduct(t *testing.T) {
	product := RandomProduct()
	request := model.CreateProductRequest{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}

	testCases := []struct {
		name          string
		request       model.CreateProductRequest
		buildStubs    func(service *mockservice.MockProductService)
		checkResponse func(t *testing.T, recored *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			request: request,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(request)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireMatchProduct(t, recorder.Body, product)
			},
		},
		{
			name:    "Bad Request",
			request: model.CreateProductRequest{},

			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					CreateProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:    "Internal Server Error",
			request: request,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(request)).
					Times(1).
					Return(&model.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// given
			test := NewTest(t, "/products")
			tC.buildStubs(test.productService)

			request, err := http.NewRequest(http.MethodPost, test.url, toReader(t, tC.request))
			require.NoError(t, err)

			// when
			test.server.router.ServeHTTP(test.recorder, request)

			// then
			tC.checkResponse(t, test.recorder)
		})
	}
}

func TestListProduct(t *testing.T) {
	n := 5
	products := make([]*model.Product, 0)
	for i := 0; i < n; i++ {
		products = append(products, RandomProduct())
	}

	testCases := []struct {
		name          string
		pageID        int32
		pageSize      int32
		buildStubs    func(service *mockservice.MockProductService)
		checkResponse func(t *testing.T, recored *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			pageID:   1,
			pageSize: int32(n),
			buildStubs: func(service *mockservice.MockProductService) {
				arg := model.ListProductsRquest{
					PageID:   1,
					PageSize: int32(n),
				}

				service.EXPECT().
					ListProducts(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(products, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				returnedProducts := readBody(t, recorder.Body)

				require.Equal(t, http.StatusOK, recorder.Code)
				require.Len(t, returnedProducts, n)

				for _, p := range returnedProducts {
					require.NotEmpty(t, p)
				}
			},
		},
		{
			name:     "Bad Request",
			pageID:   0,
			pageSize: 100,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					ListProducts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:     "Internal Server Error",
			pageID:   1,
			pageSize: int32(n),
			buildStubs: func(service *mockservice.MockProductService) {
				arg := model.ListProductsRquest{
					PageID:   1,
					PageSize: int32(n),
				}

				service.EXPECT().
					ListProducts(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return([]*model.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// given
			url := fmt.Sprintf("/products?page_id=%d&page_size=%d",
				tC.pageID, tC.pageSize)

			test := NewTest(t, url)
			tC.buildStubs(test.productService)

			request, err := http.NewRequest(http.MethodGet, test.url, nil)
			require.NoError(t, err)

			// when
			test.server.router.ServeHTTP(test.recorder, request)

			// then
			tC.checkResponse(t, test.recorder)
		})
	}
}

func TestInactiveProduct(t *testing.T) {
	productID := utils.RandomProductID()
	testCases := []struct {
		name          string
		productID     int32
		buildStubs    func(service *mockservice.MockProductService)
		checkResponse func(t *testing.T, recored *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productID: productID,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					InactiveProduct(gomock.Any(), gomock.Eq(productID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "Not Found",
			productID: productID,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					InactiveProduct(gomock.Any(), gomock.Eq(productID)).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "Bad Request",
			productID: 0,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					InactiveProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "Internal Server Error",
			productID: productID,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					InactiveProduct(gomock.Any(), gomock.Eq(productID)).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// given
			url := fmt.Sprintf("/products/%d", tC.productID)

			test := NewTest(t, url)
			tC.buildStubs(test.productService)

			request, err := http.NewRequest(http.MethodDelete, test.url, nil)
			require.NoError(t, err)

			// when
			test.server.router.ServeHTTP(test.recorder, request)

			// then
			tC.checkResponse(t, test.recorder)
		})
	}
}

func TestUpdateProductStatus(t *testing.T) {
	product := RandomProduct()
	request := model.UpdateProductStatusRequest{
		ID:     product.ID,
		Status: "out_of_stock",
	}

	testCases := []struct {
		name          string
		request       model.UpdateProductStatusRequest
		buildStubs    func(service *mockservice.MockProductService)
		checkResponse func(t *testing.T, recored *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			request: request,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					UpdateProductStatus(gomock.Any(), gomock.Eq(request)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Bad Request",
			request: model.UpdateProductStatusRequest{
				ID:     0,
				Status: "out_of_stock",
			},
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					UpdateProductStatus(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:    "Not Found",
			request: request,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					UpdateProductStatus(gomock.Any(), gomock.Eq(request)).
					Times(1).
					Return(&model.Product{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "Internal Server Error",
			request: request,
			buildStubs: func(service *mockservice.MockProductService) {
				service.EXPECT().
					UpdateProductStatus(gomock.Any(), gomock.Eq(request)).
					Times(1).
					Return(&model.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			// given
			url := "/products"

			test := NewTest(t, url)
			tC.buildStubs(test.productService)

			request, err := http.NewRequest(http.MethodPatch, test.url, toReader(t, tC.request))
			require.NoError(t, err)

			// when
			test.server.router.ServeHTTP(test.recorder, request)

			// then
			tC.checkResponse(t, test.recorder)
		})
	}
}

func RandomProduct() *model.Product {
	return &model.Product{
		ID:          utils.RandomProductID(),
		Name:        utils.RandomProductName(),
		Price:       utils.RandomProductPrice(),
		Description: utils.RandomProductDescription(),
	}
}

func requireMatchProduct(t *testing.T, body *bytes.Buffer, product *model.Product) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var responseProduct model.Product
	err = json.Unmarshal(data, &responseProduct)
	require.NoError(t, err)

	require.Equal(t, product, &responseProduct)
}

func toReader(t *testing.T, body any) io.Reader {
	data, err := json.Marshal(body)
	require.NoError(t, err)

	return bytes.NewReader(data)
}

func readBody(t *testing.T, boyd *bytes.Buffer) []*model.Product {
	data, err := io.ReadAll(boyd)
	require.NoError(t, err)

	var products []*model.Product
	err = json.Unmarshal(data, &products)
	require.NoError(t, err)

	return products
}
