package pkg

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	Instance *sql.DB
}

func MustNewDbInstance(url string) *Db {
	db, err := sql.Open("mysql", url)

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	return &Db{
		db,
	}
}
