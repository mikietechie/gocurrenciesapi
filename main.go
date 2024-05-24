package main

import (
	"github.com/mikietechie/gocurrenciesapi/internal/initialize"
	"github.com/mikietechie/gocurrenciesapi/internal/server"
)

// @securityDefinitions.apikey	Bearer
// @in                         	header
// @name                       	Authorization
// / @description              	Type "Bearer" followed by a space and JWT token.
// @securityDefinitions.apikey 	ApiKeyAuth
// @in                         	query
// @name                       	apikey
// / @description               Client API Key
func main() {
	initialize.Init()
	defer initialize.Tear()
	server.RunServer()
}
