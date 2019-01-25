package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Todo - todo model
type Todo struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

// TableName - sign todo table name
func (t *Todo) TableName() string {
	return "todo"
}

// Response - response struct
type Response struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// InitialMigration - Migrate database
func InitialMigration() {
	db := openDB()
	defer db.Close()
	db.AutoMigrate(&Todo{})
}

// AllTodos - get all todos
func AllTodos(c *gin.Context) {
	db := openDB()
	defer db.Close()

	var todos []Todo
	if err := db.Find(&todos).Error; err != nil {
		response := Response{Code: 1, Message: "Get todos failed.", Data: err}
		c.JSON(http.StatusNoContent, response)
	} else {
		response := Response{Code: 0, Message: "Successfully get all todos.", Data: todos}
		c.JSON(http.StatusOK, response)
	}
}

// NewTodo - add a todo
func NewTodo(c *gin.Context) {
	name := c.PostForm("name")
	todo := Todo{Name: name}

	db := openDB()
	defer db.Close()

	if err := db.Create(&todo).Error; err != nil {
		response := Response{Code: 1, Message: "Add todo failed.", Data: err}
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := Response{Code: 0, Message: "Successfully add todo.", Data: todo}
		c.JSON(http.StatusOK, response)
	}
}

// DeleteTodo - delete todo with id
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := Response{Code: 0, Message: "id is no uint type", Data: c.Param("id")}
		c.JSON(http.StatusExpectationFailed, response)
		return
	}

	db := openDB()
	defer db.Close()

	var todo Todo
	if err := db.First(&todo, "id = ?", id).Error; err != nil {
		response := Response{Code: 1, Message: "Delete todo failed.", Data: err}
		c.JSON(http.StatusInternalServerError, response)
	} else {
		if err := db.Delete(&todo).Error; err != nil {
			response := Response{Code: 1, Message: "Delete todo failed.", Data: err}
			c.JSON(http.StatusInternalServerError, response)
		} else {
			response := Response{Code: 0, Message: "Successfully deleted the todo.", Data: todo}
			c.JSON(http.StatusOK, response)
		}
	}
}

func openDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345678@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("Failed to connect to mysql %+v\n", err)
		panic(err)
	}
	db.LogMode(true)
	return db
}
