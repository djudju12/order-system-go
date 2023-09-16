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

func TestGetProduct(t *testing.T) {
	product := createRandomProduct(t)

	product2, err := testQueries.GetProduct(context.Background(), product.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product, product2)
}

func TestListProduct(t *testing.T) {
	n := 10

	for i := 0; i < n; i++ {
		createRandomProduct(t)
	}

	arg := ListProductsParams{
		Limit:  int32(n / 2),
		Offset: int32(n / 2),
	}

	products, err := testQueries.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, n/2)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}
