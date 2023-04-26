package controllers

import (
	"fmt"
	"golang-api/middleware"
	"golang-api/models"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func getValueOf(obj any, field string) any {
	value := reflect.ValueOf(obj)
	val := reflect.Indirect(value).FieldByName(field)
	if val.IsValid() {
		return val
	}
	return nil
}

func RegisterUser(c *gin.Context) {
	var input models.CreateUser
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("email=?", input.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error"})
		return
	}

	userData := models.User{Name: input.Name, Password: string(hashedPassword), Email: input.Email}

	models.DB.Create(&userData)

	c.JSON(http.StatusCreated, gin.H{"data": userData})
}

func LoginUser(c *gin.Context) {
	var input models.LoginUser
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("email=?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not exist"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password in valid"})
		return
	}
	tokenString, exp := middleware.GenerateJWT(user.Email, user.Name, user.ID)
	var response = models.LoginResponse{tokenString, exp, user}
	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func UserInfo(c *gin.Context) {
	var claims any
	claims, _ = c.Get("user")
	c.JSON(http.StatusOK, gin.H{"data": claims})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	var input models.UpdateUser
	var claims any
	claims, _ = c.Get("user")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var id = getValueOf(claims, "Id")
	id = fmt.Sprintf("%d", id)
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error"})
		return
	}

	models.DB.Model(&user).Updates(models.User{Name: input.Name, Password: string(hashedPassword)})

	c.JSON(http.StatusOK, gin.H{"data": user})
}
