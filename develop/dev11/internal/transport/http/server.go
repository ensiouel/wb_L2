package http

import (
	"calendar/internal/transport/http/handler"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (server *Server) Handle(eventHandler *handler.EventHandler) *Server {
	eventHandler.Register(server.mux)
	return server
}

func (server *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, server.mux)
}
