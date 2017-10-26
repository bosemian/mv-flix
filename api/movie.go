package api

import (
	"encoding/json"
	"net/http"

	"github.com/bosemian/mv-flix/model"
)

type Movier interface {
	List() (*MovieList, error)
}

type MovieList struct {
	Movies []*model.Movie `json:"movies"`
}

func (m *MovieList) List() (*MovieList, error) {
	movies := model.ListMovies()
	m.Movies = movies
	return m, nil
}

func movie(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	movie := model.GetMovie(id)
	mJson, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(mJson)
}
