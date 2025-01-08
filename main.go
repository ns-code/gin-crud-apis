package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ns-code/gin-crud-apis/docs"
	"github.com/ns-code/gin-crud-apis/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Create a CORS config with allowed origins, methods, headers etc.

	corsConfig := cors.Config{

		AllowOrigins: []string{"http://localhost:4200"},

		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},

		// AllowHeaders:     []string{"Content-Type", "Authorization"},

		// AllowCredentials: true, // Allow sending cookies with requests

	}

	// @Summary get users
	// @Description get string by ID
	// @Produce  json


	docs.SwaggerInfo.Title = "gin-crud-apis"
	docs.SwaggerInfo.Description = "gin crud apis"
	docs.SwaggerInfo.Version = ""
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}


	err := models.ConnectUserDatabase()
	checkErr(err)

	r := gin.Default()
	r.Use(cors.New(corsConfig))

	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// url := ginSwagger.URL("http://localhost:8080/docs/swagger.json")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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



// @Tags         users
// @Produce      json
// @Success 200 {array} models.User
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
