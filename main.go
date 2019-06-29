package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/oniikal3/finalexam/customers"
	"github.com/oniikal3/finalexam/database"
)

func main() {
	h := setupEnv()
	r := setupRouter(h)
	defer h.DB.Close()
	r.Run(getPort())
}

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": http.StatusText(http.StatusUnauthorized),
		})
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized: invalid token"))
		return
	}
	c.Next()
}

func setupEnv() *customers.Handler {
	db := startConnection()
	handler := &customers.Handler{
		DB: db,
	}
	err := database.CreateCustomersTable(db)
	if err != nil {
		panic(err)
	}
	return handler
}

func setupRouter(h *customers.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(authMiddleware)
	c := r.Group("/customers")
	{
		c.GET("/:id", h.GetByIdHandler)
		c.GET("/", h.GetAllHandler)
		c.POST("/", h.PostHandler)
		c.PUT("/:id", h.PutByIdHandler)
		c.DELETE("/:id", h.DeleteByIdHandler)
	}
	return r
}

func getPort() string {
	return ":2019"
}

func startConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error", err.Error())
	}
	return db
}
