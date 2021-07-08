package controllers

import (
	"net/http"
)

func GetBook(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Book"))
}