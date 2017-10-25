package app

import (
	"net/http"

	"github.com/bosemian/mv-flix/pkg/model"
	"github.com/bosemian/mv-flix/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	movies := model.ListMovies()
	view.Index(w, &view.MovieData{
		Movies: movies,
	})
}
