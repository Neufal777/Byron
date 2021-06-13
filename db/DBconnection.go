package db

import "database/sql"

const (
	HOST   = "127.0.0.1"
	DBNAME = "byrondb"
	DBUSER = "root"
	DBPASS = "root"
	DB     = "mysql"
	PORT   = "3306"
)

var (
	connectionInformation = DBUSER + ":" + DBPASS + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME
)

func ConnectDb(dbName string, connectionData string) (*sql.DB, error) {

	db, err := sql.Open(dbName, connectionData)

	return db, err
}
