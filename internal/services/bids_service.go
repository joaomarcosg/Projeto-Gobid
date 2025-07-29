package services

import "github.com/joaomarcosg/Projeto-Gobid/internal/store"

type BidService struct {
	Store store.BidStore
}

func NewBidService(store store.BidStore) *BidService {
	return &BidService{
		Store: store,
	}
}
