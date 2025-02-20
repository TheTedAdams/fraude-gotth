package handlers

import (
	"goth/internal/templates"
	"net/http"
)

type PostHomeHandler struct{}

func NewPostHomeHandler() *PostHomeHandler {
	return &PostHomeHandler{}
}

func (h *PostHomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("item")

	err := templates.ItemListElement(name).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
