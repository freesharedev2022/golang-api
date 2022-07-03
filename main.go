package main

import (
	"fmt"
	"github.com/PrinceNorin/todo-go/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	echo "github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"goelster/model"
	"net/http"
	"strconv"
)

// allUsers godoc
// @Summary Get all user
// @Description Get all user item
// @Tags users
// @Accept json
// @Produce json
// @Param user body types.User true "All User"
// @Success 201 {object} types.User
// @Failure 400 {object} HTTPError
// @Router /users [get]
func allUsers(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var users []model.User
		db.Find(&users)
		fmt.Println("{}", users)
		return c.JSON(http.StatusOK, users)
	}
}

// newUser godoc
// @Summary Create a user
// @Description Create a new user item
// @Tags users
// @Accept json
// @Produce json
// @Param user body types.User true "New User"
// @Success 201 {object} types.User
// @Failure 400 {object} HTTPError
// @Router /user [post]
func newUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		Username := c.FormValue("Username")
		Password := c.FormValue("Password")
		Fullname := c.FormValue("Fullname")
		db.Create(&model.User{Username: Username, Password: Password, Fullname: Fullname})
		return c.String(http.StatusOK, Username+" user successfully created")
	}
}

func deleteUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		Username := c.Param("username")
		var user model.User
		db.Where("username = ?", Username).Find(&user)
		db.Delete(&user)
		return c.String(http.StatusOK, Username+" user successfully deleted")
	}
}

func updateUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		Username := c.Param("Username")
		Password := c.Param("Password")
		Fullname := c.Param("Fullname")
		var user model.User
		db.Where("username=?", Username).Find(&user)
		user.Fullname = Fullname
		user.Password = Password
		db.Save(&user)
		return c.String(http.StatusOK, Username+" user successfully updated")
	}
}

func usersByPage(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		size, _ := strconv.Atoi(c.QueryParam("size"))
		page, _ := strconv.Atoi(c.QueryParam("page"))
		var result []model.User
		db.Limit(size).Offset(size * (page - 1)).Find(&result)
		return c.JSON(http.StatusOK, result)
	}
}

func handleRequest(db *gorm.DB) {
	e := echo.New()
	e.HTTPErrorHandler = handler.ErrorHandler
	v1 := e.Group("/api/v1")
	{
		v1.GET("/users", allUsers(db))
		v1.GET("/user", usersByPage(db))
		v1.POST("/user", newUser(db))
		v1.DELETE("/user/:name", deleteUser(db))
		v1.PUT("/user/:name/:email", updateUser(db))
	}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8000"))
}

func initialMigration(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}

// @title User Application
// @description This is a use list management application
// @version 1.0
// @host localhost:8000
// @BasePath /api/v1
func main() {
	db, err := gorm.Open("mysql", "root:12345678@(127.0.0.1:3306)/golang_api?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		fmt.Printf("error connect to database %s", err)
	} else {
		initialMigration(db)
	}
	handleRequest(db)
}
