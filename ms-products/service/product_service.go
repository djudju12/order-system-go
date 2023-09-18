package service

import (
	"context"

	db "github.com/djudju12/ms-products/db/sqlc"
	"github.com/djudju12/ms-products/model"
)

type ProductService interface {
	GetProduct(ctx context.Context, productID int32) (*model.Product, error)
	CreateProduct(ctx context.Context, req model.CreateProductRequest) (*model.Product, error)
	ListProducts(ctx context.Context, req model.ListProductsRquest) ([]*model.Product, error)
	DeleteProduct(ctx context.Context, productID int32) error
}

type productService struct {
	repository db.Querier
}

var _ ProductService = (*productService)(nil)

func NewProductService(repository db.Querier) ProductService {
	return &productService{
		repository: repository,
	}
}

func (ps *productService) GetProduct(ctx context.Context, productID int32) (*model.Product, error) {
	product, err := ps.repository.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	return model.ProductDbToModel(product), nil
}

func (ps *productService) CreateProduct(ctx context.Context, req model.CreateProductRequest) (*model.Product, error) {
	arg := req.ToDB()

	product, err := ps.repository.CreateProduct(ctx, arg)
	if err != nil {
		return nil, err
	}

	return model.ProductDbToModel(product), nil
}

func (ps *productService) ListProducts(ctx context.Context, req model.ListProductsRquest) ([]*model.Product, error) {
	arg := req.ToDB()

	products, err := ps.repository.ListProducts(ctx, arg)
	if err != nil {
		return nil, err
	}

	return model.ListProductsDbToModel(products), nil
}

func (ps *productService) DeleteProduct(ctx context.Context, productID int32) error {
	_, err := ps.repository.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	return ps.repository.DeleteProduct(ctx, productID)
}
