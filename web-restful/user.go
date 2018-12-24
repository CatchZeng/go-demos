package main

import (
	"fmt"
	"net/http"

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

func AllUsers(c *gin.Context) {
	db := openDB()
	defer db.Close()

	var users []User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}

func NewUser(c *gin.Context) {
	c.String(http.StatusOK, "New User Endpoint Hint")
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusOK, "Delete User Endpoint Hint")
}

func UpdateUser(c *gin.Context) {
	c.String(http.StatusOK, "Update User Endpoint Hint")
}
