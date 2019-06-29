package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/oniikal3/finalexam/customers"
)

func main() {
	r := setupRouter()
	r.Run(getPort())
}

func authMiddleware(c *gin.Context) {
	// token := c.GetHeader("Authorization")
	c.Next()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(authMiddleware)
	handler := customers.Handler{
		DB: startConnection(),
	}
	err := handler.CreateCustomersTable()
	if err != nil {
		// return error
	}
	c := r.Group("/customers")
	{
		c.GET("/:id", handler.GetByIdHandler)
		c.GET("/", handler.GetAllHandler)
		c.POST("/", handler.PostHandler)
		c.PUT("/:id", handler.PutByIdHandler)
		c.DELETE("/:id", handler.DeleteByIdHandler)
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
