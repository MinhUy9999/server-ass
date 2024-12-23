package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Định nghĩa struct Product để đại diện cho 1 sản phẩm
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Tạo sẵn 1 slice các sản phẩm mẫu để giả lập dữ liệu
var products = []Product{
	{ID: 1, Name: "Sản phẩm A", Price: 10000},
	{ID: 2, Name: "Sản phẩm B", Price: 20000},
	{ID: 3, Name: "Sản phẩm C", Price: 30000},
}

func main() {
	r := gin.Default()

	// GET /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// GET /hello/:name
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	// GET /user/:id
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "User with ID = %s", id)
	})

	// GET /product/:id (lấy 1 sản phẩm theo ID)
	r.GET("/product/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Product with ID = %s", id)
	})

	// GET /cart/:id
	r.GET("/cart/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Cart with ID = %s", id)
	})

	// GET /order/:id
	r.GET("/order/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Order with ID = %s", id)
	})

	// POST /user (tạo mới user với các trường name, email, password)
	r.POST("/user", func(c *gin.Context) {
		type User struct {
			Name     string `json:"name"     binding:"required"`
			Email    string `json:"email"    binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}

		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Ở đây bạn có thể thêm logic lưu CSDL, kiểm tra hợp lệ,...
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"user":    user,
		})
	})

	// MỚI THÊM: GET /products - trả về danh sách tất cả sản phẩm
	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	r.Run(":3000")
}
