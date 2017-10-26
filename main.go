package main

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/bosemian/mv-flix/app"
)

const (
	port = ":3000"
)

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, cors.Default().Handler(mux))
}
