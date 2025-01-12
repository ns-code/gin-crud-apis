package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ns-code/gin-crud-apis/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ns-code/gin-crud-apis/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// @Summary get users
	// @Description get string by ID
	docs.SwaggerInfo.Title = "gin-crud-apis"
	docs.SwaggerInfo.Description = "gin crud apis"
	docs.SwaggerInfo.Version = ""
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}


	err := models.ConnectUserDatabase()
	checkErr(err)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api")
	{

		v1.GET("users", getUsers)
		v1.POST("users", addUser)
		v1.PUT("users/:user_id", updateUser)
		v1.DELETE("users/:user_id", deleteUser)
	}

	r.Run()
}

// @Description get all users
// @Tags         users
// @Produce      json
// @Success 200 {array} models.User
// @Failure   400  "Bad Request"  
// @Router /api/users [get]
func getUsers(c *gin.Context) {

	users, err := models.GetUsers(10)

	checkErr(err)

	if users == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// @Description Add a new user
// @Tags        users
// @Accept       json
// @Param		users body models.UserDTO true "Create User"
// @Produce      json
// @Success 201  {object} models.User "User created"
// @Failure 409  "User Name exists"  
// @Failure 400  "Bad Request"  
// @Router /api/users [post]
func addUser(c *gin.Context) {

	var json models.User

	log.Println(c.Request.Body)

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := models.AddUser(json)

	if success {
		c.JSON(http.StatusCreated, gin.H{"message": "Add User Success"})
	} else {
		if strings.Contains(strings.ToLower(err.Error()), "unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": err})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
	}
}

// @Description update a user
// @Tags users
// @Accept       json
// @Param		 userId path string true "update user by id"
// @Param		 user body models.UserDTO true  "Update user"
// @Success 200  {object} models.User "User updated"
// @Failure   409  "User Name exists"  
// @Failure   400  "Bad Request"  
// @Router /api/users/{userId} [put]
func updateUser(c *gin.Context) {

	var json models.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := strconv.Atoi(c.Param("user_id"))

	fmt.Printf("Updating id %d", userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
	}

	success, err := models.UpdateUser(json, userId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Update User Success"})
	} else {
		if strings.Contains(strings.ToLower(err.Error()), "unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": err})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
	}
}

// @Description delete a user by user_id
// @Tags users
// @Param        userId     path    int     true        "User ID"
// @Success 204  "No Content"
// @Failure   409  "User Name exists"  
// @Failure   400  "Bad Request"  
// @Router /api/users/{userId} [delete]
func deleteUser(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
	}

	success, err := models.DeleteUser(userId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Delete User Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

/* func getPersonById(c *gin.Context) {

	// grab the Id of the record want to retrieve
	id := c.Param("id")

	person, err := models.GetPersonById(id)

	checkErr(err)
	// if the name is blank we can assume nothing is found
	if person.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": person})
	}
}

*/

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
