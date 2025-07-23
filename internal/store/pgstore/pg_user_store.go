package pgstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
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

func (pgu *PGUserStore) CreateUser(
	ctx context.Context,
	userName,
	email string,
	password []byte,
	bio string) (uuid.UUID, error) {

	id, err := pgu.Queries.CreateUser(ctx, CreateUserParams{
		UserName:     userName,
		Email:        email,
		PasswordHash: password,
		Bio:          bio,
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil

}

func (pgu *PGUserStore) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {

	user, err := pgu.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return uuid.UUID{}, nil
	}

	return user.ID, nil

}

func (pgu *PGUserStore) GetUserByEmail(ctx context.Context, email string) (store.User, error) {

	user, err := pgu.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return store.User{}, nil
	}

	return store.User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Bio:          user.Bio,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}, nil

}

func (pgu *PGUserStore) GetUserById(ctx context.Context, id uuid.UUID) (store.User, error) {

	user, err := pgu.Queries.GetUserById(ctx, id)
	if err != nil {
		return store.User{}, err
	}

	return store.User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Bio:          user.Bio,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}, nil

}
