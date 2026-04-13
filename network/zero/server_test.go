package zero

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/router"
)

func setupRouter() http.Handler {
	r := router.NewRouter()

	r.Handle(http.MethodGet, "/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, map[string]string{
			"message": "pong",
		})
	}))

	r.Handle(http.MethodGet, "/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, map[string]string{
			"message": "Hello, World!",
		})
	}))

	r.Handle(http.MethodGet, "/users/:id", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// 从 URL 路径中提取 ID
		id := req.URL.Path[len("/users/"):]
		user := User{
			ID:   id,
			Name: "User " + id,
		}
		httpx.OkJson(w, user)
	}))

	return r
}

func TestPingEndpoint(t *testing.T) {
	handler := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["message"] != "pong" {
		t.Errorf("Expected message 'pong', got '%s'", response["message"])
	}
}

func TestHelloEndpoint(t *testing.T) {
	handler := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["message"] != "Hello, World!" {
		t.Errorf("Expected message 'Hello, World!', got '%s'", response["message"])
	}
}

func TestGetUserEndpoint(t *testing.T) {
	handler := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/users/123", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var user User
	json.Unmarshal(w.Body.Bytes(), &user)

	if user.ID != "123" {
		t.Errorf("Expected user ID '123', got '%s'", user.ID)
	}

	if user.Name != "User 123" {
		t.Errorf("Expected user name 'User 123', got '%s'", user.Name)
	}
}

func TestNotFoundEndpoint(t *testing.T) {
	handler := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}
