package db

import "database/sql"

func ConnectDb(dbName string, connectionData string) (*sql.DB, error) {

	db, err := sql.Open(dbName, connectionData)

	return db, err
}
