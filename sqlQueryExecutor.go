package databasesmng

import (
	"database/sql"
	"encoding/json"
)

type QueryExecutor struct {
	DB *sql.DB
}

// NewQueryExecutor receives an existing connection
func NewQueryExecutor(db *sql.DB) *QueryExecutor {
	return &QueryExecutor{DB: db}
}

// ExecuteQuery runs a SELECT query and returns JSON result
func (q *QueryExecutor) ExecuteQuery(query string) (string, error) {
	rows, err := q.DB.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	results := []map[string]interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return "", err
		}

		entry := map[string]interface{}{}
		for i, col := range columns {
			entry[col] = values[i]
		}
		results = append(results, entry)
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// ExecuteNonQuery runs INSERT, UPDATE, DELETE statements
func (q *QueryExecutor) ExecuteNonQuery(query string) (int64, error) {
	result, err := q.DB.Exec(query)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
