package controller

import (
	"net/http/httptest"
	"os"
	"testing"

	mockservice "github.com/djudju12/ms-products/service/mock"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

type TestProductController struct {
	ctrl           *gomock.Controller
	productService *mockservice.MockProductService
	server         *Server
	recorder       *httptest.ResponseRecorder
	url            string
}

func NewTest(t *testing.T, url string) *TestProductController {
	ctrl := gomock.NewController(t)
	productService := mockservice.NewMockProductService(ctrl)
	productController := New(productService)
	server := NewServer(productController)
	recorder := httptest.NewRecorder()

	return &TestProductController{
		ctrl:           ctrl,
		productService: productService,
		server:         server,
		recorder:       recorder,
		url:            url,
	}
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
