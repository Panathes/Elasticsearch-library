package models

import (
	"context"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"strings"
)

type Books struct {
	Title         	string  `json:"title"`
	Author          string  `json:"author"`
	Abstract 		string 	`json:"abstract"`
}

type Elastic struct {
	Es *elasticsearch.Client
	Book Books
}

func (e *Elastic) AddBook() *esapi.Response {
	var b strings.Builder
	b.WriteString(`{"title" : "`)
	b.WriteString(e.Book.Title)
	b.WriteString(`","author" : "`)
	b.WriteString(e.Book.Author)
	b.WriteString(`","abstract" : "`)
	b.WriteString(e.Book.Abstract)
	b.WriteString(`"}`)

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      "book",
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), e.Es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status())
	}

	return res
}
