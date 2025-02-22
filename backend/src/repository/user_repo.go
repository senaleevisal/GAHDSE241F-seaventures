package repository

import (
	"seaventures/src/config"
	"seaventures/src/models"
)

func Register(user *models.User) error {
	return config.DB.Create(user).Error
}

func Login(user *models.User) error {
	return config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(userName string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

