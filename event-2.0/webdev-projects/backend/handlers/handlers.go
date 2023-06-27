package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Peter-Immanuel/nacomes-projects/event-2.0/webdev-projects/backend/models"
	"github.com/gin-gonic/gin"
)

func APIHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Backend is healthty",
	})
}

func GetRecommendation(c *gin.Context) {
	var userRequest models.UserRequest
	var movieResponse map[string]any // apiResponse from movie endpoint

	// deserialization of request body
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Function to get movies from api endpoint
	getMovies(userRequest.MovieCategory, c, &movieResponse)

	// serialize json output
	//result, err := json.Marshal(movieResponse)
	//
	//// Handle error
	//if err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"message": "Sorry we couldn't get your recommendation, try again later",
	//	})
	//}
	c.JSON(http.StatusOK, movieResponse)
}

func getMovies(search string, c *gin.Context, response *map[string]any) {
	url := os.Getenv("MOVIE_API") + "&s=" + search

	// API Call to movie endpoint. This would be replaced with a call to the ML model.
	apiResponse, err := http.Get(url)

	// Handle error
	if err != nil {
		//log error
		c.JSON(http.StatusBadGateway, gin.H{
			"error:getMovies": err.Error(),
		})
		return
	}
	// Close request util
	defer apiResponse.Body.Close()

	body, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		//log error
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Sorry we couldn't generate a recommendation now, try again later",
		})
		return
	}

	// validate json response
	if !json.Valid(body) {
		//log error
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Sorry, We couldn't generate a recommendation now, try again later",
		})
		return
	}
	json.Unmarshal(body, response)
	return
}
