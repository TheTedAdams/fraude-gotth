package handlers

import (
	"context"
	"goth/internal/templates"
	"log"
	"net/http"

	"github.com/ollama/ollama/api"
)

type PostHomeHandler struct{}

func NewPostHomeHandler() *PostHomeHandler {
	return &PostHomeHandler{}
}

func (h *PostHomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := r.PostFormValue("message")

	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	messages := []api.Message{
		{
			Role:    "user",
			Content: message,
		},
	}

	ctx := context.Background()
	req := &api.ChatRequest{
		Model:    "deepseek-r1",
		Messages: messages,
		Stream:   new(bool),
	}

	respFunc := func(resp api.ChatResponse) error {
		err = templates.MessageListElement(resp.Message.Content).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
		return nil
	}

	err = client.Chat(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
