FROM golang:latest

COPY . .

RUN go mod download

# RUN go build -o bin/dockeredgo

# RUN chmod +x ./bin/dockeredgo

# ENTRYPOINT [ "./bin/dockeredgo" ]
ENTRYPOINT [ "go", "run", "src/cmd/main/main.go"]
