package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ns-code/gin-crud-apis/models"
	"github.com/ns-code/gin-crud-apis/util"
)

// @Description get all users
// @Tags         users
// @Produce      json
// @Success 200 {array} models.User
// @Failure   400  "Bad Request"
// @Router /api/users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {

		users, err := models.GetUsers(10)
		util.CheckErr(err, "GetUsers users error")

		usersBytes, err := json.Marshal(users)
		util.CheckErr(err, "GetUsers usersBytes error")

		if users == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No Records Found"))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(usersBytes)
		}	
}

/* 
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

	if !ServerError(c) {
		var user models.User

		log.Println(c.Request.Body)

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		lastInsertedId, err := models.AddUser(user)

		if lastInsertedId > 0 {
			user.UserId = lastInsertedId
			c.JSON(http.StatusCreated, user)
		} else {
			if strings.Contains(strings.ToLower(err.Error()), "unique constraint") {
				c.JSON(http.StatusConflict, gin.H{"error": err})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			}
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
func UpdateUser(c *gin.Context) {
	if !ServerError(c) {

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
}

// @Description delete a user by user_id
// @Tags users
// @Param        userId     path    int     true        "User ID"
// @Success 204  "No Content"
// @Failure   409  "User Name exists"
// @Failure   400  "Bad Request"
// @Router /api/users/{userId} [delete]
func DeleteUser(c *gin.Context) {
	if !ServerError(c) {

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
}

func ServerError(c *gin.Context) bool {
	if models.USERDBERR {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return true
	}
	return false
}

 *//* func getPersonById(c *gin.Context) {

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


