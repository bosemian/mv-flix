package app

import (
	"net/http"
)

// Mount function to mount path url
func Mount(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/movies", list)
	mux.Handle("/movie/", http.StripPrefix("/movie", http.HandlerFunc(get)))
	mux.HandleFunc("/favorite", favorite)
	mux.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("static"))))
}
