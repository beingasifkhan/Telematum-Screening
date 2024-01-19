package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// CreateConnection creates a connection to a PostgreSQL database
func CreateConnection() (*sql.DB, error) {
	connStr := "host=localhost user=postgres password=25205089pAr@ dbname=screening port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("error opening database connection:", err)
		return nil, err
	}

	// Checking if the db connection is established
	err = db.Ping()
	if err != nil {
		log.Println("error pinging database:", err)
		db.Close()
		return nil, err
	}

	fmt.Println("connected to the database")
	return db, nil
}
