package database

import "go.mongodb.org/mongo-driver/mongo"

// Collection Name
const (
	users = "users"
)

// UserCol...
func UserCol() *mongo.Collection {
	return db.Collection(users)
}