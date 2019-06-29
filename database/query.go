package database

import "database/sql"

func QueryCustomer(db *sql.DB, id int) (*sql.Row, error) {
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customers WHERE id=$1;")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	return row, err
}

func QueryCustomers(db *sql.DB) (*sql.Rows, error) {
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customers;")
	if err != nil {
		return nil, err
	}
	return stmt.Query()
}
