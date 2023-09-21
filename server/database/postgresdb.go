package postgresdb

import (
	"fmt"
	"sync"
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

var (
	instance *Database
	once sync.Once
)

func NewDatabase() (*Database, error) {
	dbURI := "user=postgres dbname=gowebsocket sslmode=disable"

	once.Do(func() {
		db, err := sql.Open("postgres", dbURI)
		if err != nil {
			panic(err)
		}

		if err = db.Ping(); err != nil {
			panic(err)
		}
		fmt.Println("Connected to database: ", db)
		instance = &Database{db}
	})

	return instance, nil
}

func (d *Database) getDB() *sql.DB {
	if instance == nil {
		panic("Database connection is not initialized")
	}
	return d.db
}

func (d *Database) Close() error {
	return d.db.Close()
}
