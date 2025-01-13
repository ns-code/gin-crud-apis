package main

import (
	"net/http"
	"net/http/httptest"
	"encoding/json"
    "bytes"
	// "strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ns-code/gin-crud-apis/handlers/mock"
	"github.com/ns-code/gin-crud-apis/models"
	"github.com/stretchr/testify/assert"
)

func SetupMockRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api")
	{

		v1.GET("users", mock.GetUsers)
		v1.POST("users", mock.AddUser)
		v1.PUT("users/:user_id", mock.UpdateUser)
		v1.DELETE("users/:user_id", mock.DeleteUser)
	}
	return r;
}

func TestGetUsers(t *testing.T) {
	// router := GetUsersMockRouter()
	router := SetupMockRouter(gin.Default()) // Initialize your routes

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddUserWithUserNameExistsError(t *testing.T) {
	router := gin.Default()

	SetupMockRouter(router) // Initialize routes

	// Create a test payload
	payload := models.UserDTO{
		UserName: "u123",
		FirstName:   "fn222",
		LastName: "ln222",
		Email: "u222@test.com",
		UserStatus: "T",
		Department: "",
	}

	payloadJSON, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(payloadJSON))

	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert on response
	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestAddUserWithSuccess(t *testing.T) {
	router := gin.Default()

	SetupMockRouter(router) // Initialize routes

	// Create a test payload
	payload := models.UserDTO{
		UserName: "u224",
		FirstName:   "fn222",
		LastName: "ln222",
		Email: "u222@test.com",
		UserStatus: "T",
		Department: "",
	}

	payloadJSON, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(payloadJSON))

	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert on response
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateUserWithSuccess(t *testing.T) {
	router := gin.Default()

	SetupMockRouter(router) // Initialize routes

	// Create a test payload
	payload := models.UserDTO{
		UserName: "u123",
		FirstName:   "fn222",
		LastName: "ln222",
		Email: "u222@test.com",
		UserStatus: "T",
		Department: "",
	}

	payloadJSON, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPut, "/api/users/123", bytes.NewReader(payloadJSON))

	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert on response
	assert.Equal(t, http.StatusOK, w.Code)
}


func TestDeleteUserWithSuccess(t *testing.T) {
	router := gin.Default()

	SetupMockRouter(router) // Initialize routes

	req := httptest.NewRequest(http.MethodDelete, "/api/users/123", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert on response
	assert.Equal(t, http.StatusOK, w.Code)
}
