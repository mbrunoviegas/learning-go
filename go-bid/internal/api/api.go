package api

import (
	"go-bid/internal/services"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

// Defines handlers that we would use
type Api struct {
	Router      *chi.Mux
	UserService services.UserService
	Sessions    *scs.SessionManager
}
