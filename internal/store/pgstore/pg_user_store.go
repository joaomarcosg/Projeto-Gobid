package pgstore

import "github.com/jackc/pgx/v5/pgxpool"

type PGUserStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGUserStore(pool *pgxpool.Pool) PGUserStore {
	return PGUserStore{
		Queries: New(pool),
		Pool:    pool,
	}
}
