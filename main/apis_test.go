package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func GetUsersMockRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/users", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func TestGetUsers(t *testing.T) {
	router := GetUsersMockRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}

