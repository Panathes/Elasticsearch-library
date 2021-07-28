package models

import (
	"context"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
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

func Request(author *string, title *string, abstract *string) string  {
	str := ""
	if *author != "" {
		str += "author: " + *author + ", "
	}
	if *title != "" {
		str += "title: " + *title + ", "
	}
	if  *abstract != "" {
		str += "abstract: " + *abstract + ", "
	}
	return str
}

func (e *Elastic) GetBook(author *string, title *string, abstract *string) *esapi.Response {
	res, err := e.Es.Search(
		e.Es.Search.WithContext(context.Background()),
		e.Es.Search.WithIndex("books"),
		e.Es.Search.WithQuery(Request(author, title, abstract)),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error getting response")
	}

	return res
}
