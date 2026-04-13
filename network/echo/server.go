package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func StartServer() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		user := User{
			ID:   id,
			Name: "User " + id,
		}
		return c.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
