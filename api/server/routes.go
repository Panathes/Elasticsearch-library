package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Pouet(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("Pouet"))
}

func (s *Server) InitialiseRoutes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", Pouet).Methods("GET")
}