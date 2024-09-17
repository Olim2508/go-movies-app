package controllers

import (
	"github.com/gin-gonic/gin"
	"go-movies-app/database"
	"go-movies-app/models"
	"net/http"
)

// CreateDirector godoc
// @Summary Create a new director
// @Description Add a new director to the database
// @Tags directors
// @Accept json
// @Produce json
// @Param director body models.Director true "Director Data"
// @Success 200 {object} models.Director
// @Failure 400 {object} map[string]interface{}
// @Router /directors [post]
func CreateDirector(c *gin.Context) {
	var director models.Director
	if err := c.ShouldBindJSON(&director); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&director)
	c.JSON(http.StatusOK, director)
}

// GetDirectors godoc
// @Summary Get all directors
// @Description Retrieve all directors in the database
// @Tags directors
// @Produce json
// @Success 200 {array} models.Director
// @Router /directors [get]
func GetDirectors(c *gin.Context) {
	var directors []models.Director
	database.DB.Find(&directors)
	c.JSON(http.StatusOK, directors)
}
