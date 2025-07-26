package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	SellerID    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
	IsSold      bool      `json:"is_sold"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductStore interface {
	CreateProduct(
		ctx context.Context,
		sellerID uuid.UUID,
		productName,
		description string,
		baseprice float64,
		auctionEnd time.Time,
	) (uuid.UUID, error)
	GetProductById(ctx context.Context, productID uuid.UUID) (Product, error)
}
