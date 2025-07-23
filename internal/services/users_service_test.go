package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
	"github.com/stretchr/testify/assert"
)

type MockUserStore struct{}

func (m *MockUserStore) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	return id, nil
}

func (m *MockUserStore) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (m *MockUserStore) GetUserByEmail(ctx context.Context, email string) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (m *MockUserStore) GetUserById(ctx context.Context, id uuid.UUID) (store.User, error) {
	return store.User{}, nil
}

func TestCreateUser(t *testing.T) {
	mockStore := MockUserStore{}
	userService := NewUserService(&mockStore)

	id, err := userService.CreateUser(context.Background(), "Marcos", "marcos@gmail.com", "Senha123456", "Hellom world!")

	assert.NoError(t, err)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", id)

}
