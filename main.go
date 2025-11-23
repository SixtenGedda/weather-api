package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var weather = "sunny"

// getAlbums responds with the list of all albums as JSON.
func getWeather(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, weather)
}

func main() {
	router := gin.Default()
	router.GET("/weather", getWeather)

	router.Run("localhost:8080")
}
