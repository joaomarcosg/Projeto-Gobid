package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
)

type MockUserStore struct{}

func (m *MockUserStore) CreateUser(ctx context.Context, userName, email, password, bio string) (store.User, error) {
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	return store.User{
		ID:           id,
		UserName:     userName,
		Email:        email,
		PasswordHash: []byte("mocked_hash_password"),
		Bio:          bio,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
