package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ShehabEl-DeenAlalkamy/residencia/internal/config"
	"github.com/ShehabEl-DeenAlalkamy/residencia/internal/handlers"
	"github.com/ShehabEl-DeenAlalkamy/residencia/internal/models"
	"github.com/ShehabEl-DeenAlalkamy/residencia/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// to become available in routes.go & middleware.go
var app config.AppConfig

// a way to make session available in main package
var session *scs.SessionManager

// main is the main application function
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// store Reservation data structure type into the session
	gob.Register(models.Reservation{})

	var app config.AppConfig

	// change this to true when in production
	app.InProduction = false

	// if it was session := scs.New() this would be a different variable other than session defined in the global scope (shadowing)
	session = scs.New()
	// how much time you want your session to live
	session.Lifetime = 24 * time.Hour
	// should the cookie persist after the browser window is closed by the end user
	session.Cookie.Persist = true
	// how strict you want to be about what site this cookie applies to
	session.Cookie.SameSite = http.SameSiteLaxMode
	// should it insist that the cookies be encrypted and the connection is from https instead
	// in production it should be set to true
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	return nil
}
