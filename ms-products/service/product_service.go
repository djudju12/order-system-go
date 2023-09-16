package service

import (
	"database/sql"
	"net/http"

	"github.com/djudju12/order-system/ms-products/model"
	dbproducts "github.com/djudju12/order-system/ms-products/repository/sqlc"
	"github.com/gin-gonic/gin"
)

type ProductService struct {
	repository *dbproducts.Queries
	db         *sql.DB
}

func NewProcutService(db *sql.DB) *ProductService {
	return &ProductService{
		db:         db,
		repository: dbproducts.New(db),
	}
}

func (ps *ProductService) GetProduct(ctx *gin.Context) {
	var req model.GetProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := ps.repository.GetProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (ps *ProductService) CreateProduct(ctx *gin.Context) {
	var req model.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := dbproducts.CreateProductParams{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	product, err := ps.repository.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
