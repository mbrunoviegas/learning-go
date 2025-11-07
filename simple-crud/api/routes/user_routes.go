package routes

import (
	"simple-crud/api/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(handler *handlers.UserHandler) chi.Router {
	r := chi.NewRouter()
	r.Route("/users", func(sr chi.Router) {
		sr.Post("/", handler.CreateUser)
		sr.Get("/", handler.ListUsers)
		sr.Get("/{id}", handler.GetUser)
		sr.Put("/{id}", handler.UpdateUser)
		sr.Delete("/{id}", handler.DeleteUser)
	})

	return r
}
