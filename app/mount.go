package app

import (
	"net/http"
)

// Mount function to mount path url
func Mount(mux *http.ServeMux) {
	middleware := http.NewServeMux()
	mux.Handle("/", checkPathCorrect(middleware))
	middleware.Handle("/", http.FileServer(http.Dir("./static")))
	middleware.HandleFunc("/movies", list)
	middleware.Handle("/movie/", http.StripPrefix("/movie", http.HandlerFunc(get)))
	middleware.HandleFunc("/favorite", favorite)
	middleware.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("static"))))
}

func checkPathCorrect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Todo Handle uncorrect path to index
		// log.Println(r.URL.Path)
		// if r.URL.Path != "/" {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	w.
		// 	return
		// }
		h.ServeHTTP(w, r)
	})
}
