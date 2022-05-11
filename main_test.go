package main

import (
	"example.com/m/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	beforeCount := getCount(c, rds)

	handleCount(c, rds, "UP")

	afterCount := getCount(c, rds)

	assert.Equal(t, beforeCount+1, afterCount)
}
