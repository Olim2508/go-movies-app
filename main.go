package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-movies-app/database"
	_ "go-movies-app/docs"
	"go-movies-app/routes"
)

// @title Movie API
// @version 1.0
// @description A simple Movie and Director API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	// Connect to the database
	database.Connect()

	// Set up the Gin router
	r := gin.Default()

	// Define routes
	routes.SetupRoutes(r)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server on port 8080
	r.Run(":8080")
}
