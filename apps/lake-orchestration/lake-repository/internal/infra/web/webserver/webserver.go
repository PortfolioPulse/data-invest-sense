package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	return &WebServer{
		Router:        router,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, method string, pattern string, handler http.HandlerFunc) {
	fmt.Printf("Adding handler %s %s\n", method, path)
	switch method {
	case "GET":
		s.Router.Get(pattern, handler)
	case "POST":
		s.Router.Post(pattern, handler)
	case "PUT":
		s.Router.Put(pattern, handler)
	case "DELETE":
		s.Router.Delete(pattern, handler)
	}
}

func (s *WebServer) Start() {
	http.ListenAndServe(fmt.Sprintf(":%s", s.WebServerPort), s.Router)
}
