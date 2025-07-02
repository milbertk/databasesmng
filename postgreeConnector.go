package databasesmng

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresConnector struct {
	DB *sql.DB
}

// NewPostgresConnector creates and connects to PostgreSQL
func NewPostgresConnector(host, port, user, password, dbname string) (*PostgresConnector, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresConnector{DB: db}, nil
}

// Close closes the DB connection
func (p *PostgresConnector) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}
