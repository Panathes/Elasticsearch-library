package server

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Router 	*mux.Router
	Logger 	*log.Logger
}

func (s *Server) Run() *Server {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("FRONT_URL"), os.Getenv("DOCUMENTATION_URL")},
		AllowedHeaders: []string{"Authorization", "Content-Type", "accept"},
		AllowedMethods: []string{"POST", "GET", "PUT"},
		AllowCredentials: true,
		Debug: false,
	})
	handler := c.Handler(s.Router)
	// defer s.DB.Close()
	err := http.ListenAndServe(os.Getenv("PORT"), handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return s
}
