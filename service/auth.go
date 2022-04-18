package service

import (
	"errors"
	"golang-mongo-api/dao"
	"golang-mongo-api/models"
	"golang-mongo-api/utils"
)

// AdminLogin ...
func AdminLogin(loginBody models.AdminLoginBody) (string, error) {
	// find admin in db
	admin, err := dao.AdminFindByEmail(loginBody.Email)

	if err != nil {
		return "", err
	}

	// verify admin password
	if admin.Password != loginBody.Password {
		return "", errors.New("wrong password")
	}

	data := map[string]interface{}{
		"id":      admin.ID,
		"isAdmin": true,
	}

	// return JWT token
	return utils.GenerateUserToken(data), err
}
