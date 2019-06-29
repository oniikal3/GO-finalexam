package database

import "database/sql"

func InsertCustomer(db *sql.DB, name string, email string, status string) *sql.Row {
	query := `INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id, name, email, status`
	row := db.QueryRow(query, name, email, status)
	return row
}
