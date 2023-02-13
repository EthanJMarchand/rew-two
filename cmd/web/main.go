package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ethanjmarchand/rewtwo/pkg/config"
	"github.com/ethanjmarchand/rewtwo/pkg/handlers"
	"github.com/ethanjmarchand/rewtwo/pkg/render"
)

// This is the port number that the server is launching on.
const portNumber = ":80"

var app config.AppConfig
var session *scs.SessionManager

// main - This is the entry point into the application
func main() {

	app.InProduction = false

	session = scs.New()
	// Lifetime sets the duration of the session cookie
	session.Lifetime = 24 * time.Hour
	// Persist makes it so the session cookie survives a browser close, and reopen
	session.Cookie.Persist = true
	// Samesite lax mode allows them to come from a different area, and still be logged in.
	session.Cookie.SameSite = http.SameSiteLaxMode
	// Secure sets the token or cookie to be encrypted.
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port%s", portNumber)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
