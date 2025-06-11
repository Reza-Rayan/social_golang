package main

import "log"

func main() {
	cfg := config{
		address: ":5000",
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
