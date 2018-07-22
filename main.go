package main

import (
	//...
	// The libn/pq driver is used for postgres
	"database/sql"

	_ "github.com/lib/pq"
	//...
)

func main() {
	// ...
	connString := "dbname=<your main db name> sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})
	//...
}
