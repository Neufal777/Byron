package db

import "database/sql"

const (
	HOST   = "localhost"
	DBNAME = "byrondb"
	DBUSER = "root"
	DBPASS = "rootroot"
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
