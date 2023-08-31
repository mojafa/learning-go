package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mojafa/building-web-app-in-go-fundamentals/optmizing-cache-using-application-config/pkg/config"
	"github.com/mojafa/building-web-app-in-go-fundamentals/optmizing-cache-using-application-config/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
