package main

import (
	"net/http"

	"github.com/Coding-Bruh/bookings/pkg/config"
	"github.com/Coding-Bruh/bookings/pkg/handlers"

	//"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(mux.Middlewares()...)
	//mux.Use(WriteToConsole)
	mux.Use(noSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
