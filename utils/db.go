package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Instance *sql.DB
}

// To initialize postgres db
func (pg *PostgresDB) InitializeDB(connectionString string) (*sql.DB, error) {
	// Connect to db
	var err error
	pg.Instance, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	// Ping db and throw error if not connected
	if err = pg.Instance.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected!")
	return pg.Instance, nil
}
