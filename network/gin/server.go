package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func StartServer() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := User{
			ID:   id,
			Name: "User " + id,
		}
		c.JSON(http.StatusOK, user)
	})

	r.Run(":8080")
}
