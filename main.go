package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sixtengedda/weather-api/config"
	"io"
	"log"
	"net/http"
)

func checkcache(c *gin.Context, rdb *redis.Client, ctx context.Context) {
	key := c.Query("location")
	value, _ := rdb.Get(ctx, key).Result()

	if value != "" {
		c.JSON(200, gin.H{
			"location": key,
			"weather":  value,
			"cached":   true,
		})
		return
	}
	c.JSON(404, gin.H{
		"error": "No weather found for this location",
	})

}

func callAPI(c *gin.Context) {
	response, err := http.Get(config.API_URL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		log.Fatalf("Response test failed with status code: %d and\nbody: %s\n", response.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	c.Data(http.StatusOK, "application/json", body)

}

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	router := gin.Default()
	router.GET("/weather", func(c *gin.Context) {
		checkcache(c, rdb, ctx)
	})

	router.Run(":8080")
}
