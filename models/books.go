package models

import (
//    "github.com/jinzhu/gorm"
)

// Book model
type Book struct {
//	gorm.Model
    Isbn   string  `json:"isbn" binding:"required"`
    Title  string `json:"title" binding:"required"`
    Author string `json:"author" binding:"required"`
    Price  float32 `json:"price" binding:"required"`
}

/**
Book DAO to deal with DB opertaions 
*/
func GetAllBooks() ([]Book) {
	// Query multiple records
	books := []Book{}
	db.Find(&books)
    return books
}

func GetBook(isbn string) (Book) {
  var book Book
// Read
   db.First(&book, "isbn = ?", isbn)// find product with id 
   
 return book
}

func CreateBook(book Book) {
	  db.Create(&book)
}

func UpdateBook(book Book) {
	  // Update multiple attributes with `struct`, will only update those changed & non blank fields
	db.Model(&book).Updates(Book{Title: book.Title,Author: book.Author, Price: book.Price})
}

func DeleteBook(isbn string)  {
	book := GetBook(isbn)
	db.Where("Isbn = ?", isbn).Delete(&book)
}

