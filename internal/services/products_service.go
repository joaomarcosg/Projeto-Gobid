package services

import (
	"context"
	"time"

	"github.com/google/uuid"
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
