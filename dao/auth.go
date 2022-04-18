package dao

import (
	"context"
	"golang-mongo-api/database"
	"golang-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AdminFindByEmail(email string) (models.Admin, error) {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
		admin    models.Admin
	)

	// find
	filter := bson.M{
		"email": email,
	}
	err := adminCol.FindOne(ctx, filter).Decode(&admin)

	if err != nil {
		return admin, err
	}

	return admin, nil
}

func InitAdminUser() {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
	)

	count, _ := adminCol.CountDocuments(ctx, bson.D{})

	if count == 0 {
		admin := models.Admin{
			ID:       primitive.NewObjectID(),
			Email:    "admin",
			Password: "abc123@",
			Name:     "Ngo Thanh Tuan",
			//Age:      12,
		}
		adminCol.InsertOne(ctx, admin)
	}
}
