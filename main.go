package main

import (
	dbmanager "SQLITE/db_manager"
	"SQLITE/users"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	new_user := users.User{}
	new_user.Id = uuid.New().String()

	fmt.Print("Type your username: ")
	fmt.Scanln(&new_user.Username)

	//Our DB Manager implementation, now is SQLite
	dbManager := dbmanager.SQLite{}

	imp := dbmanager.DBImplementation{ //We pass our objecto to this struct for accessing the interface methods
		Conn: dbManager,
	}

	myconn := imp.Conn.GetDBConnection() //We need to create a connection for creating our table for storing our users

	statement, myerror := myconn.Prepare("CREATE TABLE IF NOT EXISTS user (id TEXT PRIMARY KEY, username TEXT)")
	if myerror != nil {
		log.Fatal(myerror)
	}
	statement.Exec()

	//Here we user our object for creating users and storing in database without compromissing our bussines logic
	userdao := users.UserDAO{
		Conn: imp,
	}

	userdao.CreateUser(new_user)

}
