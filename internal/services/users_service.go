package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicatedEmailOrUserName = errors.New("username or email already exists")
	ErrInvalidCredentials        = errors.New("invalid credentials")
)

type UserService struct {
	Store store.UserStore
}

func NewUserService(store store.UserStore) *UserService {
	return &UserService{
		Store: store,
	}
}

func (us *UserService) CreateUser(
	ctx context.Context,
	userName,
	email,
	password,
	bio string,
) (uuid.UUID, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := us.Store.CreateUser(ctx, userName, email, hash, bio)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicatedEmailOrUserName
		}

	}

	return id, nil
}

func (us *UserService) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {

	user, err := us.Store.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrInvalidCredentials
		}
		return uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return uuid.UUID{}, err
		}

		return uuid.UUID{}, err
	}

	return user.ID, nil

}
