/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package main

import (
	"github.com/mikietechie/gocurrenciesapi/internal/initialize"
	"github.com/mikietechie/gocurrenciesapi/internal/server"
)

// @title                       Go Currencies API
// @version                     0.1
// @description                 Go currencies API
// @contact.name                Mike Zinyoni
// @contact.email               mzinyoni7
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
