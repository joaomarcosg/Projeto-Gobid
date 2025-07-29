package pgstore

import (
	"github.com/jackc/pgx/v5/pgxpool"
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
