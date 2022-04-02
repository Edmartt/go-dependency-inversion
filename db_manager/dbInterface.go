package dbmanager

import "database/sql"

type IDatabaseConnection interface{

	 GetDBConnection() *sql.DB

}