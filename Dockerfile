# Date Created		    1 May 2024
# Author				Mike Z
# Email				    mzinyoni7@outlook.com
# Website				https://mikeio.web.app
# Status				Looking for a job!
# Description			A Fintech Data Service
# Inspired by			https://freecurrencyapi.com

FROM golang:latest

COPY . .

RUN go mod download

RUN go run internal/cmd/su/main.go

ENTRYPOINT [ "go", "run", "main.go"]
