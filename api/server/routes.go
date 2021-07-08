package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("Home"))
}

func GetBook(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Book"))
}

func (s *Server) InitialiseRoutes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", Home).Methods("GET")
	// s.Router.HandleFunc("/book", GetBook).Methods("GET")
}