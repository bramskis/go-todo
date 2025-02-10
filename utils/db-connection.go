package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	os.Getenv("POSTGRES_HOST"),
	os.Getenv("POSTGRES_PORT"),
	os.Getenv("POSTGRES_USER"),
	os.Getenv("POSTGRES_PASSWORD"),
	os.Getenv("POSTGRES_NAME"),
)

var db *sql.DB

func initializeDBConnection() error {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("unable to open connection to database: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("unable to ping database connection: %s", err.Error())
	}

	return nil
}

// GetDBConnection returns a connection to the Postgresql database
// This should NEVER be used without including a call to close on the *sql.DB returned
func GetDBConnection() (*sql.DB, error) {
	if db == nil {
		err := initializeDBConnection()
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
