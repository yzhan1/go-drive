package db

import (
	"fmt"
	db "github.com/yzhan1/go-drive/db/mysql"
)

func SignUp(username string, password string) bool {
	statement, err := db.GetConn().Prepare("insert ignore into tbl_user (`user_name`, `user_pwd`) values (?, ?)")
	if err != nil {
		fmt.Println("Failed to insert user: " + err.Error())
		return false
	}
	defer statement.Close()

	ret, err := statement.Exec(username, password)
	if err != nil {
		fmt.Println("Failed to insert user: " + err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return true
	}
	return false
}

func SignIn(username string, password string) bool {
	statement, err := db.GetConn().Prepare("select * from tbl_user where user_name = ? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	row, err := statement.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if row == nil {
		fmt.Println("Username " + username + " not found")
		return false
	}

	parsedRow := db.ParseRows(row)
	if len(parsedRow) > 0 && string(parsedRow[0]["user_pwd"].([]byte)) == password {
		return true
	}
	return false
}

func UpdateToken(username string, token string) bool {
	statement, err := db.GetConn().Prepare("replace into tbl_user_token (`user_name`, `user_token`) values (?, ?) ")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer statement.Close()

	_, err = statement.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

func GetUserInfo(username string) (User, error) {
	user := User{}

	statement, err := db.GetConn().Prepare("select user_name, signup_at from tbl_user where user_name = ? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer statement.Close()

	err = statement.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	return user, nil
}
