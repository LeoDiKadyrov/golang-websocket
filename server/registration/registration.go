package registration

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

// should I connect to db here and do all work here? Or split connection to db into one file and all registration / auth handling to another?

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	http.ServeFile(w, r, "client/registration.html")
}

func Registration() {
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

	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", "john", "hashed_password")
	if err != nil {
		log.Fatal(err)
	}
}
