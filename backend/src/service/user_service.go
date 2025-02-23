package service

import (
	"seaventures/src/models"
	"seaventures/src/repository"
	"golang.org/x/crypto/bcrypt"
	"errors"
	
)

func RegisterUser(user *models.User) error {

	existingUser, err := repository.GetUserByEmail(user.Email)
    if err == nil && existingUser != nil {
        return errors.New("email is already taken")
    }

	existingUser, err = repository.GetUserByUsername(user.UserName)
    if err == nil && existingUser != nil {
        return errors.New("username is already taken")
    }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

	return repository.Register(user)
}

func Login(user *models.User) error {

	existingUser, err := repository.GetUserByEmail(user.Email)
    if err != nil {
        return errors.New("invalid credentials")
    }

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
    if err != nil {
        return errors.New("invalid credentials")
    }

	user.ID = existingUser.ID

	return nil
}



	
