package main

import (
    "bookstore_web/models"
    "fmt"
    "net/http"
    "strconv"
)

func main() {
    models.InitDB("root:root@tcp(localhost:3306)/gowebapp")

    http.HandleFunc("/books", getAllBooks)
    http.HandleFunc("/book/show", getBook)
    http.HandleFunc("/book/create", createBook)
    http.HandleFunc("/book/delete", deleteBook)
    http.ListenAndServe(":8080", nil)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }
    bks, err := models.GetAllBooks()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}

func getBook(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  isbn := r.FormValue("isbn")
  if isbn == "" {
    http.Error(w, http.StatusText(400), 400)
    return
  }

  bk, err := models.GetBook(isbn)
  
	if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
}

func createBook(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  isbn := r.FormValue("isbn")
  title := r.FormValue("title")
  author := r.FormValue("author")
  if isbn == "" || title == "" || author == "" {
    http.Error(w, http.StatusText(400), 400)
    return
  }
  price, err := strconv.ParseFloat(r.FormValue("price"), 32)
  if err != nil {
    http.Error(w, http.StatusText(400), 400)
    return
  }

  rowsAffected, err := models.CreateBook(isbn, title, author, price)
    
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", isbn, rowsAffected)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
  if r.Method != "DELETE" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  isbn := r.FormValue("isbn")
  if isbn == "" {
    http.Error(w, http.StatusText(400), 400)
    return
  }

   rowsAffected, err := models.DeleteBook(isbn)
  
	if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  fmt.Fprintf(w, "Book %s deleted successfully (%d row affected)\n", isbn, rowsAffected)
}