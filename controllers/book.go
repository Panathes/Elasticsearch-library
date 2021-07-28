package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Elasticsearch-library/models"
	u "github.com/Elasticsearch-library/utils"
	"github.com/elastic/go-elasticsearch"
)

type Controller struct {
	Es *elasticsearch.Client
}

func (c *Controller) WriteJson(w http.ResponseWriter, value interface{}) {
	output, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(output)
}


func (c *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	books := models.Elastic{
		Es: c.Es,
		Book: models.Books{},
	}
	err = json.Unmarshal(body, &books.Book)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	resp := books.AddBook()
	c.WriteJson(w, resp)
}

func (c *Controller) GetBook(w http.ResponseWriter, r *http.Request) {

	books := models.Elastic{
		Es: c.Es,
		Book: models.Books{},
	}

	author := r.FormValue("author")
	title := r.FormValue("title")
	abstract := r.FormValue("abstract")

	resp := books.GetBook(&author, &title, &abstract)
	c.WriteJson(w, resp)

}