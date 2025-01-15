package models

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var USERDB *sql.DB
var USERDBERR bool

func ConnectUserDatabase() error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		USERDBERR = true
		return err
	}
	USERDBERR = false
	USERDB = db
	return nil
}

type UserDTO struct {
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	UserStatus string `json:"userStatus"`
	Department string `json:"department"`
}

type User struct {
	UserId     int64  `json:"userId"`
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	UserStatus string `json:"userStatus"`
	Department string `json:"department"`
}

func GetUsers(count int) ([]User, error) {

	rows, err := USERDB.Query("SELECT user_id, user_name, first_name, last_name, email, user_status, department from user LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0, count)

	for rows.Next() {
		singleUser := User{}
		err = rows.Scan(&singleUser.UserId, &singleUser.UserName, &singleUser.FirstName, &singleUser.LastName, &singleUser.Email, &singleUser.UserStatus, &singleUser.Department)

		if err != nil {
			return nil, err
		}

		users = append(users, singleUser)
		fmt.Println(">> Users count: ", len(users))
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return users, err
}

func AddUser(newUser User) (int64, error) {

	tx, err := USERDB.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	stmt, err := tx.Prepare("INSERT INTO user (user_name, first_name, last_name, email, user_status, department) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	defer stmt.Close()

	res, errres := stmt.Exec(newUser.UserName, newUser.FirstName, newUser.LastName, newUser.Email, newUser.UserStatus, newUser.Department)

	if errres != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	lastInsertedId, _ := res.LastInsertId()

	return lastInsertedId, nil
}

func UpdateUser(ourUser User, userId int) (bool, error) {

	tx, err := USERDB.Begin()
	if err != nil {
		tx.Rollback()
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE user SET user_name = ?, first_name = ?, last_name = ?, email = ?, user_status = ?, department = ? WHERE user_id = ?")

	if err != nil {
		tx.Rollback()
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ourUser.UserName, ourUser.FirstName, ourUser.LastName, ourUser.Email, ourUser.UserStatus, ourUser.Department, userId)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}

func DeleteUser(userId int) (bool, error) {

	tx, err := USERDB.Begin()

	if err != nil {
		tx.Rollback()
		return false, err
	}

	stmt, err := USERDB.Prepare("DELETE from user where user_id = ?")

	if err != nil {
		tx.Rollback()
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}

/*
func GetUserById(id string) (User, error) {

	stmt, err := DB.Prepare("SELECT id, first_name, last_name, email, ip_address from people WHERE id = ?")

	if err != nil {
		return User{}, err
	}

	person := User{}

	sqlErr := stmt.QueryRow(id).Scan(&person.Id, &person.FirstName, &person.LastName, &person.Email, &person.IpAddress)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return User{}, nil
		}
		return User{}, sqlErr
	}
	return person, nil
}





*/
