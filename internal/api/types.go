package api

import "github.com/google/uuid"

type SignupUserResponse struct {
	UserID uuid.UUID `json:"user_id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
