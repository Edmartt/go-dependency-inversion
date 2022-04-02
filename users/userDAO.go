package users

import (
	dbmanager "SQLITE/db_manager"
	"log"
)

type UserDAO struct {
	Conn dbmanager.DBImplementation //We need an object that implements user's interface, here we pass the db manager implementation we want
}

func (dao UserDAO) CreateUser(user User) {
	query := "INSERT INTO user (id, username) VALUES (?, ?)"

	connection := dao.Conn.Conn.GetDBConnection()

	preparedStatement, myerror := connection.Prepare(query)

	if myerror != nil {
		log.Fatal(myerror)
	}
	preparedStatement.Exec(user.Id, user.Username)

}
