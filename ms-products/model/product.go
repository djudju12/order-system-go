package model

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       string `json:"price" binding:"required,price"`
	Description string `json:"descripion" binding:"required"`
}

type GetProductRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}
