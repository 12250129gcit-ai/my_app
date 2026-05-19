package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	// host = "db" because that's the service name in docker-compose
	connStr := "host=db port=5432 user=postgres password=postgres dbname=my_db sslmode=disable"
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// Verify connection
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	println("Database successfully connected")
}
