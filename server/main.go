package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//User is define a user
type User struct {
	Username string
	Email    string
	Password string
}

func main() {
	db, err := sql.Open("sqlite3", "database/test.db")
	checkError(err)

	db.Exec("create table if not exists users (username text, email text, password text)")

	addUser(db, "name1", "name1@gmail.com", "name1sifre")
	fmt.Println("done!")
	fmt.Println(getUser(db, "name1@gmail.com"))

}

func addUser(db *sql.DB, username string, email string, pass string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into users(username,email,password) values(?,?,?)")
	_, err := stmt.Exec(username, email, pass)
	checkError(err)
	tx.Commit()
}

func getUser(db *sql.DB, mail string) User {
	rows, err := db.Query("select * from users")
	checkError(err)
	for rows.Next() {
		var tempUser User
		err = rows.Scan(&tempUser.Username, &tempUser.Email, &tempUser.Password)
		checkError(err)
		if tempUser.Email == mail {
			return tempUser
		}

	}
	return User{}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
