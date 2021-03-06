package main

import (
	"example.com/m/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	database.ConnectRedis()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("no port variables")
	}

	redis, err := database.GetRedis()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	setUpRouter(r, redis)
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setUpRouter(r *gin.Engine, rds *redis.Client) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/up", func(c *gin.Context) {
		handleCount(c, rds, "UP")

		c.JSON(200, gin.H{
			"item": getCount(c, rds),
		})
	})

	r.GET("/up/:count", func(c *gin.Context) {
		paramCount := c.Param("count")
		countVal, _ := strconv.ParseInt(paramCount, 10, 64)
		count := int(countVal)

		for i := 0; i < count; i++ {
			handleCount(c, rds, "UP")
			time.Sleep(5 * time.Second)
		}

		c.JSON(200, gin.H{
			"item": getCount(c, rds),
		})
	})

	r.GET("/down", func(c *gin.Context) {
		handleCount(c, rds, "DOWN")
		c.JSON(200, gin.H{
			"item": getCount(c, rds),
		})

	})

	r.GET("/get", func(c *gin.Context) {
		handleCount(c, rds, "GET")
		c.JSON(200, gin.H{
			"item": getCount(c, rds),
		})
	})
}

func getCount(c *gin.Context, rds *redis.Client) int64 {
	var cntVal int64
	cnt := database.GetItem(c, rds, "cnt")
	if cnt == "" {
		return 0
	}

	cntVal, _ = strconv.ParseInt(cnt, 10, 64)
	return cntVal
}

func handleCount(c *gin.Context, rds *redis.Client, opt string) {
	cntVal := getCount(c, rds)
	switch opt {
	case "UP":
		rds.Set(c, "cnt", cntVal+1, 0)
	case "DOWN":
		rds.Set(c, "cnt", cntVal-1, 0)
	default:
		break
	}
}
