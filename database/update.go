package database

import "database/sql"

func UpdateCustomer(db *sql.DB, id int, name string, email string, status string) (*sql.Row, error) {
	stmt, err := db.Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1;")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id, name, email, status)
	return row, err
}
