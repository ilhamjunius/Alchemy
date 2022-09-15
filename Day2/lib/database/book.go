package database

import (
	"Alterra/batch5/ORM/Part1/configs"
	"Alterra/batch5/ORM/Part1/models"
)

// get all books

func GetAllBooksController() ([]models.Book, error) {
	var books []models.Book

	if err := configs.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

// get book by id
func GetBookController(id int) ([]models.Book, error) {
	// your solution here
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Incorrect book ID")
	// }

	var book []models.Book

	if err := configs.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// create new book
func CreateBookController(u models.Book) ([]models.Book, error) {
	var book []models.Book

	if err := configs.DB.Save(&u).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// delete book by id
func DeleteBookController(id int) ([]models.Book, error) {
	// your solution here
	var book []models.Book
	if err := configs.DB.Find(&book, "id=?", id).Error; err != nil {
		return book, err
	}
	if err := configs.DB.Delete(&book).Error; err != nil {
		return book, nil
	}
	return book, nil
}

// update book by id
func UpdateBookController(tittle, author string, id int) ([]models.Book, error) {
	// your solution here
	var book []models.Book

	if err := configs.DB.Model(&models.Book{}).Where("id = ?", id).Updates(models.Book{Tittle: tittle, Author: author}).Error; err != nil {
		return book, err
	}
	// fmt.Println(book)
	return book, nil

}
