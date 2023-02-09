package postgres

import (
	"database/sql"
	"log"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(connectionString string) *Postgres {
	p := &Postgres{}
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed to open database connection, %v", err)
	}
	p.DB = db
	err = p.DB.Ping()
	if err != nil {
		log.Fatalf("failed to ping database, %v", err)
	}
	return p
}
