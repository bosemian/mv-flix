package app

import (
	"net/http"
)

// Mount function to mount path url
func Mount(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/movies", index)
	// mux.Handle("/movie/", http.StripPrefix("/movie", http.HandlerFunc(movie)))
	mux.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("static"))))
}
