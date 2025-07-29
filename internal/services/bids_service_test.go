package services

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
	"github.com/stretchr/testify/assert"
)

type MockBidStore struct{}

func (m *MockBidStore) CreateBid(ctx context.Context, productID, bidderID uuid.UUID, bidAmount float64) (store.Bid, error) {
	return store.Bid{
		ID:        uuid.New(),
		ProductID: productID,
		BidderID:  bidderID,
		BidAmount: bidAmount,
		CreatedAt: time.Now(),
	}, nil
}

func (m *MockBidStore) GetBidsByProduct(ctx context.Context, productID uuid.UUID) ([]store.Bid, error) {
	return []store.Bid{}, nil
}

func (m *MockBidStore) GetHighestBidByProductId(ctx context.Context, productID uuid.UUID) (store.Bid, error) {
	id := uuid.New()
	bidID := uuid.New()
	return store.Bid{
		ID:        id,
		ProductID: productID,
		BidderID:  bidID,
		BidAmount: 99.99,
		CreatedAt: time.Now(),
	}, nil
}

func TestCreateBid(t *testing.T) {

	mockStore := MockBidStore{}
	bidService := NewBidService(&mockStore)

	ctx := context.Background()
	productID := uuid.New()
	bidderID := uuid.New()
	bidAmount := 99.99

	bid, err := bidService.Store.CreateBid(ctx, productID, bidderID, bidAmount)

	id := bid.ID

	assert.NoError(t, err)
	assert.Equal(t, bidderID, bid.BidderID)
	assert.Equal(t, productID, bid.ProductID)
	assert.Equal(t, bidAmount, bid.BidAmount)
	assert.Equal(t, id, bid.ID)
}

func TestHighestBidByProductId(t *testing.T) {

	mockStore := MockBidStore{}
	bidService := NewBidService(&mockStore)

	ctx := context.Background()
	productID := uuid.New()

	highestBid, err := bidService.Store.GetHighestBidByProductId(ctx, productID)

	assert.NoError(t, err)
	assert.Equal(t, productID, highestBid.ProductID)

}
