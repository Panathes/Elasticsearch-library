package main

import (
	"github.com/Elasticsearch-library/api/server"
)

func main()  {
	api := server.Server{}
	//api.InitialiseRoutes()
	api.Run()
}