package restful

import (
	"fmt"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func StartServer() {
	ws := new(restful.WebService)
	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/hello").To(hello))
	ws.Route(ws.GET("/users/{id}").To(getUser))

	restful.Add(ws)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	resp.WriteAsJson(map[string]string{"message": "Hello, World!"})
}

func getUser(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	user := User{
		ID:   id,
		Name: "User " + id,
	}
	resp.WriteEntity(user)
}
