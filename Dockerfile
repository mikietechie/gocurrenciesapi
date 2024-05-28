FROM golang:latest

COPY . .

RUN go mod download

RUN go run internal/cmd/su/main.go

ENTRYPOINT [ "go", "run", "main.go"]
