package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: port,
	}
}

func (ws *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	ws.Handlers[path] = handler
}

func (ws *WebServer) Start() {
	for path, handler := range ws.Handlers {
		ws.Router.HandleFunc(path, handler)
	}

	log.Printf("Starting server on %s\n", ws.WebServerPort)
	if err := http.ListenAndServe(ws.WebServerPort, ws.Router); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
