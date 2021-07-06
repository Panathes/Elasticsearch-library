package server

import (
	"github.com/gorilla/mux"
)

func (s *Server) InitialiseRoutes() {
	s.Router = mux.NewRouter()
}