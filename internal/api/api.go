package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/joaomarcosg/Projeto-Gobid/internal/services"
)

type Api struct {
	Router      *chi.Mux
	UserService services.UserService
}
