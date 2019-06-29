package customers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oniikal3/finalexam/database"
)

func (h *Handler) GetByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	row, err := database.QueryCustomer(h.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	cus := Customer{}
	row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	c.JSON(http.StatusOK, cus)
}

func (h *Handler) GetAllHandler(c *gin.Context) {
	rows, err := database.QueryCustomers(h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	var customers []Customer
	for rows.Next() {
		cus := Customer{}
		err := rows.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		customers = append(customers, cus)
	}
	c.JSON(http.StatusOK, customers)
}
