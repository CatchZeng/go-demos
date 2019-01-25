package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("GO API")

	InitialMigration()
	handleRequests()
}

func handleRequests() {
	r := gin.Default()
	r.Static("/static", "./public")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "static")
	})

	r.GET("/todos", AllTodos)

	r.POST("/todo", NewTodo)

	r.DELETE("/todo/:id", DeleteTodo)

	r.Run(":8080")
}
