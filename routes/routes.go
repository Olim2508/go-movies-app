package routes

import (
	"github.com/gin-gonic/gin"
	"go-movies-app/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Movie routes
	router.POST("/movies", controllers.CreateMovie)
	router.GET("/movies", controllers.GetMovies)
	router.GET("/movies/:id", controllers.GetMovieByID)

	// Genre routes
	router.POST("/genres", controllers.CreateGenre)
	router.GET("/genres", controllers.GetGenres)

	// Director routes
	router.POST("/directors", controllers.CreateDirector)
	router.GET("/directors", controllers.GetDirectors)
}
