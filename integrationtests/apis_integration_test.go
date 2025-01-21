package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"testing"

	"github.com/ns-code/gin-crud-apis/models"

	"github.com/stretchr/testify/assert"
)

var testApi = "http://localhost:8080/api/users"
var testUserIdStr string
var testUserName = "u000"

func TestCreateUserApi(t *testing.T) {
	statusCode, err := callCreateUserApi()
	if err != nil {
		assert.Equal(t, http.StatusCreated, statusCode, "Should return statuscode 201")
	}
}

func TestCreateUserApiWithUserNameExistsError(t *testing.T) {
	statusCode, err := callCreateUserApi()
	if err != nil {
		assert.Equal(t, http.StatusConflict, statusCode, "Should return statuscode 409")
	}
}

func TestGetUsersApi(t *testing.T) {		
	users, statusCode := callGetUsersApi()
	assert.Equal(t, http.StatusOK, statusCode, "Should return statuscode 200")
	assert.Greater(t, len(users), 0, "Should return one or more users")
}

func TestUpdateUserApi(t *testing.T) {
	statusCode, err := callUpdateUserApi()
	if err != nil {
		fmt.Println(">> Update user api error: ", err)
		return
	}
	assert.Equal(t, http.StatusNoContent, statusCode, "Should return statuscode 204")
}

func TestDeleteUserApi(t *testing.T) {
	statusCode, err := callDeleteUserApi()
	if err != nil {
		fmt.Println(">> Delete user api error: ", err)
		return
	}
	assert.Equal(t, http.StatusNoContent, statusCode, "Should return statuscode 204")
}

func callGetUsersApi() ([]models.User, int) {
	resp, err := http.Get(testApi)
	if err != nil {
		fmt.Println("getUsers api http Error:", err)
		return nil, 0
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("getUsers api response read Error:", err)
		return nil, 0
	}
	var users []models.User
	json.Unmarshal(body, &users)
	return users, statusCode
}

func callCreateUserApi() (int, error) {
	user := "{\"userName\":\"" + testUserName + "\",\"firstName\":\"fname1\",\"lastName\":\"lname1\",\"email\":\"email1@test.com\",\"userStatus\":\"I\",\"department\":\"\"}"
	resp, err := http.Post(testApi, "application/json", bytes.NewBufferString(user))
	if err != nil {
		fmt.Println("createUser api http Error:", err)
		return 0, err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode == http.StatusCreated {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(">> createUser api response read Error:", err)
			return 0, err
		}
		
		var newUser models.User
		err = json.Unmarshal(body, &newUser)
		if err != nil {
			fmt.Println(">> createUser api unmarshal error:", err)
			return 0, err
		}
		testUserIdStr = strconv.FormatInt(newUser.UserId, 10)
	}
	return statusCode,  nil
}

func callUpdateUserApi() (int, error) {
	userJson := "{\"userName\":\"" + testUserName + "\",\"firstName\":\"fname1\",\"lastName\":\"lname1\",\"email\":\"email1@test.com\",\"userStatus\":\"I\",\"department\":\"\"}"
	req, err := http.NewRequest("PUT", testApi + "/" + testUserIdStr, bytes.NewBufferString(userJson))
	if err != nil {
		fmt.Println(">> Error creating PUT request:", err)
		return 500, err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := &http.Client{}

	// Send the PUT request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(">> Error sending PUT request:", err)
		return 500, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func callDeleteUserApi() (int, error) {
	req, err := http.NewRequest("DELETE", testApi + "/" + testUserIdStr, nil)
	if err != nil {
		fmt.Println(">> Error creating DELETE user request: ", err)
		return 500, err
	}

	// Create an HTTP client
	client := &http.Client{}

	// Send the PUT request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(">> Error sending DELETE user request: ", err)
		return 500, err
	}

	return resp.StatusCode, nil
}
