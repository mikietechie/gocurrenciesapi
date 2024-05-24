package server

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/docs"
	"github.com/mikietechie/gocurrenciesapi/internal/api"
	"github.com/mikietechie/gocurrenciesapi/internal/config"

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
	api.AuthRouter(*api_router.Group("/auth"))
	api.CurrenciesRouter(*api_router.Group("/currencies"))
	api.UsersRouter(*api_router.Group("/users"))
	api.ClientsRouter(*api_router.Group("/clients"))
	api.IndexRouter(*api_router.Group("/"))
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Run Server
	ADDRESS := "http://" + config.SERVER_ADDRESS
	log.Println(ADDRESS)
	log.Println(ADDRESS + "/docs/index.html")
	http.ListenAndServe(config.SERVER_ADDRESS, server)
}
