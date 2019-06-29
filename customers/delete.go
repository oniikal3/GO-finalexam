package customers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oniikal3/finalexam/database"
)

func (h *Handler) DeleteByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	err = database.DeleteCustomer(h.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}
