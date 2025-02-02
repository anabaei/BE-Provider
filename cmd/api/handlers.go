package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"readinglist.duffney.io/internal/data"
	"strconv"
	"time"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	data := map[string]string{
		"status":      "ok",
		"environment": app.config.env,
		"version":     "1.0.0",
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	js = append(js, '\n')
	w.Write(js)

}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Display a list of books on reading list")
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Post a new book to the list")
		return
	}
}

func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		fmt.Fprintln(w, "update a book")
		app.updateBook(w, r)
	case http.MethodGet:
		fmt.Fprintln(w, "GET a book")
		app.getBook(w, r)
	case http.MethodPost:
		app.postBook(w, r)
	case http.MethodDelete:
		fmt.Fprintln(w, "Delete a book")
		app.deleteBook(w, r)
	default:
		fmt.Fprintf(w, "came here")
		app.updateBook(w, r)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(string(id), 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	book := data.Book{

		ID:        idInt,
		Title:     "The Great Gatsby",
		CreatedAt: time.Now(),
		Published: time.Date(2015, 4, 10, 0, 0, 0, 0, time.UTC),
		Pages:     218,
		Genres:    []string{"Fiction", "Tragedy"},
		Rating:    4.5,
		Version:   1,
	}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(string(id), 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "update book with ID %d\n", idInt)
}

func (app *application) postBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(string(id), 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "post book with ID %d\n", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(string(id), 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "delete book with ID %d\n", idInt)
}
