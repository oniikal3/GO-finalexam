package customers

import (
	"database/sql"
)

func main() {}

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type Handler struct {
	DB *sql.DB
}

func init() {
	// database.CreateCustomersTable()
}
