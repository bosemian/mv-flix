package app

import (
	"encoding/json"
	"net/http"

	"github.com/bosemian/mv-flix/api"
)

func index(w http.ResponseWriter, r *http.Request) {
	var rest api.MovieList
	movies, err := rest.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(js))
}
