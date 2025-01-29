package main


import (
	"fmt"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request){
	if(r.Method != http.MethodGet) {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed) ,http.StatusMethodNotAllowed )
		return
	  }
	  w.Header().Set("Content-Type", "text/plain")
	  fmt.Fprintf(w, "API is up and running\nenvironment: %s\n", app.config.env)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if(r.Method == http.MethodGet) {
		fmt.Fprintln(w, "Display a list of books on reading list")
		return
	  }
	  if(r.Method == http.MethodPost) {
		fmt.Fprintln(w, "Post a new book to the list")
		return
	  }
}


func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	
	switch r.Method {
	case http.MethodPut:
		fmt.Fprintln(w, "updte a book")
		//app.updateBook(w, r)
	case http.MethodGet:
		fmt.Fprintln(w, "GET a book")
		//app.getBook(w, r)
	case http.MethodDelete: 
	fmt.Fprintln(w, "Delete a book")
	  //app.deleteBook(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}


}