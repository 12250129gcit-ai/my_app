package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	// host = "db" because that's the service name in docker-compose
	connStr := "host=dpg-d7khg01o3t8c73co4g60-a port=5432 user=myuser password=mCmBkRFF1a52wJueTnHzhADss5PQe6PB dbname=myapp_db_x1w9 sslmode=disable"
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
