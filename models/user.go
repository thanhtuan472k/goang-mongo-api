package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
    // UserBSON
    UserBSON struct {
        ID primitive.ObjectID `bson:"_id"`
        Name string           `bson:"name"`
        Age int64             `bson:"age"`
    }

    // UserDetail
    UserDetail struct {
        ID primitive.ObjectID `json:"_id"`
        Name string           `json:"name"`
        Age int64             `json:"age"` 
    }
)
