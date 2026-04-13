package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPServer(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	expected := "Hello, World!"
	if string(body) != expected {
		t.Errorf("Expected body %q, got %q", expected, string(body))
	}
}
