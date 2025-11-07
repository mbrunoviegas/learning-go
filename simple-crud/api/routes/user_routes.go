package routes

import (
	"simple-crud/api/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(handler *handlers.UserHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/users", handler.CreateUser)

	return r
}
