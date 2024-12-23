package main

import (
	"net/http"
	"strconv"
	"sync"

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

// Định nghĩa struct CartItem để đại diện cho 1 sản phẩm trong giỏ hàng
type CartItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

// Định nghĩa struct Cart để đại diện cho giỏ hàng
type Cart struct {
	ID    int        `json:"id"`
	Items []CartItem `json:"items"`
}

// Biến để lưu trữ các giỏ hàng, sử dụng map để truy xuất nhanh theo ID
var carts = make(map[int]*Cart)
var cartMutex = &sync.Mutex{} // Mutex để đảm bảo tính thread-safe
var nextCartID = 1

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
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		for _, product := range products {
			if product.ID == id {
				c.JSON(http.StatusOK, product)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	// GET /cart/:id
	r.GET("/cart/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
			return
		}

		cartMutex.Lock()
		cart, exists := carts[id]
		cartMutex.Unlock()

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
			return
		}

		c.JSON(http.StatusOK, cart)
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

	// GET /products - trả về danh sách tất cả sản phẩm
	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	// POST /cart - tạo giỏ hàng mới
	r.POST("/cart", func(c *gin.Context) {
		cartMutex.Lock()
		defer cartMutex.Unlock()

		cart := &Cart{
			ID:    nextCartID,
			Items: []CartItem{},
		}
		carts[nextCartID] = cart
		nextCartID++

		c.JSON(http.StatusOK, cart)
	})

	// POST /cart/:id/add - thêm sản phẩm vào giỏ hàng
	r.POST("/cart/:id/add", func(c *gin.Context) {
		idStr := c.Param("id")
		cartID, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
			return
		}

		type AddItemRequest struct {
			ProductID int `json:"product_id" binding:"required"`
			Quantity  int `json:"quantity" binding:"required,gt=0"`
		}

		var req AddItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Tìm sản phẩm theo ID
		var product *Product
		for _, p := range products {
			if p.ID == req.ProductID {
				product = &p
				break
			}
		}

		if product == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		cartMutex.Lock()
		defer cartMutex.Unlock()

		cart, exists := carts[cartID]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
			return
		}

		// Kiểm tra nếu sản phẩm đã có trong giỏ hàng thì tăng số lượng
		for i, item := range cart.Items {
			if item.Product.ID == product.ID {
				cart.Items[i].Quantity += req.Quantity
				c.JSON(http.StatusOK, cart)
				return
			}
		}

		// Nếu sản phẩm chưa có trong giỏ hàng thì thêm mới
		cart.Items = append(cart.Items, CartItem{
			Product:  *product,
			Quantity: req.Quantity,
		})

		c.JSON(http.StatusOK, cart)
	})

	// POST /cart/:id/remove - xóa sản phẩm khỏi giỏ hàng
	r.POST("/cart/:id/remove", func(c *gin.Context) {
		idStr := c.Param("id")
		cartID, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
			return
		}

		type RemoveItemRequest struct {
			ProductID int `json:"product_id" binding:"required"`
		}

		var req RemoveItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cartMutex.Lock()
		defer cartMutex.Unlock()

		cart, exists := carts[cartID]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
			return
		}

		// Tìm và xóa sản phẩm khỏi giỏ hàng
		for i, item := range cart.Items {
			if item.Product.ID == req.ProductID {
				// Xóa phần tử tại vị trí i
				cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
				c.JSON(http.StatusOK, cart)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found in cart"})
	})

	r.Run(":3000")
}
