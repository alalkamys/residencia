package main

import (
	"net/http"

	"github.com/ShehabEl-DeenAlalkamy/residencia/internal/config"
	"github.com/ShehabEl-DeenAlalkamy/residencia/internal/handlers"
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
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	// http.FileServer() takes a root FileSystem and returns a Handler
	fileServer := http.FileServer(http.Dir("./static/"))

	// takes the url that go gets and modifies it from the web server request into something it knows how to handle
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
