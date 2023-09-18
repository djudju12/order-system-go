package model

import db "github.com/djudju12/ms-products/db/sqlc"

type Product struct {
	db.Product
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

func (req *ListProductsRquest) ToDB() db.ListProductsParams {
	return db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
}

func (req *CreateProductRequest) ToDB() db.CreateProductParams {
	return db.CreateProductParams{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}
}

func ProductDbToModel(product db.Product) *Product {
	return &Product{
		Product: product,
	}
}

func ListProductsDbToModel(products []db.Product) []*Product {
	result := make([]*Product, 0)
	for _, product := range products {
		result = append(result, ProductDbToModel(product))
	}

	return result
}
