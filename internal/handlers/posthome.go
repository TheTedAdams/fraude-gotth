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
	message := r.PostFormValue("message")

	// TODO: Send message to ollama

	err := templates.MessageListElement(message).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
