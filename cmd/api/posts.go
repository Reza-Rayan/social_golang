package main

import "net/http"

func (app *application) getAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show All Posts"))
}
