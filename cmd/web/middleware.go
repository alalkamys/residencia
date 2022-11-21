package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// it uses cookies to make sure that the token it generates is available on a per page basis
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		// to be applied to the entire site
		Path: "/",
		// false because we are not running https right now
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads adn saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	// provides middleware which automatically loads and saves session data for the current request, and communcates teh ression token to and from
	// the client in a cookie
	return session.LoadAndSave(next)
}
