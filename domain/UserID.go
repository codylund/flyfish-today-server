package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserID struct {
	ID primitive.ObjectID `bson:"_id"`
}
