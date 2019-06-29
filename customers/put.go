package customers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oniikal3/finalexam/database"
)

func (h *Handler) PutByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	row, err := database.UpdateCustomer(h.DB, id, cus.Name, cus.Email, cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("Update before scan: ", cus)
	row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	fmt.Println("Update after scan: ", cus)
	c.JSON(http.StatusOK, cus)
}
