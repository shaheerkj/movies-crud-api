package main

import "net/http"

func (m *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", m.home)
	mux.HandleFunc("GET /movies", m.getMovies)
	mux.HandleFunc("GET /movies/{id}", m.getMovie)
	mux.HandleFunc("POST /movies", m.createMovie)
	mux.HandleFunc("PUT /movies/{id}", m.updateMovie)
	mux.HandleFunc("DELETE /movies/{id}", m.deleteMovie)

	return mux
}
