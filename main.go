package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID 		string `json:"id"`
	Isbn 		string `json:"isbn"`
	Title 	string `json:"title"`
	Author	*Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname	string	`json:"firstname"`
	Lastname		string	`json:"lastname"`
}

// Error Struct
type Error struct {
	Error	string	`json:"error"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Find book with id

	for _, item := range books {
		if item.ID == params["id"] {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(404)
	json.NewEncoder(w).Encode(Error{Error: "Not found"})
}

// Create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(len(books) + 1)// Mock ID
	books = append(books, book)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book) // Mock ID
	
	for index, item := range books {
		if item.ID == params["id"] {
			book.ID = item.ID
			books[index] = book
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	w.WriteHeader(404)
	json.NewEncoder(w).Encode(Error{Error: "Not found"})
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "243534", Title: "Book 1", Author: &Author{Firstname: "Milton", Lastname: "Boos"}})
	books = append(books, Book{ID: "2", Isbn: "342554", Title: "Book 2", Author: &Author{Firstname: "John", Lastname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "453465", Title: "Book 3", Author: &Author{Firstname: "Lucas", Lastname: "Jhonson"}})
	books = append(books, Book{ID: "4", Isbn: "356775", Title: "Book 4", Author: &Author{Firstname: "Leo", Lastname: "Finnister"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}