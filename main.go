package main

import (
	"log"
	"net/http"
	"encoding/json"
	"bookstore_web/models"
	"github.com/julienschmidt/httprouter"
)

// This is the starting point of GO app
func main() {
	// initialize DB
	models.InitDB("root:root@tcp(localhost:3306)/gowebapp")

	// third party router library to map url's to handlers and to deal with path params
	router := httprouter.New()
	router.GET("/books", getAllBooks)
	router.GET("/book/:id", getBook)
	router.POST("/book/", createBook)
	router.PUT("/book/:id", updateBook)
	router.DELETE("/book/:id", deleteBook)

	// initialize HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}

/**
This method is to get all the bookss
*/
func getAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bks := models.GetAllBooks()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bks); err != nil {
		panic(err)
	}
}

/**
This method is to get the book by id
*/
func getBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	isbn := params.ByName("id")
	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	bk := models.GetBook(isbn)
	if err := json.NewEncoder(w).Encode(bk); err != nil {
		panic(err)
	}

}

/**
This method is to create the book
*/
func createBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var book models.Book
	if r.Body == nil {
		http.Error(w, "Please send book details to create in the request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	models.CreateBook(book)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(http.StatusOK); err != nil {
		panic(err)
	}
}

/**
This method is to update the book
*/
func updateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var book models.Book
	if r.Body == nil {
		http.Error(w, "Please send book details to create in the request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	models.UpdateBook(book)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(http.StatusOK); err != nil {
		panic(err)
	}
}


/**
This method is to delete the book
*/
func deleteBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isbn := params.ByName("id")
	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	models.DeleteBook(isbn)
	w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(http.StatusOK); err != nil {
			panic(err)
	}
}
