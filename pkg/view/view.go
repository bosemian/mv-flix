package view

import (
	"net/http"
)

// Index render index view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}
