package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/shaheerkj/movies-crud-api/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (app *application) getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(app.movies)
}
func (app *application) getMovie(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	w.Header().Set("Content-Type", "application/json")
	for _, item := range app.movies {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//json.NewEncoder(w).Encode()
}
func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(1000))
	app.movies = append(app.movies, movie)

	json.NewEncoder(w).Encode(app.movies)

}
func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {

}
func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	for index, item := range app.movies {
		if item.ID == id {
			app.movies = append(app.movies[:index], app.movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(app.movies)

}
