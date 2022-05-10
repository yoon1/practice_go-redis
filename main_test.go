package main

import (
	"example.com/m/database"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func Test_connectRedis(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	database.ConnectRedis()
	rds, err := database.GetRedis()
	if err != nil {
		t.Fatal(err)
	}

	handle(c, rds, "UP")
	// ..
}
