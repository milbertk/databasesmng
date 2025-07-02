package databasesmng

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnector struct {
	DB *sql.DB
}

// NewMySQLConnector creates and connects to a MySQL database
func NewMySQLConnector(host, port, user, password, dbname string) (*MySQLConnector, error) {
	// Example format: user:password@tcp(host:port)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &MySQLConnector{DB: db}, nil
}

// Close closes the MySQL DB connection
func (m *MySQLConnector) Close() {
	if m.DB != nil {
		m.DB.Close()
	}
}
