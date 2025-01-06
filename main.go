package main

import (
	"github.com/ns-code/gin-crud-apis/models"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors" 
)



func main() {

	    // Create a CORS config with allowed origins, methods, headers etc.

		corsConfig := cors.Config{

			AllowOrigins:     []string{"http://localhost:4200"},
	
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	
			// AllowHeaders:     []string{"Content-Type", "Authorization"},
	
			// AllowCredentials: true, // Allow sending cookies with requests
	
		}
	

	err := models.ConnectUserDatabase()
	checkErr(err)

	r := gin.Default()
	r.Use(cors.New(corsConfig))

	// API v1
	v1 := r.Group("/api")
	{
		v1.GET("users", getUsers)
		// v1.GET("person", getPersons)

/* 		v1.GET("person/:id", getPersonById)
		v1.POST("person", addPerson)
		v1.PUT("person/:id", updatePerson)
		v1.DELETE("person/:id", deletePerson)
		v1.OPTIONS("person", options) */
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}

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

func addPerson(c *gin.Context) {

	var json models.Person

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := models.AddPerson(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func updatePerson(c *gin.Context) {

	var json models.Person

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personId, err := strconv.Atoi(c.Param("id"))

	fmt.Printf("Updating id %d", personId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := models.UpdatePerson(json, personId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func deletePerson(c *gin.Context) {

	personId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := models.DeletePerson(personId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
} */

func options(c *gin.Context) {

	ourOptions := "HTTP/1.1 200 OK\n" +
		"Allow: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Origin: http://locahost:8080\n" +
		"Access-Control-Allow-Methods: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Headers: Content-Type\n"

	c.String(200, ourOptions)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

