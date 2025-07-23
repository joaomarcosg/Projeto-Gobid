package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
	"golang.org/x/crypto/bcrypt"
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

	_, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := us.Store.CreateUser(ctx, userName, email, password, bio)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
