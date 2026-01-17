package main

import (
	"flag"
	"log/slog"
	"net/http"

	"github.com/shaheerkj/movies-crud-api/models"
	_ "github.com/shaheerkj/movies-crud-api/models"
)

type application struct {
	logger *slog.Logger
	movies []models.Movie
}

func main() {

	addr := flag.String("addr", ":4000", "Http Network address")

	flag.Parse()

	app := application{}

	app.movies = append(app.movies, models.Movie{ID: "1", Isbn: "123", Title: "Hello World", Director: &models.Director{Firstname: "Aqib", Lastname: "Javed"}})
	app.movies = append(app.movies, models.Movie{ID: "2", Isbn: "456", Title: "Dead World", Director: &models.Director{Firstname: "Saqib", Lastname: "Nisar"}})

	http.ListenAndServe(*addr, app.routes())
}
