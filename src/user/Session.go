package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	UserID    primitive.ObjectID `bson:"user_id"`
	SessionID string             `bson:"session_id"`
}
