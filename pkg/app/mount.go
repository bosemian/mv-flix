package app

import "net/http"

// Mount function to mount path url
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
}
