package helper

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDatabase() (*sqlx.DB, error) {
	username := "root"
	password := ""
	database := "daerah_service"
	ip_db := "127.0.0.1"
	port_db := "3306"

	db_conn := username + ":" + password + "@" + "tcp(" + ip_db + ":" + port_db + ")/" + database
	db, err := sqlx.Open("mysql", db_conn)
	return db, err
}