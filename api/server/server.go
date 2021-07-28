package server

import (
	"github.com/Elasticsearch-library/controllers"
	"github.com/elastic/go-elasticsearch"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	Router 	*mux.Router
	elastic *elasticsearch.Client
}

func (s *Server) Run() *Server {
	s.Router = mux.NewRouter()
	controller := &controllers.Controller{
		Es: InitEs(),
	}
	s.InitialiseRoutes(controller)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("FRONT_URL"), os.Getenv("DOCUMENTATION_URL")},
		AllowedHeaders: []string{"Authorization", "Content-Type", "accept"},
		AllowedMethods: []string{"POST", "GET", "PUT"},
		AllowCredentials: true,
		Debug: false,
	})
	handler := c.Handler(s.Router)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return s
}
