package restful

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	restful "github.com/emicklei/go-restful/v3"
)

func TestHelloEndpoint(t *testing.T) {
	ws := new(restful.WebService)
	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/hello").To(hello))

	container := restful.NewContainer()
	container.Add(ws)

	req := httptest.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()

	container.ServeHTTP(w, req)

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
	ws := new(restful.WebService)
	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/users/{id}").To(getUser))

	container := restful.NewContainer()
	container.Add(ws)

	req := httptest.NewRequest("GET", "/api/users/123", nil)
	w := httptest.NewRecorder()

	container.ServeHTTP(w, req)

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
