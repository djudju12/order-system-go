package service

import (
	"context"
	"database/sql"

	"github.com/djudju12/ms-products/model"
	dbproducts "github.com/djudju12/ms-products/repository/sqlc"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetProduct(ctx context.Context, productID int32) (*model.Product, error)
	CreateProduct(ctx *gin.Context, req model.CreateProductRequest) (*model.Product, error)
	ListProducts(ctx *gin.Context, req model.ListProductsRquest) ([]*model.Product, error)
	DeleteProduct(ctx *gin.Context, productID int32) error
}

type productService struct {
	repository *dbproducts.Queries
	db         *sql.DB
}

var _ ProductService = (*productService)(nil)

func NewProcutService(db *sql.DB) ProductService {
	return &productService{
		db:         db,
		repository: dbproducts.New(db),
	}
}

func (ps *productService) GetProduct(ctx context.Context, productID int32) (*model.Product, error) {
	product, err := ps.repository.GetProduct(ctx, int32(productID))
	if err != nil {
		return nil, err
	}

	return &model.Product{Product: product}, nil
}

func (ps *productService) CreateProduct(ctx *gin.Context, req model.CreateProductRequest) (*model.Product, error) {
	arg := dbproducts.CreateProductParams{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	product, err := ps.repository.CreateProduct(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &model.Product{Product: product}, nil
}

func (ps *productService) ListProducts(ctx *gin.Context, req model.ListProductsRquest) ([]*model.Product, error) {
	arg := dbproducts.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := ps.repository.ListProducts(ctx, arg)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Product, 0)
	for _, product := range products {
		result = append(result, &model.Product{Product: product})
	}

	return result, nil
}

func (ps *productService) DeleteProduct(ctx *gin.Context, productID int32) error {
	_, err := ps.repository.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	return ps.repository.DeleteProduct(ctx, productID)
}
