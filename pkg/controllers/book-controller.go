package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ShazFr/go-book-mgmt/pkg/models"
	"github.com/ShazFr/go-book-mgmt/pkg/utils"
	"github.com/gorilla/mux"
)

// var newBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)

	createdBook := models.CreateBook(newBook)

	res, _ := json.Marshal(createdBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["id"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing book ID when Deleting Book")
	}

	deletedBook := models.DeleteBook(id)
	res, _ := json.Marshal(deletedBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	// Create new book from request
	newBookDetails := &models.Book{}
	utils.ParseBody(r, newBookDetails)

	// Get Book ID of book to be updated
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	// get book's existing details
	bookDetails, db := models.GetBookById(id)

	// Update the details
	if newBookDetails.Name != "" {
		bookDetails.Name = newBookDetails.Name
	}
	if newBookDetails.Author != "" {
		bookDetails.Author = newBookDetails.Author
	}

	// Save the book details in DB
	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func summ() {

}
