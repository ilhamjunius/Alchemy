package database

import (
	"Alterra/batch5/ORM/Part1/configs"
	"Alterra/batch5/ORM/Part1/models"
)

// get all users

func GetAllUsersController() ([]models.User, error) {
	var users []models.User

	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// get user by id
func GetUserController(id int) ([]models.User, error) {
	// your solution here
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Incorrect user ID")
	// }

	var user []models.User

	if err := configs.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// create new user
func CreateUserController(u models.User) (models.User, error) {
	// var user []models.User

	if err := configs.DB.Save(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// delete user by id
func DeleteUserController(id int) ([]models.User, error) {
	// your solution here
	var user []models.User
	if err := configs.DB.Find(&user, "id=?", id).Error; err != nil {
		return user, err
	}
	if err := configs.DB.Delete(&user).Error; err != nil {
		return user, nil
	}
	return user, nil
}

// update user by id
func UpdateUserController(name, email, password string, id int) ([]models.User, error) {
	// your solution here
	var user []models.User

	if err := configs.DB.Model(&models.User{}).Where("id = ?", id).Updates(models.User{Name: name, Email: email, Password: password}).Error; err != nil {
		return user, err
	}
	// fmt.Println(user)
	return user, nil

}
func Login(email, password string) (models.User, error) {
	foundUser := models.User{}
	if err := configs.DB.Where("email = ? AND password = ?", email, password).Find(&foundUser).Error; err != nil {
		return foundUser, err
	}
	return foundUser, nil
}
