package database

import (
	"github.com/alchemy/Day2/models"
)

var books = []models.Book{
	{ID: 1, Tittle: "Science", Author: "Ilham Junius"},
	{ID: 2, Tittle: "Technology", Author: "Ilham"},
	{ID: 3, Tittle: "Art", Author: "Junius"},
	{ID: 4, Tittle: "Tale", Author: "Rudi"},
}

// get all books

func GetAllBooksController() ([]models.Book, error) {

	return books, nil
}

// get book by id
func GetBookController(id int) ([]models.Book, error) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Incorrect book ID")
	// }

	var book []models.Book
	book = append(book, books[id-1])
	return book, nil
}

// create new book
func CreateBookController(book models.Book) ([]models.Book, error) {
	book.ID = len(books) + 1
	books = append(books, book)

	return books, nil
}

// delete book by id
func DeleteBookController(id int) ([]models.Book, error) {
	var newBooks = []models.Book{}
	for _, book := range books {
		if book.ID != id {
			newBooks = append(newBooks, book)
		}

	}
	books = newBooks
	return books, nil
}

// update book by id
func UpdateBookController(tittle, author string, id int) ([]models.Book, error) {
	var newBooks []models.Book

	books[id-1].Author = author
	books[id-1].Tittle = tittle
	for _, book := range books {
		if book.ID == id {
			newBooks = append(newBooks, book)
			break
		}

	}
	return newBooks, nil

}
