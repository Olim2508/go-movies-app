package controllers

import (
	"github.com/gin-gonic/gin"
	"go-movies-app/database"
	"go-movies-app/models"
	"net/http"
)

// CreateMovie godoc
// @Summary Create a new movie
// @Description Add a new movie to the database
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body models.Movie true "Movie Data"
// @Success 200 {object} models.Movie
// @Failure 400 {object} map[string]interface{}
// @Router /movies [post]
func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&movie)
	c.JSON(http.StatusOK, movie)
}

// GetMovies godoc
// @Summary Get all movies
// @Description Get a list of all movies in the database
// @Tags movies
// @Produce json
// @Success 200 {array} models.Movie
// @Router /movies [get]
func GetMovies(c *gin.Context) {
	var movies []models.Movie
	database.DB.Preload("Director").Find(&movies)
	c.JSON(http.StatusOK, movies)
}

// GetMovieByID godoc
// @Summary Get a movie by ID
// @Description Retrieve a specific movie by ID
// @Tags movies
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} models.Movie
// @Failure 404 {object} map[string]interface{}
// @Router /movies/{id} [get]
func GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := database.DB.Preload("Director").First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// CreateGenre godoc
// @Summary Create a new genre
// @Description Add a new movie to the database
// @Tags genres
// @Accept json
// @Produce json
// @Param movie body models.Genre true "Genre Data"
// @Success 200 {object} models.Genre
// @Failure 400 {object} map[string]interface{}
// @Router /genres [post]
func CreateGenre(c *gin.Context) {
	var genre models.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&genre)
	c.JSON(http.StatusOK, genre)
}

// GetGenres godoc
// @Summary Get all genres
// @Description Get a list of all genres in the database
// @Tags genres
// @Produce json
// @Success 200 {array} models.Genre
// @Router /genres [get]
func GetGenres(c *gin.Context) {
	var genres []models.Genre
	database.DB.Find(&genres)
	c.JSON(http.StatusOK, genres)
}
