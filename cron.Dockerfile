FROM golang:latest

COPY . .

RUN go mod download

ENTRYPOINT [ "go", "run", "internal/cmd/cron/main.go"]
