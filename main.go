package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Movie struct represents the Movie table
type Movie struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `json:"title"`
	Rating int    `json:"rating"`
}

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set up the PostgreSQL connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var errDB error
	DB, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatalf("Failed to connect to the database: %v", errDB)
	}

	// Migrate the Movie struct to create the table
	DB.AutoMigrate(&Movie{})
}

func main() {
	r := gin.Default()

	// Define routes for handling movies
	r.POST("/movies", createMovie)
	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMovieByID)

	r.Run(":8080") // Run the server on port 8080
}

// Handlers
func createMovie(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	DB.Create(&movie)
	c.JSON(201, movie)
}

func getMovies(c *gin.Context) {
	var movies []Movie
	DB.Find(&movies)
	c.JSON(200, movies)
}

func getMovieByID(c *gin.Context) {
	id := c.Param("id")
	var movie Movie
	if err := DB.First(&movie, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(200, movie)
}
