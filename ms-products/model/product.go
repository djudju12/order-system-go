package model

import (
	"time"

	db "github.com/djudju12/ms-products/db/sqlc"
)

type Product struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Price       string    `json:"price"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ProductDbToModel(product db.Product) *Product {
	return &Product{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ProductModelToDb(product *Product) *db.Product {
	return &db.Product{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ListProductsDbToModel(products []db.Product) []*Product {
	result := make([]*Product, 0)
	for _, product := range products {
		result = append(result, ProductDbToModel(product))
	}

	return result
}

type ListProductsRquest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (req *ListProductsRquest) ToDB() db.ListProductsParams {
	return db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
}

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       string `json:"price" binding:"required,price"`
	Description string `json:"description" binding:"required"`
}

func (req *CreateProductRequest) ToDB() db.CreateProductParams {
	return db.CreateProductParams{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}
}

type UpdateProductStatusRequest struct {
	ID     int32  `json:"id" binding:"required,min=1"`
	Status string `json:"status" binding:"required,oneof=out_of_stock available"`
}

func (req *UpdateProductStatusRequest) ToDB() db.UpdateProductStatusParams {
	return db.UpdateProductStatusParams{
		ID:     req.ID,
		Status: req.Status,
	}
}

type GetProductRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type DeleteProductRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}
