package database

import "database/sql"

type Database struct {
	Instance *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{Instance: db}
}
