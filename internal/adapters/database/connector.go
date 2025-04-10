package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func OpenConnectionDB() (*sql.DB, error) {
	var dbURl string //todo: configurar .env

	db, err := sql.Open("postgres", dbURl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return db, err
}
