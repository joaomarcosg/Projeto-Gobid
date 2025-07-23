package pgstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

func (pgu *PGUserStore) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {

	id, err := pgu.Queries.CreateUser(ctx, CreateUserParams{
		UserName:     userName,
		Email:        email,
		PasswordHash: []byte(password),
		Bio:          bio,
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil

}
