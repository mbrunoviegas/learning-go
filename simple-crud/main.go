package main

import (
	"log/slog"
	"net/http"
	"simple-crud/api/handlers"
	"simple-crud/api/handlers/validators"
	"simple-crud/api/routes"
	"simple-crud/services"
	"simple-crud/storage"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			storage.NewUserDb,
			services.NewUserService,
			validators.NewUserValidator,
			handlers.NewUserHandler,
		),
		fx.Invoke(registerRoutes),
	)

	app.Run()

}

func registerRoutes(handler *handlers.UserHandler) {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Mount("/api", routes.UserRoutes(handler))

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      r,
	}

	slog.Info("Server is running on https://localhost:8080")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			slog.Error("Fatal error", "error", err)
		}
	}()
}
