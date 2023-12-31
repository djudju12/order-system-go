// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"context"
)

type Querier interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	GetProduct(ctx context.Context, id int32) (Product, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	UpdateProductStatus(ctx context.Context, arg UpdateProductStatusParams) (Product, error)
}

var _ Querier = (*Queries)(nil)
