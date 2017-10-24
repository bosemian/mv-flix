package view

import (
	"net/http"

	"github.com/bosemian/mv-flix/pkg/model"
)

type MovieData struct {
	Movies []*model.Movie
}

// Index render index view
func Index(w http.ResponseWriter, data *MovieData) {
	render(tpIndex, w, data)
}
