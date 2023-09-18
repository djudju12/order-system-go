package controller

import (
	"database/sql"
	"net/http"

	"github.com/djudju12/order-system/ms-products/model"
	"github.com/djudju12/order-system/ms-products/service"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	getProduct(ctx *gin.Context)
	createProduct(ctx *gin.Context)
	listProducts(ctx *gin.Context)
	deleteProduct(ctx *gin.Context)
}

type productController struct {
	service service.ProductService
}

func New(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

func (pc *productController) getProduct(ctx *gin.Context) {
	var req model.GetProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	p, err := pc.service.GetProduct(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, p)
}

func (pc *productController) createProduct(ctx *gin.Context) {
	var req model.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := pc.service.CreateProduct(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (pc *productController) listProducts(ctx *gin.Context) {
	var req model.ListProductsRquest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	products, err := pc.service.ListProducts(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (pc *productController) deleteProduct(ctx *gin.Context) {
	var req model.DeleteProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := pc.service.DeleteProduct(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
