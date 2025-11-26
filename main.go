package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"io"
	"net/http"
	"os"
	"time"
)

func checkcache(c *gin.Context, rdb *redis.Client, ctx context.Context) {
	location := c.Param("location")

	cachedData, err := rdb.Get(ctx, location).Result()

	if err == nil && cachedData != "" {
		c.String(http.StatusOK, cachedData)
		return
	}

	weatherData, err := callAPI(location)
	if err != nil {
		fmt.Println("API error:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch weatherData"})
		return
	}

	rdb.Set(ctx, location, weatherData, 12*time.Hour)

	c.String(http.StatusOK, weatherData)

}

func callAPI(location string) (string, error) {

	apiKey := os.Getenv("API_KEY")

	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/next7days?unitGroup=metric&key=%s&contentType=json", location, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", fmt.Errorf("API returned status %d: %s", response.StatusCode, string(body))
	}

	return string(body), nil
}

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	ctx := context.Background()
	router := gin.Default()

	router.GET("/:location", func(c *gin.Context) {
		checkcache(c, rdb, ctx)
	})

	router.Run(":8080")

}
