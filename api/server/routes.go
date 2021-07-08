package server

import (
	"github.com/Elasticsearch-library/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("Home"))
}

func (s *Server) InitialiseRoutes(controller *controllers.Controller) {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", Home).Methods("GET")
	s.Router.HandleFunc("/", controller.CreateBook).Methods("POST")
	// s.Router.HandleFunc("/book", GetBook).Methods("GET")
}