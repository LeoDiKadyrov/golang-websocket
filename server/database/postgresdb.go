package postgresdb

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

type singletonDB struct {
	DB  *sql.DB
	err error
}

var singleInstanceDB *singletonDB

func (s *singletonDB) NewDatabase() error {
	dbURI := "user=postgres dbname=gowebsocket sslmode=disable"

	s.DB, s.err = sql.Open("postgres", dbURI)
	if s.err != nil {
		panic(s.err)
	}

	if s.err = s.DB.Ping(); s.err != nil {
		panic(s.err)
	}
	fmt.Println("Connected to database: ", s.DB)

	return nil
}

func GetInstanceDB() *singletonDB {
	if singleInstanceDB == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstanceDB == nil {
			fmt.Println("Creating a single instance of db now")
			singleInstanceDB = &singletonDB{}
			singleInstanceDB.NewDatabase()
		} else {
			fmt.Println("Single instance has already been created")
		}
	} else {
		fmt.Println("Single instance has already been created")
	}

	return singleInstanceDB
}

func Close() error {
	return singleInstanceDB.DB.Close()
}
