package models

import (
	"github.com/ShazFr/go-book-mgmt/pkg/config"
	"gorm.io/gorm"
)

func init() {

}

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(book *Book) *Book {
	db.Create(book)
	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	result := db.First(&book, id) // shorthand for WHERE id = ?
	return &book, result
}

func DeleteBook(id int64) *Book {
	var bookToGet Book
	db.Where("ID=?", id).Delete(&bookToGet)
	return &bookToGet
}
