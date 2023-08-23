package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/mojafa/building-web-app-in-go-fundamentals/optmizing-cache-using-application-config/pkg/config"
	"github.com/mojafa/building-web-app-in-go-fundamentals/optmizing-cache-using-application-config/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
