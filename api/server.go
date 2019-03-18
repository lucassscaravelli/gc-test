package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	mux *mux.Router
}

func newServer(mux *mux.Router) *Server {
	server := Server{mux}
	server.routes()
	return &server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
