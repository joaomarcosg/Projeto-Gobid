package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
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
