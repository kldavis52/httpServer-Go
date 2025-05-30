package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port         int
	Router       *Router
	middlewares  []Middleware
	errorHandler ErrorHandler
	server       *http.Server
}

type ServerOption func(*Server)

func NewServer(port int) *Server {
	s := &Server{
		port:        port,
		middlewares: []Middleware{},
	}
	s.Router = NewRouter(s)
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) Run() error {
	addr := fmt.Sprint(":%d", s.port)
	s.server = &http.Server{
		Addr:    addr,
		Handler: s,
	}

	fmt.Printf("Server starting on port %d\n", s.port)
	return s.server.ListenAndServe()
}
