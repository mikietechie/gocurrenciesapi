/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package server

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/docs"
	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"github.com/mikietechie/gocurrenciesapi/internal/controllers"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func RunServer() {
	docs.SwaggerInfo.Title = "Gin Swagger"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.ReadDoc()

	server := gin.Default()
	// server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Use(gin.Logger())
	server.Use(cors.Default())
	server.Use(gin.Recovery())

	// Routing
	api_router := server.Group("/api/v1")
	controllers.ServiceRouter(*api_router.Group("/service"))
	controllers.AuthRouter(*api_router.Group("/auth"))
	controllers.UsersRouter(*api_router.Group("/users"))
	controllers.ClientsRouter(*api_router.Group("/clients"))
	controllers.IndexRouter(*api_router.Group("/"))
	if config.DEV {
		server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// Run Server
	ADDRESS := "http://" + config.SERVER_ADDRESS
	log.Println(ADDRESS)
	log.Println(ADDRESS + "/docs/index.html")
	http.ListenAndServe(config.SERVER_ADDRESS, server)
}
