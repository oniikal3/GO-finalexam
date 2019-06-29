package database

import "database/sql"

func DeleteCustomer(db *sql.DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM customers WHERE id=$1;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
