package main

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

var (
	db      *gorm.DB
	validate *validator.Validate
)

func init() {
	// Initialize Validator
	validate = validator.New()
}

func GetUsersController(c echo.Context) error {
	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "user not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"user":     user,
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
	if err := validate.Struct(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid data",
		})
	}
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = newUser.Password
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

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid data",
		})
	}
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

func main() {
	e := echo.New()

	// Database Connection
	dsn := "root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
