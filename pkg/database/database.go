package database

import (
	"database/sql"
	"strconv"
	"time"
)

func New(
	host string,
	port int,
	user string,
	password string,
	dbname string,
) (*sql.DB, error) {
	// Define MySQL DSN
	dsn := user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbname

	// Open MySQL connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	db.SetMaxOpenConns(25)                  // Limit the number of open connections
	db.SetMaxIdleConns(25)                  // Limit the number of idle connections
	db.SetConnMaxLifetime(15 * time.Minute) // Limit the maximum lifetime of a connection

	return db, nil
}
