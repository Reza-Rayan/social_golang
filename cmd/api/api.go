package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	address string
}

// EndPoints Handler
func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	//Middlewares
	r.Use(middleware.Recoverer) // Recover from a panic
	r.Use(middleware.Logger)    // add logs to terminal for calling routes
	r.Use(middleware.RequestID) //Injects a request ID into the context of each request
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(time.Second * 30)) // Timed out for each request duration
	r.Use(middleware.Throttle(100))             // limitation request counts per minute

	//EndPoints
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		//	Post Routes
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", app.getAllPostsHandler)
		})
	})
	return r
}

// Run Server
func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("listening on %s", app.config.address)

	return srv.ListenAndServe()
}

//End here
