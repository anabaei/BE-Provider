package main

import (
	"net/http"
)

// create a method, receiver and a pointer to reciever 
// the method is called routes 
// return serve mux
func (app *application) routes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/v1/books", app.getCreateBooksHandler)
	mux.HandleFunc("/v1/books/", app.getUpdateDeleteBooksHandler)
    return mux 
}
