package app

import (
	"encoding/json"
	"net/http"

	"github.com/bosemian/mv-flix/api"
)

func list(w http.ResponseWriter, r *http.Request) {
	var rest api.MovieList
	movies, err := rest.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(json))
}

func get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	var rest api.MovieData
	movie, err := rest.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(json))
}
