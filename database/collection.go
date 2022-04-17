package database

import "go.mongodb.org/mongo-driver/mongo"

// Collection Name
const (
	admins = "admins"
	users  = "users"
)

// AdminCol
func AdminCol() *mongo.Collection {
	return db.Collection(admins)
}

// UserCol...
func UserCol() *mongo.Collection {
	return db.Collection(users)
}
