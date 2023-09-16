package dbproducts

import (
	"context"
	"testing"

	"github.com/djudju12/order-system/ms-products/utils"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Name:        utils.RandomProductName(),
		Price:       utils.RandomProductPrice(),
		Description: utils.RandomProductDescription(),
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, product.Name, arg.Name)
	require.Equal(t, product.Price, arg.Price)
	require.Equal(t, product.Description, arg.Description)

	require.NotZero(t, product.ID)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}
