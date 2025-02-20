FROM golang:1.24-alpine
WORKDIR /app

RUN apk add --no-cache make gcc g++ curl

RUN go install github.com/a-h/templ/cmd/templ@latest && \
  go install github.com/air-verse/air@latest

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v4.0.7/tailwindcss-linux-x64-musl && \
  chmod +x tailwindcss-linux-x64-musl && \
  mv tailwindcss-linux-x64-musl /usr/local/bin/tailwindcss

COPY dev/scripts/generate.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/generate.sh

COPY dev/scripts/entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]
