package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Gobid/internal/store"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type MockUserStore struct{}

func (m *MockUserStore) CreateUser(
	ctx context.Context,
	userName,
	email string,
	password []byte,
	bio string) (uuid.UUID, error) {
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	return id, nil
}

func (m *MockUserStore) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	return id, nil
}

func (m *MockUserStore) GetUserByEmail(ctx context.Context, email string) (store.User, error) {
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")

	hash, _ := bcrypt.GenerateFromPassword([]byte("Senha123456"), bcrypt.DefaultCost)

	return store.User{
		ID:           id,
		UserName:     "Marcos",
		Email:        "marcos@gmail.com",
		PasswordHash: hash,
		Bio:          "Hello, world!",
	}, nil
}

func (m *MockUserStore) GetUserById(ctx context.Context, id uuid.UUID) (store.User, error) {
	return store.User{}, nil
}

func TestCreateUser(t *testing.T) {
	mockStore := MockUserStore{}
	userService := NewUserService(&mockStore)

	id, err := userService.CreateUser(context.Background(), "Marcos", "marcos@gmail.com", "Senha123456", "Hello, world!")

	assert.NoError(t, err)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", id.String())

}

func TestAuthenticateUser(t *testing.T) {
	mockStore := MockUserStore{}
	userService := NewUserService(&mockStore)

	id, err := userService.AuthenticateUser(context.Background(), "marcos@gmail.com", "Senha123456")

	assert.NoError(t, err)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", id.String())

}

func TestGetUserByEmail(t *testing.T) {
	mockStore := MockUserStore{}
	userService := NewUserService(&mockStore)

	user, err := userService.Store.GetUserByEmail(context.Background(), "marcos@gmail.com")
	assert.NoError(t, err)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", user.ID.String())
	assert.Equal(t, "Marcos", user.UserName)
	assert.Equal(t, "marcos@gmail.com", user.Email)
	assert.Equal(t, "Hello, world!", user.Bio)

}
