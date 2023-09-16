package postgresqdb

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Postgresqdb() {
	defer db.Close()
	dbURI := "user=postgres dbname=gowebsocket sslmode=disable password=admin"

	var err error
	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database: ", db)

	// Just test inserting - all working but crashes with error: "panic: runtime error: invalid memory address or nil pointer dereference"

	// _, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", "john", "hashed_password")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}