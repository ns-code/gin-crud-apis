package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

// @Description Add a new user
// @Tags        users
// @Accept       json
// @Param		users body models.UserDTO true "Create User"
// @Produce      json
// @Success 201  {object} models.User "User created"
// @Failure 409  "User Name exists"
// @Failure 400  "Bad Request"
// @Router /api/users [post]
func AddUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cannot map json to User"))
	} else {
		fmt.Println(">> new user: ", user.UserName, user.FirstName, user.LastName)
		lastInsertedId, err := models.AddUser(user)

		if lastInsertedId > 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			user.UserId = lastInsertedId
			usersBytes, err := json.Marshal(user)
			util.CheckErr(err, "GetUsers usersBytes error")
			w.Write(usersBytes)
		} else {
			if strings.Contains(strings.ToLower(err.Error()), "unique constraint") {
				w.WriteHeader(http.StatusConflict)
			} else {
				w.WriteHeader(http.StatusBadRequest)
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
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	toks := strings.Split(r.URL.Path, "/")
	userIdStr := toks[len(toks)-1]

	fmt.Println(">> queryValues: ", r.URL.Path, userIdStr)
	// Access individual query parameters
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		fmt.Println(">> parse err: ", err.Error())
		http.Error(w, "Error parsing query parameters", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cannot map json to User"))
	} else {
		isUpdateSuccess, err := models.UpdateUser(user, userId)

		if isUpdateSuccess {
			w.WriteHeader(http.StatusNoContent)
		} else {
			if strings.Contains(strings.ToLower(err.Error()), "unique constraint") {
				w.WriteHeader(http.StatusConflict)
			} else {
				w.WriteHeader(http.StatusBadRequest)
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
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// userIdStr := r.PathValue("userId")
	toks := strings.Split(r.URL.Path, "/")
	userIdStr := toks[len(toks)-1]

	fmt.Println(">> queryValues: ", r.URL.Path, userIdStr)
	// Access individual query parameters
	userId, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		fmt.Println(">> parse err: ", err.Error())
		http.Error(w, "Error parsing query parameters", http.StatusBadRequest)
		return
	}
	success, _ := models.DeleteUser(userId)

	if success {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

