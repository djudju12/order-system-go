package controller

import (
	"github.com/djudju12/order-system/ms-products/service"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	getProduct(ctx *gin.Context)
	createProduct(ctx *gin.Context)
	// ListProducts(ctx *gin.Context)
	// DeleteProduct(ctx *gin.Context)
}

type productController struct {
	service *service.ProductService
}

func New(service *service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

func (pc *productController) getProduct(ctx *gin.Context) {
	pc.service.GetProduct(ctx)
}

func (pc *productController) createProduct(ctx *gin.Context) {
	pc.service.CreateProduct(ctx)
}
