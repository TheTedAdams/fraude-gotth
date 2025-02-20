package handlers

import (
	"net/http"

	"goth/internal/templates"
)

type HomePageData struct {
	Title string
	Items []string
}

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Index([]string{"Hello! How can I help you today?"})
	err := templates.Layout(c, "Home").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
