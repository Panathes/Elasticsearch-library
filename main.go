package main

// import (
// 	"log"

// 	"github.com/elastic/go-elasticsearch/v8"
// )

//   func main() {
//     es, _ := elasticsearch.NewDefaultClient()
//     log.Println(elasticsearch.Version)
//     log.Println(es.Info())
//   }
import (
	"github.com/Elasticsearch-library/api/server"
)

func main()  {
	api := server.Server{}
	//api.InitialiseRoutes()
	api.Run()
}