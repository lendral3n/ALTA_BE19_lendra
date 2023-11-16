package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

type Product struct {
	gorm.Model
	ProductName string  `json:"productName" form:"productName"`
	Description string  `json:"description" form:"description"`
	Stok        int     `json:"stok" form:"stok"`
	Price       float64 `json:"price" form:"price"`
}

type Customer struct {
	gorm.Model
	CustomerName string `json:"customerName" form:"customerName"`
	Address      string `json:"address" form:"address"`
	Phone        string `json:"phone" form:"phone"`
}

var db *gorm.DB

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	
	user.Role = "pegawai"

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error creating user",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}


func GetUsersController(c echo.Context) error {
	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "user not found",
		})
	}
	db.Delete(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user by id",
		"user":     user,
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "user not found",
		})
	}
	newUser := User{}
	c.Bind(&newUser)
	user.Username = newUser.Username
	user.Password = newUser.Password
	user.Role = newUser.Role
	if err := db.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error updating user",
			"error":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user by id",
		"user":     user,
	})
}

func CreateProductController(c echo.Context) error {
	product := Product{}
	c.Bind(&product)
	if err := db.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error creating product",
			"error":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create product",
		"product":  product,
	})
}

func GetProductsController(c echo.Context) error {
	var products []Product
	db.Find(&products)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all products",
		"products": products,
	})
}

func DeleteProductController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "product not found",
		})
	}
	db.Delete(&product)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete product by id",
		"product":  product,
	})
}

func UpdateProductController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "product not found",
		})
	}
	newProduct := Product{}
	c.Bind(&newProduct)
	product.ProductName = newProduct.ProductName
	product.Description = newProduct.Description
	product.Stok = newProduct.Stok
	product.Price = newProduct.Price
	if err := db.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error updating product",
			"error":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update product by id",
		"product":  product,
	})
}


func CreateCustomerController(c echo.Context) error {
	customer := Customer{}
	c.Bind(&customer)
	if err := db.Create(&customer).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error creating customer",
			"error":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create customer",
		"customer": customer,
	})
}


func GetCustomersController(c echo.Context) error {
	var customers []Customer
	db.Find(&customers)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get all customers",
		"customers": customers,
	})
}

func DeleteCustomerController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer Customer
	if err := db.First(&customer, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "customer not found",
		})
	}
	db.Delete(&customer)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete customer by id",
		"customer": customer,
	})
}


func UpdateCustomerController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer Customer
	if err := db.First(&customer, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "customer not found",
		})
	}

	newCustomer := Customer{}
	c.Bind(&newCustomer)

	customer.CustomerName = newCustomer.CustomerName
	customer.Address = newCustomer.Address
	customer.Phone = newCustomer.Phone

	if err := db.Save(&customer).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error updating customer",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update customer by id",
		"customer": customer,
	})
}

func main() {
	e := echo.New()

	dsn := "root:@tcp(localhost:3306)/project1api?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Customer{})

	e.POST("/users", CreateUserController)
	e.GET("/users", GetUsersController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.POST("/products", CreateProductController)
	e.GET("/products", GetProductsController)
	e.DELETE("/products/:id", DeleteProductController)
	e.PUT("/products/:id", UpdateProductController)

	e.POST("/customers", CreateCustomerController)
	e.GET("/customers", GetCustomersController)
	e.DELETE("/customers/:id", DeleteCustomerController)
	e.PUT("/customers/:id", UpdateCustomerController)

	e.Logger.Fatal(e.Start(":8000"))
}

