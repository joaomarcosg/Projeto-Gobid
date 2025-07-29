package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	BidderID  uuid.UUID `json:"bidder_id"`
	BidAmount float64   `json:"bid_amount"`
	CreatedAt time.Time `json:"created_at"`
}

type BidStore interface {
	CreateBid(ctx context.Context, productID, bidderID uuid.UUID, bidAmount float64) (Bid, error)
	GetBidsByProduct(ctx context.Context, productID uuid.UUID) ([]Bid, error)
	GetHighestBidByProductId(ctx context.Context, productID uuid.UUID) (Bid, error)
}
