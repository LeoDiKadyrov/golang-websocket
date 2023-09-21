package postgresdb

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

func Postgresqdb() {
	dbURI := "user=postgres dbname=gowebsocket sslmode=disable"

	var err error
	PostgresDB, err = sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	if err = PostgresDB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database: ", PostgresDB)

	/*
			CREATE TABLE users (
		    	user_id SERIAL PRIMARY KEY,
		    	username VARCHAR(255) NOT NULL UNIQUE,
		    	password_hash BYTEA NOT NULL,
		    	salt BYTEA NOT NULL
			);
	*/
}
