package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
)

type ProductService struct {
	Store store.ProductStore
}

func NewProductService(store store.ProductStore) *ProductService {
	return &ProductService{
		Store: store,
	}
}

func (ps *ProductService) CreateProduct(
	ctx context.Context,
	sellerID uuid.UUID,
	productName,
	description string,
	baseprice float64,
	auctionEnd time.Time,
) (uuid.UUID, error) {

	productID, err := ps.Store.CreateProduct(
		ctx,
		sellerID,
		productName,
		description,
		baseprice,
		auctionEnd,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return productID, nil

}

var ErrProductNotFound = errors.New("product not found")

func (ps *ProductService) GetProductById(ctx context.Context, productID uuid.UUID) (store.Product, error) {

	product, err := ps.Store.GetProductById(ctx, productID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return store.Product{}, ErrProductNotFound
		}
		return store.Product{}, err
	}

	return store.Product{
		ID:          product.ID,
		SellerID:    product.ID,
		ProductName: product.ProductName,
		Description: product.Description,
		Baseprice:   product.Baseprice,
		AuctionEnd:  product.AuctionEnd,
		IsSold:      product.IsSold,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil

}
