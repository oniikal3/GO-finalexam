package customers

import (
	"database/sql"

	"github.com/oniikal3/finalexam/database"
)

func main() {}

type Customer struct {
	ID     int    `json: "id"`
	Name   string `json: "name"`
	Email  string `json: "email"`
	Status string `json: "status"`
}

type Handler struct {
	DB *sql.DB
}

func init() {
	// database.CreateCustomersTable()
}

func (h *Handler) CreateCustomersTable() error {
	err := database.CreateCustomersTable(h.DB)
	return err
}
