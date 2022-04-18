package service

import (
	"errors"
	"myapp/dao"
	"myapp/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create new user
func CreateUser(createBody models.UserCreateBody) error {
	// check email existed
	isEmailExisted := dao.CheckUserEmailExisted(createBody.Email)

	if isEmailExisted {
		return errors.New("email existed")
	}

	// create user
	userCreateBSON := models.User{
		ID:       primitive.NewObjectID(),
		Email:    createBody.Email,
		Password: createBody.Password,
		Name:     createBody.Name,
	}

	if err := dao.CreateUser(userCreateBSON); err != nil {
		return errors.New("can not create user")
	}
	return nil
}

// Get list user
func GetListUser() []models.User {
	return dao.GetListUser()
}

// Update user
func UpdateUser(ID string, userUpdateBody models.UserUpdateBody) error {
	user := models.User{
		Name: userUpdateBody.Name,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	return dao.UpdateUser(objID, user)
}

func DeleteUser(ID string) error {
	objID, _ := primitive.ObjectIDFromHex(ID)
	return dao.DeletePlayer(objID)
}
