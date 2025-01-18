package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ns-code/gin-crud-apis/util"
)

var USERDB *sqlx.DB
var USERDBERR bool

func ConnectUserDatabase() error {
	db, err := sqlx.Connect("sqlite3", "users.db")
	if err != nil {
		USERDBERR = true
		return err
	}
	USERDBERR = false
	USERDB = db
	return nil
}

type UserDTO struct {
	UserName   string `db:"USER_NAME"; json:"userName"`
	FirstName  string `db:"FIRST_NAME"; json:"firstName"`
	LastName   string `db:"LAST_NAME"; json:"lastName"`
	Email      string `db:"EMAIL"; json:"email"`
	UserStatus string `db:"USER_STATUS"; json:"userStatus"`
	Department string `db:"DEPARTMENT"; json:"department"`
}

type User struct {
	UserId     int64  `db:"USER_ID"; json:"userId"`
	UserName   string `db:"USER_NAME"; json:"userName"`
	FirstName  string `db:"FIRST_NAME"; json:"firstName"`
	LastName   string `db:"LAST_NAME"; json:"lastName"`
	Email      string `db:"EMAIL"; json:"email"`
	UserStatus string `db:"USER_STATUS"; json:"userStatus"`
	Department string `db:"DEPARTMENT"; json:"department"`
}

func GetUsers(count int) ([]User, error) {

	users := []User{}
	err := USERDB.Select(&users, "SELECT * FROM user")
	util.CheckErr(err, "users.db SELECT error")
	return users, err
}


/* 
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

 *//*
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
