package mock

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ns-code/gin-crud-apis/models"
)

var usersDB = []models.User{
	{UserId: 123, UserName: "u123", FirstName: "fname1", LastName: "lname1", Email: "email1@test.com", UserStatus: "I", Department: ""},
	{UserId: 456, UserName: "u456", FirstName: "fname2", LastName: "lname2", Email: "email2@test.com", UserStatus: "A", Department: ""},
}

// @Description get all users
// @Tags         users
// @Produce      json
// @Success 200 {array} models.User
// @Failure   400  "Bad Request"
// @Router /api/users [get]
func GetUsers(c *gin.Context) {

	users := usersDB

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
func AddUser(c *gin.Context) {

	var user models.User

	log.Println(c.Request.Body)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if userName already exists
	errmsg := ""
	for _, usr := range usersDB {
		if usr.UserName == user.UserName {
			errmsg = "User Name: " + user.UserName + " exists"
			break
		}
	}

	lenBef := len(usersDB)
	usersDB := append(usersDB, user)
	lenAft := len(usersDB)

	if errmsg == "" && lenAft == lenBef+1 {
		c.JSON(http.StatusCreated, user)
	} else if len(errmsg) > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": errmsg})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
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
func UpdateUser(c *gin.Context) {

	var payload models.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	fmt.Printf("Updating id %d", userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
	}

	// update usersDB
	errmsg := ""
	for _, usr := range usersDB {
		if usr.UserName != payload.UserName {
			errmsg = "User Name " + payload.UserName + " exists."
			break
		} else if usr.UserId == userId {
			usr = payload
			usr.UserId = userId
			break
		}
	}

	if errmsg == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Update User Success"})
	} else {
		if strings.Contains(errmsg, "unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": errmsg})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Update User error"})
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
func DeleteUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
	}

	newItems := []models.User{}
	for _, item := range usersDB {
		if item.UserId != userId {
			newItems = append(newItems, item)
		}
	}

	if len(newItems) == len(usersDB)-1 {
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

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
