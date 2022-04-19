package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// User
	User struct {
		ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // meta-tag for update & delete
		Email    string             `bson:"email"`
		Password string             `bson:"password"`
		Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	}

	// UserCreateBody - UserRegisterBody
	UserCreateBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	// UserLoginBody
	UserUpdateBody struct {
		Name string `json:"name"`
	}

	Page struct {
	}
)

func (u UserCreateBody) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Name, validation.Required),
	)
}

func (u UserUpdateBody) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required),
	)
}
