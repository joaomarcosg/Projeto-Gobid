package services

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
	"github.com/stretchr/testify/assert"
)

type MockProductStore struct{}

func (m *MockProductStore) CreateProduct(
	ctx context.Context,
	sellerID uuid.UUID,
	productName,
	description string,
	baseprice float64,
	auctionEnd time.Time,
) (uuid.UUID, error) {
	productID := uuid.New()
	return productID, nil
}

func (m *MockProductStore) GetProductById(ctx context.Context, productID uuid.UUID) (store.Product, error) {
	return store.Product{
		ID:          productID,
		SellerID:    uuid.New(),
		ProductName: "Producto de teste",
		Description: "Descrição do produto",
		Baseprice:   99.99,
		AuctionEnd:  time.Now().Add(2 * time.Hour),
	}, nil
}

func TestCreateProduct(t *testing.T) {

	mockStore := MockProductStore{}
	productService := NewProductService(&mockStore)

	ctx := context.Background()
	sellerID := uuid.New()
	productName := "Produto de teste"
	description := "Descrição do produto"
	baseprice := 99.99
	auctionEnd := time.Now().Add(2 * time.Hour)

	productID, err := productService.CreateProduct(ctx, sellerID, productName, description, baseprice, auctionEnd)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, productID)

}

func TestGetProductById(t *testing.T) {

	mockStore := MockProductStore{}
	productService := NewProductService(&mockStore)

	ctx := context.Background()
	productID := uuid.New()

	product, err := productService.Store.GetProductById(ctx, productID)

	assert.NoError(t, err)
	assert.Equal(t, productID, product.ID)

}
