package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Password string
}

func (u *User) TableName() string {
	return "user"
}

func InitialMigration() {
	db := openDB()
	defer db.Close()
	db.AutoMigrate(&User{})
}

func openDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345678@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("Faailed to connecrt to mysql %+v\n", err)
		panic(err)
	}
	db.LogMode(true)
	return db
}

func FindUser(c *gin.Context) {
	name := c.Param("name")

	db := openDB()
	defer db.Close()

	var user User
	if err := db.First(&user, "name = ?", name).Error; err != nil {
		c.String(http.StatusNoContent, "Without the user information.")
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func AllUsers(c *gin.Context) {
	db := openDB()
	defer db.Close()

	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusNoContent, "Without users information.")
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func NewUser(c *gin.Context) {
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))
	password := c.DefaultPostForm("password", "123456")

	user := User{Name: name, Age: age, Password: password}

	db := openDB()
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	name := c.Param("name")

	db := openDB()
	defer db.Close()

	var user User
	if err := db.First(&user, "name = ?", name).Error; err != nil {
		c.String(http.StatusNoContent, "Without the user information.")
	} else {
		if err := db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, "Successfully Deleted User.")
		}
	}
}

func UpdateUser(c *gin.Context) {
	name := c.Param("name")
	age, _ := strconv.Atoi(c.Param("age"))
	password := c.Param("password")

	db := openDB()
	defer db.Close()

	var user User
	if err := db.First(&user, "name = ?", name).Error; err != nil {
		c.String(http.StatusNoContent, "Without the user information.")
	} else {
		user.Name = name
		user.Age = age
		if len(password) > 0 {
			user.Password = password
		}

		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, "Successfully Updated User.")
		}
	}
}
