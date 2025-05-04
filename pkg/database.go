package pkg

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func MustConnectDatabase(url string) *sql.DB {

	const op = "pkg.database.driver"

	db, err := sql.Open("mysql", url)

	if err != nil {
		panic("Can not connect to database: " + err.Error())
	}

	return db
}
