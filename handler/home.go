package handler

import (
	"dreampixai/view/home"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) {
	home.Index().Render(r.Context(), w)
}
