package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alchemy/Day2/lib/database"
	"github.com/alchemy/Day2/models"
	"github.com/labstack/echo/v4"
)

func GetAllBooksController(c echo.Context) error {
	res, err := database.GetAllBooksController()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   res,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Incorrect book ID")
	}

	res, err := database.GetBookController(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"books":   res,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	_, err := database.CreateBookController(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Incorrect book ID")
	}

	if _, err := database.DeleteBookController(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Incorrect book ID")
	}
	if _, err := database.UpdateBookController(book.Tittle, book.Author, id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	book.ID = id
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success update book id %v", id),
	})

}
