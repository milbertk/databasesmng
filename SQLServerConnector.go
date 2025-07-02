package databasesmng

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type SQLServerConnector struct {
	DB *sql.DB
}

// NewSQLServerConnector creates and connects to a SQL Server database
func NewSQLServerConnector(host, port, user, password, dbname string) (*SQLServerConnector, error) {
	connString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s;",
		host, port, user, password, dbname)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &SQLServerConnector{DB: db}, nil
}

// Close closes the SQL Server DB connection
func (s *SQLServerConnector) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
}
