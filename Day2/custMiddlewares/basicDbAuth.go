package custMiddlewares

import (
	"Alterra/batch5/ORM/Part1/configs"
	"Alterra/batch5/ORM/Part1/models"

	"github.com/labstack/echo/v4"
)

func DBBasicAuth(username, password string, c echo.Context) (bool, error) {
	checkUser := models.User{Name: username, Password: password}
	if err := configs.DB.Where("name = ? AND password = ?", checkUser.Name, checkUser.Password).Find(&checkUser).Error; err != nil || checkUser.ID == 0 {
		return false, nil
	}
	// fmt.Println(checkUser)
	return true, nil
}
