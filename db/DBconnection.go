package db

import "database/sql"

// const (
// 	HOST   = "localhost"
// 	DBNAME = "byrondb"
// 	DBUSER = "root"
// 	DBPASS = "rootroot"
// 	DB     = "mysql"
// 	PORT   = "3306"
// )

//mysql://ba8e89decaafd6:ded630b2@eu-cdbr-west-01.cleardb.com/heroku_4806f7088be7d12?reconnect=true

const (
	HOST   = "eu-cdbr-west-01.cleardb.com"
	DBNAME = "heroku_4806f7088be7d12"
	DBUSER = "ba8e89decaafd6"
	DBPASS = "ded630b2"
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
