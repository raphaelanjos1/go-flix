package main

import (
	"log"
	"net/http"

	"github.com/raphaelanjos1/go-flix/internal/routes"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	router := routes.NewRouter()

	return router
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: app.mount(),
	}
	log.Printf("Starting server on %s", app.config.addr)
	return srv.ListenAndServe()

}
