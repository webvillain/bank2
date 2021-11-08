package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *sql.DB

func ConnectDb() {
	// err :=
	os.Remove("/home/vijay/Desktop/workspace/go/src/github.com/webvillain/bank2/db/test.db")
	// this is the absolute path for removing any existing file in my app
	// so everytime i run this program , the previos file/database will automatically deleted
	// if err != nil {
	// 	log.Fatal("Cann't Remove Existing/Previous Database")
	// }
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection To Database Is Successful .")
	defer db.Close()
	DB = db

	mySCHEMA := `
	CREATE TABLE IF NOT EXISTS users(Id INTEGER PRIMARY KEY AUTOINCREMENT , Name TEXT NOT NULL , Email TEXT NOT NULL);
	`
	stmt, err := db.Prepare(mySCHEMA)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
	fmt.Println("Table Created Successfully.")

	// i think upto this part of this application will work find
	// but lets be sure by running this program once more
	var Name string
	var Email string
	rows, err := db.Query("SELECT Name ,Email FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&Name, &Email)
	}
	fmt.Println(Name, Email)

}

// get all users info
func GetAllUsers(db *sql.DB) {

}

// create new user into database
func CreateNewUser(db *sql.DB) {

	stmt, err := db.Prepare("INSERT INTO users VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(1, "Vikash", "Vikash@test.com")
	if err != nil {
		log.Fatal(err)
	}
	res.RowsAffected()
	fmt.Println("Data Is Inserted Successfully into table")
}

func GetSingleUser(db *sql.DB, Id *User) {

}
