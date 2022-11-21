package main

import (
	"net/http"

	"github.com/ShehabEl-DeenAlalkamy/residencia/pkg/config"
	"github.com/ShehabEl-DeenAlalkamy/residencia/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(_ *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// now when we put that middleware in place because we've declared this before the routes, it's actually available
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// http.FileServer() takes a root FileSystem and returns a Handler
	fileServer := http.FileServer(http.Dir("./static/"))

	// takes the url that go gets and modifies it from the web server request into something it knows how to handle
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
