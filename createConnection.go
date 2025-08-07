package databasesmng

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/milbertk/class"
)

var (
	db    *sql.DB
	once  sync.Once
	dbErr error
)

// GetPostgresDB initializes or returns the existing DB connection
func CreateConnection() (*sql.DB, error) {
	once.Do(func() {
		// Load config
		reader, err := class.NewJSONReader("./connection.json")
		if err != nil {
			dbErr = fmt.Errorf("❌ Failed to read JSON config: %w", err)
			return
		}

		// Extract values
		host, ok1 := reader.GetValue("host")
		port, ok2 := reader.GetValue("port")
		user, ok3 := reader.GetValue("user")
		pass, ok4 := reader.GetValue("pass")
		dbname, ok5 := reader.GetValue("database")

		if !(ok1 && ok2 && ok3 && ok4 && ok5) {
			dbErr = fmt.Errorf("❌ Missing DB connection config")
			return
		}

		conn, err := NewPostgresConnector(host, port, user, pass, dbname)
		if err != nil {
			dbErr = fmt.Errorf("❌ Failed to connect to PostgreSQL: %w", err)
			return
		}

		db = conn.DB
	})
	return db, dbErr
}
