package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oniikal3/finalexam/database"
)

func (h *Handler) PostHandler(c *gin.Context) {
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	row := database.InsertCustomer(h.DB, cus.Name, cus.Email, cus.Status)
	row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	c.JSON(http.StatusCreated, cus)
}
