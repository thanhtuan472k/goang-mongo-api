package dao

import (
	"context"
	"golang-mongo-api/database"
	"golang-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckUserEmailExisted(email string) bool {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	// count documents having same email
	count, err := userCol.CountDocuments(ctx, bson.M{"email": email})

	if err != nil || count <= 0 {
		return false
	}

	return true
}

func CreateUser(doc models.User) error {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	// Insert
	_, err := userCol.InsertOne(ctx, doc)

	return err
}

func GetListUser() []models.User {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
		users   []models.User
	)

	cursor, err := userCol.Find(ctx, bson.D{})

	if err != nil {
		return nil
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil
	}

	return users
}

func UpdateUser(id primitive.ObjectID, user models.User) error {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	// Update
	_, err := userCol.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})

	if err != nil {
		return err
	}

	return nil
}

func DeletePlayer(ID primitive.ObjectID) error {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	if _, err := userCol.DeleteOne(ctx, bson.D{{"_id", ID}}); err != nil {
		return err
	}

	return nil
}
