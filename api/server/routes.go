package server

import (
	"net/http"

	"github.com/Elasticsearch-library/controllers"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("Home"))
}

func (s *Server) InitialiseRoutes(controller *controllers.Controller) {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", Home).Methods("GET")
	s.Router.HandleFunc("/", controller.CreateBook).Methods("POST")
	s.Router.HandleFunc("/book", controller.GetBook).Methods("GET")
}