package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// AdminBSON
	Admin struct {
		ID       primitive.ObjectID `bson:"_id"`
		Email    string             `bson:"email"`
		Password string             `bson:"password"`
		Name     string             `bson:"name"`
	}

	// AdminLoginBody
	AdminLoginBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// AdminRegisterBody
	AdminRegisterBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
)

func (a AdminLoginBody) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required),
		validation.Field(&a.Password, validation.Required),
	)
}
