package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
)

var (
	ps             *ProductService
	ErrBidIsTooLow = errors.New("the bid values is too low")
)

type BidService struct {
	Store store.BidStore
}

func NewBidService(store store.BidStore) *BidService {
	return &BidService{
		Store: store,
	}
}

func (bs *BidService) PlaceBid(ctx context.Context, productID, bidderID uuid.UUID, amount float64) (store.Bid, error) {

	product, err := ps.GetProductById(ctx, productID)
	if err != nil {
		return store.Bid{}, err
	}

	highestBid, err := bs.Store.GetHighestBidByProductId(ctx, product.ID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return store.Bid{}, err
		}
	}

	if product.Baseprice >= amount || highestBid.BidAmount >= amount {
		return store.Bid{}, ErrBidIsTooLow
	}

	highestBid, err = bs.Store.CreateBid(ctx, productID, bidderID, amount)
	if err != nil {
		return store.Bid{}, err
	}

	return highestBid, nil

}
