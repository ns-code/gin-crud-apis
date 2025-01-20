package models

import (
	"errors"
	"fmt"

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
	UserName   string `db:"USER_NAME" json:"userName"`
	FirstName  string `db:"FIRST_NAME" json:"firstName"`
	LastName   string `db:"LAST_NAME" json:"lastName"`
	Email      string `db:"EMAIL" json:"email"`
	UserStatus string `db:"USER_STATUS" json:"userStatus"`
	Department string `db:"DEPARTMENT" json:"department"`
}

type User struct {
	UserId     int64  `db:"USER_ID" json:"userId"`
	UserName   string `db:"USER_NAME" json:"userName"`
	FirstName  string `db:"FIRST_NAME" json:"firstName"`
	LastName   string `db:"LAST_NAME" json:"lastName"`
	Email      string `db:"EMAIL" json:"email"`
	UserStatus string `db:"USER_STATUS" json:"userStatus"`
	Department string `db:"DEPARTMENT" json:"department"`
}

func GetUsers(count int) ([]User, error) {

	users := []User{}
	err := USERDB.Select(&users, "SELECT * FROM user")
	util.CheckErr(err, "users.db SELECT error")
	return users, err
}

func AddUser(newUser User) (int64, error) {

	query := `INSERT INTO user (user_name, first_name, last_name, email, user_status, department) VALUES (:USER_NAME, :FIRST_NAME, :LAST_NAME, :EMAIL, :USER_STATUS, :DEPARTMENT)`
    sqlResult, err := USERDB.NamedExec(query, newUser)
	if err != nil {
		fmt.Println(">> sqlResult, err: ", sqlResult, err)
		return 0, err
	}
	return sqlResult.LastInsertId()
}

func DeleteUser(userId int64) (bool, error) {

	isDeleted := false
	tx := USERDB.MustBegin()
	fmt.Println(">> bef del: ", userId)
	_, err := tx.Exec("DELETE from user where user_id = $1", userId)
	if err == nil {
		isDeleted = true
		tx.Commit()
	} else {
		tx.Rollback()
	}
	fmt.Println(">> isDel: ", isDeleted)
	return isDeleted, err
}

func UpdateUser(updUser User, userId int64) (bool, error) {

 	tx := USERDB.MustBegin()
	sqlResult := tx.MustExec("UPDATE user SET user_name = $1, first_name = $2, last_name = $3, email = $4, user_status = $5, department = $6 WHERE user_id = $7", updUser.UserName, updUser.FirstName, updUser.LastName, updUser.Email, updUser.UserStatus, updUser.Department, userId)
	
	if sqlResult != nil {
		tx.Commit()
		return true, nil
	} else {
		tx.Rollback()
	}
	return false, errors.New("Update user db error")
 
/* 	query := `UPDATE user SET user_name = :USER_NAME, first_name = :FIRST_NAME, last_name = :LAST_NAME, email = :EMAIL, user_status = :USER_STATUS, department = :DEPARTMENT WHERE user_id = :USER_ID`
    sqlResult, err := USERDB.NamedExec(query, updUser)
	if err != nil {
		fmt.Println(">> sqlResult, err: ", sqlResult, err)
		return false, err
	}
	return true, nil
 */
}


