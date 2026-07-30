package database

import (
	"database/sql"

	_ "github.com/lib/pq" //psql driver
)

//connecting go to postgres database

func NewPostgresDB() (*sql.DB, error) {
	//connection string
	connStr := "postgres://postgres:postgres@localhost:5432/email_dispatcher?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	//ping the database to check if connection is established
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
