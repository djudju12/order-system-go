package model

import dbproducts "github.com/djudju12/order-system/ms-products/repository/sqlc"

type Product struct {
	dbproducts.Product
}

type ListProductsRquest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       string `json:"price" binding:"required,price"`
	Description string `json:"description" binding:"required"`
}

type GetProductRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type DeleteProductRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}
