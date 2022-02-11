package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books
func getBooks(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(books)
}

// Get one book
func getBook(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // Get Parameters

	//Loop through the books and find the ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(responseWriter).Encode(item)
			return
		}
	}
	json.NewEncoder(responseWriter).Encode(&Book{})
}

// Insert a new book
func createBook(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(responseWriter).Encode(book)
}

// Update a book - this reuses code for create and delete book
func updateBook(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(responseWriter).Encode(book)
			return
		}
		json.NewEncoder(responseWriter).Encode(books)
	}
}

// Delete a Book
func deleteBook(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(responseWriter).Encode(books)
	}
}

func main() {

	// Init Router
	router := mux.NewRouter()

	// Hardcoded data
	books = append(books, Book{ID: "1", Isbn: "42355", Title: "Book One", Author: &Author{FirstName: "Ahmed", LastName: "Ali"}})
	books = append(books, Book{ID: "2", Isbn: "14672", Title: "Book Two", Author: &Author{FirstName: "Ronald", LastName: "McDonald"}})
	books = append(books, Book{ID: "3", Isbn: "73452", Title: "Book Three", Author: &Author{FirstName: "Thomas", LastName: "Jefferson"}})
	books = append(books, Book{ID: "4", Isbn: "67821", Title: "Book Four", Author: &Author{FirstName: "Tom", LastName: "Jerry"}})

	// Router handlers and endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Starts the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
