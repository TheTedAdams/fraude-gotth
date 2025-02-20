FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest

ENTRYPOINT ["templ", "generate"]

# CMD "templ generate --watch --proxy=\"http://localhost:3000\" --cmd=\"go run cmd/main.go\""
