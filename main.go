package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Định nghĩa struct User để ánh xạ dữ liệu gửi lên
type User struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// GET: /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// GET: /hello/:name
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	// GET: /user/:id
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "User with ID = %s", id)
	})

	// GET: /product/:id
	r.GET("/product/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Product with ID = %s", id)
	})

	// GET: /cart/:id
	r.GET("/cart/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Cart with ID = %s", id)
	})

	// GET: /order/:id
	r.GET("/order/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Order with ID = %s", id)
	})

	// POST: /user
	// Tạo mới user với name, email, password
	r.POST("/user", func(c *gin.Context) {
		var user User
		// Gọi ShouldBindJSON để parse dữ liệu JSON và gán cho 'user'
		if err := c.ShouldBindJSON(&user); err != nil {
			// Nếu dữ liệu sai format hoặc thiếu trường, trả về lỗi
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Sau khi parse thành công, có thể xử lý lưu database hoặc logic khác
		// Ở đây ví dụ trả về user vừa nhận được
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"user":    user,
		})
	})

	r.Run(":3000")
}
