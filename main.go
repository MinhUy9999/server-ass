package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})
	// viết thêm user
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "User with ID = %s", id)
	})
	// viết thêm product
	r.GET("/product/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Product with ID = %s", id)
	})
	// viết thêm cart
	r.GET("/cart/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Cart with ID = %s", id)
	})
	// viết thêm order
	r.GET("/order/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Order with ID = %s", id)
	})
	r.Run(":3000")
}
