package crud

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// todo: read from env
const (
	host     = "localhost"
	port     = 5432
	user     = "yourusername"
	password = "yourpassword"
	dbname   = "yourdbname"
)

func openDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Unable to open database: %v\n", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Unable to connect to the database: %v\n", err)
		return nil, err
	}
	return db, err
}
