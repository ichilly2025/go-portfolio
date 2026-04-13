package zero

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func StartServer() {
	server := rest.MustNewServer(rest.RestConf{
		Port: 8080,
	})
	defer server.Stop()

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ping",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			httpx.OkJson(w, map[string]string{
				"message": "pong",
			})
		},
	})

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/hello",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			httpx.OkJson(w, map[string]string{
				"message": "Hello, World!",
			})
		},
	})

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/users/:id",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			// 从 URL 路径中提取 ID
			id := r.URL.Path[len("/users/"):]
			user := User{
				ID:   id,
				Name: "User " + id,
			}
			httpx.OkJson(w, user)
		},
	})

	server.Start()
}
