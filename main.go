package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sixtengedda/weather-api/config"
	"io"
	"log"
	"net/http"
)

func callAPI(c *gin.Context) {
	response, err := http.Get(config.API_URL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	c.Data(http.StatusOK, "application/json", body)

}

func main() {
	router := gin.Default()
	router.GET("/weather", callAPI)

	router.Run("localhost:8080")
}
