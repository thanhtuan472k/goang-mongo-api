package service

import (
	"errors"
	"myapp/dao"
	"myapp/models"
	"myapp/utils"
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
