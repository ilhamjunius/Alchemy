package controllers

import (
	"net/http"
	"strconv"

	"github.com/alchemy/Day2/lib/database"
	"github.com/alchemy/Day2/models"
	"github.com/labstack/echo/v4"
)

func GetAllUsersController(c echo.Context) error {
	res, err := database.GetAllUsersController()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   res,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Incorrect user ID")
	}

	res, err := database.GetUserController(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"users":   res,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	res, err := database.CreateUserController(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    res,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Incorrect user ID")
	}

	if _, err := database.DeleteUserController(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	id, err := strconv.Atoi(c.Param("id"))
	user.ID = uint(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Incorrect user ID")
	}

	if _, err := database.UpdateUserController(user.Name, user.Email, user.Password, id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})

}
