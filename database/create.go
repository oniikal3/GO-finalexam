package database

import "database/sql"

// CreateCustomersTable will create table `customers` if not exist
func CreateCustomersTable(db *sql.DB) error {
	stmt := `
	CREATE TABLE IF NOT EXISTS customers(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`
	_, err := db.Exec(stmt)
	return err
}
