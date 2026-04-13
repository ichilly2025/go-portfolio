package echo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func setupRouter() *echo.Echo {
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

	return e
}

func TestPingEndpoint(t *testing.T) {
	e := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)

	if response["message"] != "pong" {
		t.Errorf("Expected message 'pong', got '%s'", response["message"])
	}
}

func TestHelloEndpoint(t *testing.T) {
	e := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)

	if response["message"] != "Hello, World!" {
		t.Errorf("Expected message 'Hello, World!', got '%s'", response["message"])
	}
}

func TestGetUserEndpoint(t *testing.T) {
	e := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/users/123", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var user User
	json.Unmarshal(rec.Body.Bytes(), &user)

	if user.ID != "123" {
		t.Errorf("Expected user ID '123', got '%s'", user.ID)
	}

	if user.Name != "User 123" {
		t.Errorf("Expected user name 'User 123', got '%s'", user.Name)
	}
}

func TestNotFoundEndpoint(t *testing.T) {
	e := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", rec.Code)
	}
}
