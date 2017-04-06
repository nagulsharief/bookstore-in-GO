package main

import (
	"bookstore_web/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// This is the starting point of GO app
func main() {
	// initialize DB
	models.InitDB("root:root@tcp(localhost:3306)/gowebapp")

	// third party router library to map url's to handlers and to deal with path params
	router := httprouter.New()
	router.GET("/books", getAllBooks)
	router.GET("/book/:id", getBook)
	router.POST("/book/create", createBook)
	router.DELETE("/book/:id", deleteBook)

	// initialize HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}

/**
This method is to get all the bookss
*/
func getAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bks, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(bks); err != nil {
		panic(err)
	}
}

/**
This method is to get the book by id
*/
func getBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isbn := params.ByName("id")
	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	bk, err := models.GetBook(isbn)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

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

	rowsAffected, err := models.CreateBook(book)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	} else if rowsAffected == 1 {
		if err := json.NewEncoder(w).Encode(http.StatusOK); err != nil {
			panic(err)
		}
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
	rowsAffected, err := models.DeleteBook(isbn)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	} else if rowsAffected == 1 {
		if err := json.NewEncoder(w).Encode(http.StatusOK); err != nil {
			panic(err)
		}
	}
}
