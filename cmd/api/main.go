package main

import (
	"github.com/Reza-Rayan/social_golang/internal/env"
	"log"
)

func main() {
	cfg := config{
		address: env.GetString("ADDRESS", ":5000"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
