package app

import (
	"net/http"

	"github.com/bosemian/mv-flix/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	view.Index(w, nil)
}
