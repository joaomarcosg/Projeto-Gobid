package pgstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
)

type PGBidStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGBidStore(pool *pgxpool.Pool) *PGBidStore {
	return &PGBidStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func (pgb *PGBidStore) CreateBid(
	ctx context.Context,
	productID,
	bidderID uuid.UUID,
	bidAmount float64,
) (store.Bid, error) {

	bid, err := pgb.Queries.CreateBid(ctx, CreateBidParams{
		ProductID: productID,
		BidderID:  bidderID,
		BidAmount: bidAmount,
	})

	if err != nil {
		return store.Bid{}, err
	}

	return store.Bid{
		ID:        bid.ID,
		ProductID: bid.ProductID,
		BidderID:  bid.BidderID,
		BidAmount: bid.BidAmount,
	}, nil

}

func (pgb *PGBidStore) GetBidsByProduct(ctx context.Context, productID uuid.UUID) ([]store.Bid, error) {

	dbBids, err := pgb.Queries.GetBidsByProduct(ctx, productID)
	if err != nil {
		return []store.Bid{}, err
	}

	bids := make([]store.Bid, len(dbBids))

	for i, b := range dbBids {
		bids[i] = store.Bid{
			ID:        b.ID,
			ProductID: b.ProductID,
			BidderID:  b.BidderID,
			BidAmount: b.BidAmount,
			CreatedAt: b.CreatedAt,
		}
	}
	return bids, nil

}

func (pgb *PGBidStore) GetHighestBidByProductId(ctx context.Context, productID uuid.UUID) (store.Bid, error) {

	bid, err := pgb.Queries.GetHighestBidByProductId(ctx, productID)
	if err != nil {
		return store.Bid{}, err
	}

	return store.Bid{
		ID:        bid.ID,
		ProductID: bid.ProductID,
		BidderID:  bid.BidderID,
		BidAmount: bid.BidAmount,
		CreatedAt: bid.CreatedAt,
	}, nil

}
