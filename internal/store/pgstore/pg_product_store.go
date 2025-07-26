package pgstore

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
)

type PGProductStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGProductStore(pool *pgxpool.Pool) *PGProductStore {
	return &PGProductStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func (pgp *PGProductStore) CreateProduct(
	ctx context.Context,
	sellerID uuid.UUID,
	productName,
	description string,
	baseprice float64,
	auctionEnd time.Time,
) (uuid.UUID, error) {
	productID, err := pgp.Queries.CreateProduct(ctx, CreateProductParams{
		SellerID:    sellerID,
		ProductName: productName,
		Description: description,
		Baseprice:   baseprice,
		AuctionEnd:  auctionEnd,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	return productID, nil
}

func (pgp *PGProductStore) GetProductById(ctx context.Context, productID uuid.UUID) (store.Product, error) {

	product, err := pgp.Queries.GetProductById(ctx, productID)
	if err != nil {
		return store.Product{}, err
	}

	return store.Product{
		ID:          product.ID,
		SellerID:    product.SellerID,
		ProductName: product.ProductName,
		Description: product.Description,
		Baseprice:   product.Baseprice,
		AuctionEnd:  product.AuctionEnd,
		IsSold:      product.IsSold,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}
