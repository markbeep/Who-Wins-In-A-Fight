FROM golang:1.21-bookworm
WORKDIR /app

RUN apt-get update
RUN apt-get install -y curl libvips-dev pkg-config
RUN curl -o /bin/tailwindcss -sSL https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64
RUN chmod +x /bin/tailwindcss
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY tailwind.config.js .
COPY go.mod go.sum main.go ./
COPY internal internal
COPY components components
COPY models models
COPY migrations migrations

RUN templ generate components
RUN go get
RUN go build

COPY static static
RUN tailwindcss -i static/tw.css -o static/main.css --minify

CMD ./compare
