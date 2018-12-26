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

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	r.GET("/users", AllUsers)

	r.GET("/user/:name", FindUser)

	r.POST("/user", NewUser)

	r.DELETE("/user/:name", DeleteUser)

	r.PUT("/user/:name/:age/:password", UpdateUser)

	r.Run(":8080")
}
